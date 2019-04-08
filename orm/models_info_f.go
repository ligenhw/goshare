package orm

import "reflect"

type fieldInfo struct {
	column string
	value  interface{}
}

func newFieldInfo(sf reflect.StructField, val reflect.Value) *fieldInfo {
	column := snakeString(sf.Name)
	return &fieldInfo{
		column: column,
		value:  val.Interface(),
	}
}
