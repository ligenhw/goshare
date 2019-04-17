package api

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/ligenhw/goshare/auth"

	"github.com/ligenhw/goshare/session"
	"github.com/ligenhw/goshare/user"
)

// /user   need auth
func Get(w http.ResponseWriter, r *http.Request) (err error) {
	globalSession := session.Instance

	var session session.Store
	session, err = globalSession.SessionStart(w, r)
	if err != nil {
		return
	}

	var userID int
	userID, err = auth.Auth(session)
	if err != nil {
		return
	}

	user := user.User{Id: userID}
	err = user.QueryByID()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(user)
	return
}

// /user
func Post(w http.ResponseWriter, r *http.Request) (err error) {
	// TODO: close the body
	decoder := json.NewDecoder(r.Body)
	u := user.User{}
	err = decoder.Decode(&u)
	if err != nil {
		log.Println(err)
		return
	}
	err = u.Create()
	return
}

// /user/:id
func Delete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	u := user.User{Id: id}
	err = u.Delete()
	return
}

// /user
func Put(w http.ResponseWriter, r *http.Request) (err error) {
	decoder := json.NewDecoder(r.Body)
	u := user.User{}
	err = decoder.Decode(&u)
	if err != nil {
		return
	}

	err = u.Update()
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method)

	var err error
	switch r.Method {
	case http.MethodGet:
		err = Get(w, r)
	case http.MethodPost:
		err = Post(w, r)
	case http.MethodDelete:
		err = Delete(w, r)
	case http.MethodPut:
		err = Put(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/api/user/", UserHandler)
}
