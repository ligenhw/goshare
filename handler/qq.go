package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/session"
)

// QqLogin qq login
func QqLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	var ses session.Store
	if ses, err = session.Instance.SessionStart(w, r); err != nil {
		handleError(err, w)
		return
	}

	defer r.Body.Close()

	var req OAuthReq
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(err, w)
		return
	}

	var uid int
	if uid, err = auth.QQLogin(req.Code); err != nil {
		handleError(err, w)
		return
	}

	ses.Set("userID", uid)
	return
}
