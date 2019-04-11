package orm

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/ligenhw/goshare/store"
)

func TestReflect(t *testing.T) {
	type T1 struct {
		A int
		B string
	}
	t1 := T1{23, "skidoo"}
	fmt.Println(t1)

	s := reflect.ValueOf(&t1).Elem()
	typeofT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s %s = %v\n", typeofT.Field(i).Name, f.Type(), f.Interface())
	}
}

type UserInfo struct {
	Id       int    `orm:"auto"`
	UserName string `orm:"varchar(20)"`
	PassWord string `orm:"varchar(100)"`
	Age      int
	Ext      string
	Time     time.Time `orm:"-"`
}

func TestOrm(t *testing.T) {
	u := UserInfo{
		UserName: "ggg",
		PassWord: "123",
		Age:      10,
		Ext:      "22",
	}
	registerModel(&u)

	db := store.Db
	o := NewOrm(db)

	id, err := o.Insert(&u)
	if err != nil {
		panic(err)
	}
	log.Println(id)
}

func TestOrmQuery(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	u := UserInfo{
		Id: 7,
	}
	registerModel(&u)

	db := store.Db
	o := NewOrm(db)

	err := o.Read(&u)
	if err != nil {
		panic(err)
	}

	log.Println(u)
}

func TestOrmQuery1(t *testing.T) {
	u := UserInfo{
		Id:       10,
		UserName: "ggg",
		Age:      11,
	}
	registerModel(&u)

	db := store.Db
	o := NewOrm(db)

	err := o.Read(&u, "id", "user_name")
	if err != nil {
		panic(err)
	}

	log.Println(u)
}
