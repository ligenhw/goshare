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
	UserName string `orm:"varchar(20);pk"`
	PassWord string `orm:"varchar(100)"`
	Age      int
	Time     time.Time `orm:"-"`
}

func TestOrm(t *testing.T) {
	u := UserInfo{
		UserName: "ggg",
		PassWord: "123",
		Age:      10,
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
	u
}