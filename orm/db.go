package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var (
	// ErrMissPK missing pk error
	ErrMissPK = errors.New("missed pk value")
)

type dbBase struct {
}

// get struct columns values as interface slice.
func (d *dbBase) collectValues(mi *modelInfo, ind reflect.Value, cols []string, skipAuto bool, names *[]string) (values []interface{}, err error) {
	if names == nil {
		ns := make([]string, 0, len(cols))
		names = &ns
	}
	values = make([]interface{}, 0, len(cols))

	for _, column := range cols {
		var fi *fieldInfo
		if fi = mi.fields.GetByColumn(column); fi != nil {
			column = fi.column
		} else {
			panic(fmt.Errorf("wrong db field/column name `%s` for model `%s`", column, mi.fullName))
		}

		if fi.auto && skipAuto {
			continue
		}
		value, err := d.collectFieldValue(mi, fi, ind)
		if err != nil {
			return nil, err
		}

		*names, values = append(*names, column), append(values, value)
	}

	return
}

// get one field value in struct column as interface.
func (d *dbBase) collectFieldValue(mi *modelInfo, fi *fieldInfo, ind reflect.Value) (interface{}, error) {
	var value interface{}
	typ := reflect.Indirect(fi.addrField)
	value = typ.Interface()
	return value, nil
}

// execute insert sql with given struct and given values.
// insert the given values, not the field values in struct.
func (d *dbBase) InsertValue(q *sql.DB, mi *modelInfo, names []string, values []interface{}) (int64, error) {

	marks := make([]string, len(names))
	for i := range marks {
		marks[i] = "?"
	}
	columns := strings.Join(names, ", ")
	qmarks := strings.Join(marks, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", mi.table, columns, qmarks)
	log.Println(query)
	log.Println(values)

	if !d.HasReturningID(mi, &query) {
		res, err := q.Exec(query, values...)
		if err == nil {
			return res.LastInsertId()
		}
		return 0, err
	}

	row := q.QueryRow(query, values...)
	var id int64
	err := row.Scan(&id)
	return id, err
}

// execute insert sql dbQuerier with given struct reflect.Value.
func (d *dbBase) Insert(q *sql.DB, mi *modelInfo, ind reflect.Value) (int64, error) {
	names := make([]string, 0, len(mi.fields.dbcols))
	values, err := d.collectValues(mi, ind, mi.fields.dbcols, true, &names)
	if err != nil {
		return 0, err
	}

	id, err := d.InsertValue(q, mi, names, values)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (d *dbBase) Read(q *sql.DB, mi *modelInfo, ind reflect.Value, cols []string) error {
	var whereCols []string
	var args []interface{}

	// if specify cols length > 0, then use it for where condition.
	if len(cols) > 0 {
		var err error
		whereCols = make([]string, 0, len(cols))
		args, err = d.collectValues(mi, ind, cols, false, &whereCols)
		if err != nil {
			return err
		}
	} else {
		// default use pk value as where condtion.
		pkColumn, pkValue, ok := getExistPk(mi, ind)
		if !ok {
			return ErrMissPK
		}
		whereCols = []string{pkColumn}
		args = append(args, pkValue)
	}

	sels := strings.Join(mi.fields.dbcols, ", ")
	wheres := strings.Join(whereCols, " = ? AND ")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", sels, mi.table, wheres)

	log.Println(query)
	log.Println(args)

	colsNum := len(mi.fields.dbcols)
	refs := make([]interface{}, colsNum)
	for i := range refs {
		// var ref interface{}
		column := mi.fields.dbcols[i]
		fi := mi.fields.GetByColumn(column)
		switch fi.typ.Kind() {
		case reflect.Int:
			v := new(int)
			refs[i] = v
		case reflect.String:
			v := new(string)
			refs[i] = v
		default:
			log.Println("warning not support type : ", fi.typ.Kind())
		}
	}
	row := q.QueryRow(query, args...)
	if err := row.Scan(refs...); err != nil {
		if err == sql.ErrNoRows {
			return ErrNoRows
		}
		return err
	}

	for _, r := range refs {
		r := reflect.Indirect(reflect.ValueOf(r))
		log.Println("after scan : ", r.Type().Kind(), r.Interface())
	}

	elm := reflect.New(mi.addrField.Elem().Type())
	mind := reflect.Indirect(elm)
	d.setColsValues(mi, &mind, mi.fields.dbcols, refs)
	ind.Set(mind)
	return nil
}

// set values to struct column.
func (d *dbBase) setColsValues(mi *modelInfo, ind *reflect.Value, cols []string, values []interface{}) {
	for i, column := range cols {
		val := reflect.Indirect(reflect.ValueOf(values[i])).Interface()
		fi := mi.fields.GetByColumn(column)
		field := ind.FieldByIndex(fi.fieldIndex)
		_, err := d.setFieldValue(fi, val, field)

		if err != nil {
			panic(fmt.Errorf("Raw value: `%v` %s", val, err.Error()))
		}
	}
}

// set one value to struct column field.
func (d *dbBase) setFieldValue(fi *fieldInfo, value interface{}, field reflect.Value) (interface{}, error) {
	field.Set(reflect.ValueOf(value))
	return value, nil
}

// flag of RETURNING sql.
func (d *dbBase) HasReturningID(*modelInfo, *string) bool {
	return false
}
