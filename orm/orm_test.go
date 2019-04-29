package orm

import (
	"testing"
	"time"

	"github.com/ligenhw/goshare/store"
)

// avoid import cycle
type User struct {
	Id       int       `json:"id"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
	Time     time.Time `json:"time"`
}

var (
	o = NewOrm(store.Db)
)

func init() {
	RegisterModel(new(User))
}

func TestOrm(t *testing.T) {
	// create
	u := User{
		UserName: "testorm",
		Password: "testorm_pass",
		Time:     time.Now(),
	}

	if id, err := o.Insert(&u); err != nil {
		t.Fatal(err)
		t.FailNow()
	} else {
		u.Id = int(id)
	}

	// query
	u2 := User{
		UserName: u.UserName,
	}
	if err := o.Read(&u2, "user_name"); err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	if u.Password != u2.Password {
		t.Fatal("query failed.")
		t.FailNow()
	}

	// update
	u.Password = "changeme"
	if _, err := o.Update(&u); err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	// delete
	if _, err := o.Delete(&u); err != nil {
		t.Fatal(err)
		t.FailNow()
	}
}

func TestOrmAll(t *testing.T) {
	var users []*User

	qs := o.QueryTable(new(User))
	num, err := qs.Filter("user_name", "ggg").All(&users)
	t.Log(num, err)
	if err == nil {
		for _, user := range users {
			t.Log(*user)
		}
	}
}
