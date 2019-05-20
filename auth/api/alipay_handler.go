package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/session"
)

func alipayPostHandler(w http.ResponseWriter, r *http.Request) (err error) {
	session, err := session.Instance.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer r.Body.Close()

	var req Req
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		return
	}

	var uid int
	if uid, err = auth.AlipayLogin(req.Code); err != nil {
		return
	}

	session.Set("userID", uid)
	return
}

func alipayLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method)

	var err error
	switch r.Method {
	case http.MethodPost:
		err = alipayPostHandler(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
