package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/ligenhw/goshare/store"
)

var db = store.Db

type User struct {
	Id       int       `json:"id"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
	Time     time.Time `json:"time"`
	Profile  Profile   `json:"profile"`
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func (u *User) Create() (err error) {
	_, err = db.Exec("INSERT INTO user (user_name, password) VALUES (?, ?)", u.UserName, u.Password)
	return
}

// delete by Id
func (u *User) Delete() (err error) {
	_, err = db.Exec("DELETE FROM user where id = ?", u.Id)
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
func (u *User) Query() (err error) {
	err = db.QueryRow("SELECT id, user_name, password, time FROM user WHERE user_name = ?", u.UserName).Scan(&u.Id, &u.UserName, &u.Password, &u.Time)
	return
}

func GetAllUser() (users []*User, err error) {
	rows, err := db.Query("SELECT id, user_name, password, time FROM user")
	if err != nil {
		return
	}
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.UserName, &u.Password, &u.Time)
		users = append(users, &u)
	}
	return
}
