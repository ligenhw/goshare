package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/session"
)

// AlipayLogin github login
func AlipayLogin(w http.ResponseWriter, r *http.Request) {
	session, err := session.Instance.SessionStart(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer r.Body.Close()

	var req OAuthReq
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
