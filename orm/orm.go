package orm

import (
	"database/sql"
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
//		num, err = o.Delete(&u)
//	}
//

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
