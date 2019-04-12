package orm

import (
	"testing"
	"time"

	"github.com/ligenhw/goshare/store"
)

type UserInfo struct {
	Id       int    `orm:"auto"`
	UserName string `orm:"varchar(20)"`
	PassWord string `orm:"varchar(100)"`
	Age      int
	Ext      string
	Time     time.Time `orm:"-"`
}

var (
	o = NewOrm(store.Db)
)

func init() {
	registerModel(new(UserInfo))
}

func TestOrm(t *testing.T) {

	u := UserInfo{
		UserName: "ggg",
		PassWord: "1234",
		Age:      10,
	}

	id, err := o.Insert(&u)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestOrmQuery(t *testing.T) {
	u := UserInfo{
		Id: 9,
	}

	err := o.Read(&u)
	if err != nil {
		t.Error(err)
	}

	t.Log(u)
}

func TestOrmQueryWithCols(t *testing.T) {
	u := UserInfo{
		Id:       7,
		UserName: "ggg",
		Age:      11,
	}

	err := o.Read(&u, "id", "user_name")
	if err != nil {
		t.Error(err)
	}

	t.Log(u)
}

func TestOrmDelete(t *testing.T) {
	u := UserInfo{Id: 9}
	num, err := o.Delete(&u)
	if err != nil {
		t.Error(err)
	}

	t.Log(num)
}

func TestOrmUpdate(t *testing.T) {
	u := UserInfo{
		Id: 8,
		UserName: "lll",
	}

	o.Update(&u)
}
