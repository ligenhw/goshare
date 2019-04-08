package orm

import (
	"log"
	"reflect"
)

type modelInfo struct {
	name      string
	fields    []*fieldInfo
	table     string
	addrField reflect.Value
}

func newModelInfo(ind reflect.Value) *modelInfo {
	fields := make([]*fieldInfo, 0)
	typ := ind.Type()
	for i := 0; i < ind.NumField(); i++ {
		sf := typ.Field(i)
		v := ind.Field(i)
		fi := newFieldInfo(sf, v)
		log.Println(*fi)
		fields = append(fields, fi)
	}

	return &modelInfo{
		name:   typ.Name(),
		fields: fields,
	}
}
