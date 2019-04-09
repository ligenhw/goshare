package orm

import (
	"fmt"
	"os"
	"reflect"
)

type modelInfo struct {
	fields    *fields
	table     string
	fullName  string
	addrField reflect.Value
}

func newModelInfo(val reflect.Value) *modelInfo {
	mi := &modelInfo{}
	mi.addrField = val
	ind := reflect.Indirect(val)
	mi.fields = newFields()
	addModelFields(mi, ind)
	return mi
}

func addModelFields(mi *modelInfo, ind reflect.Value) {
	var (
		err error
		fi  *fieldInfo
		sf  reflect.StructField
	)

	for i := 0; i < ind.NumField(); i++ {
		sf = ind.Type().Field(i)
		field := ind.Field(i)
		fi, err = newFieldInfo(sf, field)
		if err == errSkipField {
			err = nil
			continue
		} else if err != nil {
			break
		}

		mi.fields.Add(fi)

		if fi.pk {
			if mi.fields.pk != nil {
				err = fmt.Errorf("one model must have one pk field only")
				break
			} else {
				mi.fields.pk = fi
			}
		}
	}

	if err != nil {
		fmt.Println(fmt.Errorf("field: %s.%s, %s", ind.Type(), sf.Name, err))
		os.Exit(2)
	}
}
