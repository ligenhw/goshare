package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ligenhw/goshare/auth"

	"github.com/ligenhw/goshare/session"
)

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	session, err := session.Instance.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	decoder := json.NewDecoder(r.Body)
	var user User
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = auth.Login(user.UserName, user.PassWord, session)
	log.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, err := session.Instance.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = auth.Logout(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/api/login/", handleLogin)
	http.HandleFunc("/api/logout/", handleLogout)
	http.HandleFunc("/api/ghlogin/", ghLoginHandler)
	http.HandleFunc("/api/qqlogin/", qqLoginHandler)
}
