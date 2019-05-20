package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/ligenhw/goshare/orm"
	"github.com/ligenhw/goshare/store"
)

var db = store.Db

type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatarurl"`
	Time      time.Time `json:"time"`
	Profile   Profile   `orm:"-"`
}

var (
	o = orm.NewOrm(store.Db)
)

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) Create() (err error) {
	u.Time = time.Now()
	var id int64
	id, err = o.Insert(u)
	if err != nil {
		return
	}
	u.Id = int(id)
	return
}

// delete by Id
func (u *User) Delete() (err error) {
	_, err = o.Delete(u)
	return
}

// update by id
func (u *User) Update() (err error) {
	var columes []string
	var args []interface{}
	if u.UserName != "" {
		columes = append(columes, "user_name = ?")
		args = append(args, u.UserName)
	}
	if u.Password != "" {
		columes = append(columes, "password = ?")
		args = append(args, u.Password)
	}
	if len(columes) > 0 {
		sql := strings.Join(columes, ",")
		args = append(args, strconv.Itoa(u.Id))
		update := "UPDATE user SET " + sql + " WHERE id = ?"
		_, err = db.Exec(update, args...)
		return
	} else {
		return nil
	}
}

// query by UserName
func (u *User) QueryByName() (err error) {
	err = o.Read(u, "user_name")
	return
}

func (u *User) QueryByID() (err error) {
	err = o.Read(u)
	return
}

func GetAllUser() (users []*User, err error) {
	qs := o.QueryTable(new(User))
	_, err = qs.All(&users)
	return
}

func QueryUserWithIds(id ...interface{}) (users []*User, err error) {
	qs := o.QueryTable(new(User))
	_, err = qs.In("id", id...).All(&users)

	return
}
