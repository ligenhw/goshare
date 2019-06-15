package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ligenhw/goshare/handler/context"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/session"
	"github.com/ligenhw/goshare/user"
)

// CreateUser sign up
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	u := user.User{}
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		handleError(err, w)
		return
	}
	err = u.Create()
	handleError(err, w)
	return
}

// GetUser get current user info, through cookie -> session
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := *context.UserID(r)
	user := user.User{Id: userID}
	err := user.QueryByID()
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	handleError(err, w)
	return
}

// Login add the key 'userID' in current session
func Login(w http.ResponseWriter, r *http.Request) {
	var err error

	decoder := json.NewDecoder(r.Body)
	var user user.User
	if err = decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, err = auth.Check(user.UserName, user.Password); err != nil {
		handleError(err, w)
		return
	}

	ses, err := session.Instance.SessionStart(w, r)
	if err != nil {
		handleError(err, w)
		return
	}
	ses.Set("userID", user.Id)
}

// Logout remove the key 'userID' in current session
// keep the session
func Logout(w http.ResponseWriter, r *http.Request) {
	ses, err := session.Instance.SessionExist(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = ses.Delete("userID"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
