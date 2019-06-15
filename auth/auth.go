package auth

import (
	"errors"

	"github.com/ligenhw/goshare/session"

	"github.com/ligenhw/goshare/user"
)

var (
	errorInput      = errors.New("username or password is wrong")
	errorAuthFailed = errors.New("auth failed")
)

// Check password if correct return the user
func Check(username, password string) (u user.User, err error) {
	if password == "" {
		err = errorInput
		return
	}

	u = user.User{UserName: username}
	err = u.QueryByName()
	if err != nil {
		return
	}

	if u.Password == password {
		return
	}

	err = errorInput
	return
}

func Auth(session session.Store) (userID int, err error) {
	value := session.Get("userID")
	if value == nil || value.(int) == 0 {
		err = errorAuthFailed
		return
	}

	userID = value.(int)
	return
}
