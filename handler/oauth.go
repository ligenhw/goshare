package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ligenhw/goshare/auth"

	"github.com/ligenhw/goshare/session"
)

// OAuthReq oauth request body
type OAuthReq struct {
	Code string `json:"code"`
}

// OauthLogin used by github qq alipay
func OauthLogin(authFunc func(code string) (id int, err error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req OAuthReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			handleError(err, w)
			return
		}

		uid, err := authFunc(req.Code)
		if err != nil {
			handleError(err, w)
			return
		}

		ses, err := session.Instance.SessionStart(w, r)
		if err != nil {
			handleError(err, w)
			return
		}
		ses.Set("userID", uid)
	}
}

var (
	// GhLogin github oauth2.0
	GhLogin = OauthLogin(auth.GhLogin)
	// QqLogin qq oauth2.0
	QqLogin = OauthLogin(auth.QQLogin)
	// AlipayLogin alipay oauth2.0
	AlipayLogin = OauthLogin(auth.AlipayLogin)
)
