package auth

import (
	"errors"

	"github.com/ligenhw/goshare/session"

	"github.com/ligenhw/goshare/user"
)

var ErrorInput = errors.New("username or password is wrong")
var ErrorAuthFailed = errors.New("auth failed")

func Check(username, password string) (u user.User, err error) {
	u = user.User{UserName: username}
	err = u.Query()
	if err != nil {
		return
	}

	if u.Password == password {
		return
	}

	err = ErrorInput
	return
}

func Login(username, password string, session session.Store) (err error) {
	var u user.User
	if u, err = Check(username, password); err == nil {
		session.Set("userID", u.Id)
	}

	return
}

func Logout(session session.Store) (err error) {
	err = session.Delete("userID")
	return
}

func Auth(session session.Store) (userID int, err error) {
	value := session.Get("userID")
	if value == nil || value.(int) == 0 {
		err = ErrorAuthFailed
		return
	}

	userID = value.(int)
	return
}
