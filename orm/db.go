package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
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
	field := ind.FieldByIndex(fi.fieldIndex)
	value = field.Interface()
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

// TODO: implents
func (d *dbBase) InsertMulti(q *sql.DB, mi *modelInfo, sind reflect.Value) (int64, error) {
	length := sind.Len()
	for i := 1; i <= length; i++ {

		// ind := reflect.Indirect(sind.Index(i - 1))
	}

	return 0, nil
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

	refs := newRefs(mi, mi.fields.dbcols)

	row := q.QueryRow(query, args...)
	if err := row.Scan(refs...); err != nil {
		if err == sql.ErrNoRows {
			return ErrNoRows
		}
		return err
	}

	elm := reflect.New(mi.addrField.Elem().Type())
	mind := reflect.Indirect(elm)
	d.setColsValues(mi, &mind, mi.fields.dbcols, refs)
	ind.Set(mind)
	return nil
}

func (d *dbBase) Delete(q *sql.DB, mi *modelInfo, ind reflect.Value, cols []string) (int64, error) {
	var whereCols []string
	var args []interface{}

	// if specify cols length > 0, then use it for where condition.
	if len(cols) > 0 {
		var err error
		whereCols = make([]string, 0, len(cols))
		args, err = d.collectValues(mi, ind, cols, false, &whereCols)
		if err != nil {
			return 0, err
		}
	} else {
		// default use pk value as where condtion.
		pkColumn, pkValue, ok := getExistPk(mi, ind)
		if !ok {
			return 0, ErrMissPK
		}
		whereCols = []string{pkColumn}
		args = append(args, pkValue)
	}

	wheres := strings.Join(whereCols, " = ? AND ")

	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", mi.table, wheres)

	res, err := q.Exec(query, args...)
	if err == nil {
		num, err := res.RowsAffected()
		if err != nil {
			return 0, err
		}

		return num, nil
	}

	return 0, err
}

// execute update sql dbQuerier with given struct reflect.Value.
func (d *dbBase) Update(q *sql.DB, mi *modelInfo, ind reflect.Value, cols []string) (int64, error) {
	pkName, pkValue, ok := getExistPk(mi, ind)
	if !ok {
		return 0, ErrMissPK
	}

	var setNames []string

	// if specify cols length is zero, then commit all columns.
	if len(cols) == 0 {
		cols = mi.fields.dbcols
		setNames = make([]string, 0, len(mi.fields.dbcols)-1)
	} else {
		setNames = make([]string, 0, len(cols))
	}

	setValues, err := d.collectValues(mi, ind, cols, true, &setNames)
	if err != nil {
		return 0, err
	}

	setValues = append(setValues, pkValue)

	setColumns := strings.Join(setNames, " = ? ,")

	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?", mi.table, setColumns, pkName)

	log.Println(query)
	log.Println(setValues)

	res, err := q.Exec(query, setValues...)
	if err == nil {
		return res.RowsAffected()
	}
	return 0, err
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
	switch value.(type) {
	case sql.NullInt64:
		nint := value.(sql.NullInt64)
		if nint.Valid {
			val, _ := nint.Value()
			field.SetInt(val.(int64))
		} else {
			field.Set(reflect.Zero(reflect.TypeOf(new(int))))
		}
	default:
		field.Set(reflect.ValueOf(value))
	}

	return value, nil
}

// flag of RETURNING sql.
func (d *dbBase) HasReturningID(*modelInfo, *string) bool {
	return false
}

// read related records.
func (d *dbBase) ReadBatch(q *sql.DB, qs *QuerySeter, mi *modelInfo, cond *Condition, container interface{}, cols []string) (int64, error) {

	val := reflect.ValueOf(container)
	ind := reflect.Indirect(val)

	errTyp := true
	one := true

	if val.Kind() == reflect.Ptr {
		fn := ""
		if ind.Kind() == reflect.Slice {
			one = false
			typ := ind.Type().Elem()
			switch typ.Kind() {
			case reflect.Ptr:
				fn = getFullName(typ.Elem())
			case reflect.Struct:
				fn = getFullName(typ)
			}
		} else {
			fn = getFullName(ind.Type())
		}
		errTyp = fn != mi.fullName
	}

	if errTyp {
		if one {
			panic(fmt.Errorf("wrong object type `%s` for rows scan, need *%s", val.Type(), mi.fullName))
		} else {
			panic(fmt.Errorf("wrong object type `%s` for rows scan, need *[]*%s or *[]%s", val.Type(), mi.fullName, mi.fullName))
		}
	}

	var tCols []string
	if len(cols) > 0 {
		tCols = cols
	} else {
		tCols = mi.fields.dbcols
	}

	rlimit := qs.limit
	offset := qs.offset

	sels := strings.Join(tCols, ", ")
	where, args := getCondSQL(cond)
	orderBy := getOrderSQL(qs.orders)
	limit := getLimitSQL(mi, offset, rlimit)

	sqlSelect := "SELECT"
	if qs.distinct {
		sqlSelect += " DISTINCT"
	}
	query := fmt.Sprintf("%s %s FROM %s %s %s %s", sqlSelect, sels, mi.table, where, orderBy, limit)

	log.Println(query)
	log.Println(args...)

	rs, err := q.Query(query, args...)
	if err != nil {
		return 0, err
	}

	defer rs.Close()
	slice := ind

	refs := newRefs(mi, tCols)

	var cnt int64
	for rs.Next() {
		if err := rs.Scan(refs...); err != nil {
			if err == sql.ErrNoRows {
				return cnt, ErrNoRows
			}
			return cnt, err
		}

		elm := reflect.New(mi.addrField.Elem().Type())
		mind := reflect.Indirect(elm)
		d.setColsValues(mi, &mind, tCols, refs)
		slice = reflect.Append(slice, mind.Addr())
		cnt++
	}

	if cnt > 0 {
		ind.Set(slice)
	}

	return cnt, nil
}

func newRefs(mi *modelInfo, cols []string) (refs []interface{}) {
	colsNum := len(cols)
	refs = make([]interface{}, colsNum)
	for i := range refs {
		// var ref interface{}
		column := cols[i]
		fi := mi.fields.GetByColumn(column)
		switch fi.typ.Kind() {
		case reflect.Int:
			v := new(int)
			refs[i] = v
		case reflect.String:
			v := new(string)
			refs[i] = v
		default:
			switch fi.addrField.Interface().(type) {
			case time.Time:
				v := new(time.Time)
				refs[i] = v
			case *int:
				v := new(sql.NullInt64)
				refs[i] = v
			default:
				log.Println("warning not support type : ", fi.typ.Kind())
			}
		}
	}

	return
}
