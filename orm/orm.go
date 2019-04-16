package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

//	// Model Struct
//	type User struct {
//		Id   int    `orm:"auto"`
//		Name string `orm:"size(100)"`
//	}
//
//	func init() {
//		orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)
//	}
//
//	func main() {
//		o := orm.NewOrm()
//		user := User{Name: "slene"}
//		// insert
//		id, err := o.Insert(&user)
//		// update
//		user.Name = "astaxie"
//		num, err := o.Update(&user)
//		// read one
//		u := User{Id: user.Id}
//		err = o.Read(&u)
//		// delete
//      num, err = o.Delete(&u)

// var posts []*Post
// qs := o.QueryTable("post")
// num, err := qs.Filter("name", "slene").All(&posts)
//	}

var (
	DefaultRowsLimit = 1000
	ErrNoRows        = errors.New("<QuerySeter> no row found")
)

type orm struct {
	DbBaser dbBase
	db      *sql.DB
}

func NewOrm(db *sql.DB) *orm {
	o := new(orm)
	o.db = db
	return o
}

// get model info and model reflect value
func (o *orm) getMiInd(md interface{}) (mi *modelInfo, ind reflect.Value) {
	val := reflect.ValueOf(md)
	ind = reflect.Indirect(val)
	typ := ind.Type()

	name := getFullName(typ)
	if mi, ok := modelCache.getByFullName(name); ok {
		return mi, ind
	}
	panic(fmt.Errorf("<Ormer> table: `%s` not found, make sure it was registered with `RegisterModel()`", name))
}

// insert model data to database
func (o *orm) Insert(md interface{}) (int64, error) {
	mi, ind := o.getMiInd(md)
	return o.DbBaser.Insert(o.db, mi, ind)
}

// read data to model
func (o *orm) Read(md interface{}, cols ...string) error {
	mi, ind := o.getMiInd(md)
	return o.DbBaser.Read(o.db, mi, ind, cols)
}

// update model to database.
// cols set the columns those want to update.
func (o *orm) Update(md interface{}, cols ...string) (int64, error) {
	mi, ind := o.getMiInd(md)
	return o.DbBaser.Update(o.db, mi, ind, cols)
}

// delete model in database
// cols shows the delete conditions values read from. default is pk
func (o *orm) Delete(md interface{}, cols ...string) (int64, error) {
	mi, ind := o.getMiInd(md)
	num, err := o.DbBaser.Delete(o.db, mi, ind, cols)
	if err != nil {
		return num, err
	}
	return num, nil
}

func (o *orm) QueryTable(md interface{}) (qs *QuerySeter) {
	mi, _ := o.getMiInd(md)
	qs = newQuerySet(o, mi)
	return
}
