package session

import (
	"net/http"
	"sync"
	"time"
)

type SessionMaager struct {
	CookieName string
	ExpireTime time.Time
	Sessions   map[string]Session
	lock       sync.RWMutex
}

type Session interface {
	Get(k string) (interface{}, error)
	Put(k string, v interface{}) error
}

func NewSessionManager() *SessionMaager {
	return &SessionMaager{
		CookieName: "sessionId",
		lock:       sync.RWMutex{},
	}
}

func (sm *SessionMaager) Session(r *http.Request) Session {
	if c, err := r.Cookie(sm.CookieName); err != nil {
		// no session id
		return nil
	} else {
		if v, ok := sm.Sessions[c.Value]; ok {
			return v
		} else {
			// invalidate session id
			return nil
		}
	}
}
