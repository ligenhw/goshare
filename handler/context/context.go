package context

import (
	"context"
	"net/http"

	"github.com/ligenhw/goshare/session"
)

type contextKey int

const (
	uid contextKey = iota
	ses
)

// UserID returns the route variables for the current request, if any.
func UserID(r *http.Request) *int {
	if rv := contextGet(r, uid); rv != nil {
		return rv.(*int)
	}
	return nil
}

// SetUserID set user id
func SetUserID(r *http.Request, val int) *http.Request {
	return contextSet(r, uid, &val)
}

// Session get session from context
func Session(r *http.Request) session.Store {
	if rv := contextGet(r, ses); rv != nil {
		return rv.(session.Store)
	}
	return nil
}

// SetSession set session to context
func SetSession(r *http.Request, val session.Store) *http.Request {
	return contextSet(r, ses, val)
}

func contextGet(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

func contextSet(r *http.Request, key, val interface{}) *http.Request {
	if val == nil {
		return r
	}

	return r.WithContext(context.WithValue(r.Context(), key, val))
}
