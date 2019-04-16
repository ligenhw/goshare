package orm

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func RegisterModel(model interface{}) {
	val := reflect.ValueOf(model)
	ind := reflect.Indirect(val)
	typ := ind.Type()

	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("<orm.RegisterModel> cannot use non-ptr model struct `%s`", getFullName(typ)))
	}

	if typ.Kind() == reflect.Ptr {
		panic(fmt.Errorf("<orm.RegisterModel> only allow ptr model struct, it looks you use two reference to the struct `%s`", typ))
	}

	mi := newModelInfo(val)
	if mi.fields.pk == nil {
		for _, fi := range mi.fields.columns {
			if strings.ToLower(fi.column) == "id" {
				fi.pk = true
				mi.fields.pk = fi
			}
		}

		if mi.fields.pk == nil {
			fmt.Printf("<orm.RegisterModel> `%s` needs a primary key field, default is to use 'id' if not set\n", typ.Name())
			os.Exit(2)
		}

	}

	table := getTableName(ind)
	mi.table = table
	mi.fullName = getFullName(typ)

	modelCache.set(table, mi)
}
