package orm

import (
	"fmt"
	"os"
	"reflect"
)

type modelInfo struct {
	fields    []*fieldInfo
	table     string
	fullName  string
	addrField reflect.Value
}

func newModelInfo(val reflect.Value) *modelInfo {
	mi := &modelInfo{}
	mi.addrField = val
	ind := reflect.Indirect(val)
	mi.fields = make([]*fieldInfo, 0)
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
		mi.fields = append(mi.fields, fi)
	}

	if err != nil {
		fmt.Println(fmt.Errorf("field: %s.%s, %s", ind.Type(), sf.Name, err))
		os.Exit(2)
	}
}
