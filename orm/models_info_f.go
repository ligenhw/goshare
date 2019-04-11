package orm

import (
	"errors"
	"reflect"
)

var errSkipField = errors.New("skip field")

type fields struct {
	pk      *fieldInfo
	columns map[string]*fieldInfo
	dbcols  []string
}

func (f *fields) Add(fi *fieldInfo) {
	f.columns[fi.column] = fi
	f.dbcols = append(f.dbcols, fi.column)
}

// get field info by column name
func (f *fields) GetByColumn(column string) *fieldInfo {
	return f.columns[column]
}

// create new field info collection
func newFields() *fields {
	f := new(fields)
	f.columns = make(map[string]*fieldInfo)
	return f
}

type fieldInfo struct {
	column     string
	fieldIndex []int
	typ        reflect.Type
	addrField  reflect.Value
	sf         reflect.StructField
	auto       bool
	pk         bool
	null       bool
}

func newFieldInfo(sf reflect.StructField, val reflect.Value) (fi *fieldInfo, err error) {
	fi = new(fieldInfo)

	column := snakeString(sf.Name)
	attrs, tags := parseStructTag(sf.Tag.Get(defaultStructTagName))

	if _, ok := attrs["-"]; ok {
		return nil, errSkipField
	}

	if tags["column"] != "" {
		fi.column = column
	} else {
		fi.column = column
	}
	fi.addrField = val
	fi.typ = val.Type()
	fi.sf = sf

	fi.auto = attrs["auto"]
	fi.pk = attrs["pk"]
	fi.null = attrs["null"]

	return
}
