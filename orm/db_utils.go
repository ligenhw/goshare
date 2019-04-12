package orm

import "reflect"

// get pk column info.
func getExistPk(mi *modelInfo, ind reflect.Value) (column string, value interface{}, exist bool) {
	fi := mi.fields.pk
	if fi != nil {
		column = fi.column
		value = ind.FieldByIndex(fi.fieldIndex).Interface()
		exist = true
	} else {
		exist = false
	}

	return
}
