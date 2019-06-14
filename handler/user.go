package handler

import (
	"encoding/json"
	"net/http"

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
	var err error
	globalSession := session.Instance

	var session session.Store
	if session, err = globalSession.SessionStart(w, r); err != nil {
		handleError(err, w)
		return
	}

	var userID int
	if userID, err = auth.Auth(session); err != nil {
		handleError(err, w)
		return
	}

	user := user.User{Id: userID}
	if err = user.QueryByID(); err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(user)
	handleError(err, w)
	return
}

// Login add the key 'userID' in current session
func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var ses session.Store
	if ses, err = session.Instance.SessionStart(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	decoder := json.NewDecoder(r.Body)
	var user user.User
	if err = decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = auth.Login(user.UserName, user.Password, ses); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// Logout remove the key 'userID' in current session
func Logout(w http.ResponseWriter, r *http.Request) {
	ses, err := session.Instance.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = auth.Logout(ses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
