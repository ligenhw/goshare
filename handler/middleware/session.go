package middleware

import (
	"log"
	"net/http"

	"github.com/ligenhw/goshare/auth"
	"github.com/ligenhw/goshare/handler/context"

	"github.com/ligenhw/goshare/session"
)

// CheckSession if not response 403
// set session to context
func CheckSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		globalSession := session.Instance
		session, err := globalSession.SessionExist(r)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Println("CheckSession ", session)
		r = context.SetSession(r, session)

		next.ServeHTTP(w, r)
	}
}

// CheckUser first check session , then check userId in session
// set userId to context
func CheckUser(next http.HandlerFunc) http.HandlerFunc {
	return CheckSession(func(w http.ResponseWriter, r *http.Request) {
		session := context.Session(r)
		log.Println("CheckUser ", session)
		if session == nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		userID, err := auth.Auth(session)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		r = context.SetUserID(r, userID)

		next.ServeHTTP(w, r)
	})
}
