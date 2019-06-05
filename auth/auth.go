package auth

import (
	"errors"
	"net/http"

	"github.com/ligenhw/goshare/session"

	"github.com/ligenhw/goshare/user"
)

var (
	errorInput      = errors.New("username or password is wrong")
	ErrorAuthFailed = errors.New("auth failed")
)

func check(username, password string) (u user.User, err error) {
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

func Login(username, password string, session session.Store) (err error) {
	var u user.User
	if u, err = check(username, password); err == nil {
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

func GetAuthUser(w http.ResponseWriter, r *http.Request) (userID int, err error) {
	var s session.Store
	s, err = session.Instance.SessionStart(w, r)
	if err != nil {
		return
	}

	userID, err = Auth(s)
	return
}
