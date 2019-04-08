package orm

import (
	"fmt"
	"log"
	"reflect"
)

func registerModel(model interface{}) {
	val := reflect.ValueOf(model)
	ind := reflect.Indirect(val)
	typ := ind.Type()

	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("<orm.RegisterModel> cannot use non-ptr model struct `%s`", getFullName(typ)))
	}

	if typ.Kind() == reflect.Ptr {
		panic(fmt.Errorf("<orm.RegisterModel> only allow ptr model struct, it looks you use two reference to the struct `%s`", typ))
	}

	table := getTableName(ind)
	name := getFullName(typ)
	mi := newModelInfo(ind)

	mi.table = table

	log.Println(mi, table, name)
}
