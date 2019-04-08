package orm

import "reflect"

func getTableName(ind reflect.Value) string {
	return snakeString(ind.Type().Name())
}

func getFullName(typ reflect.Type) string {
	return typ.PkgPath() + "/" + typ.Name()
}
