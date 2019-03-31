package session

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Store contains all data for one session process with specific id.
type Store interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

type Provider interface {
	SessionInit(gclifetime int64) error
	SessionRead(sid string) (Store, error)
	SessionExist(sid string) bool
	SessionDestroy(sid string) error
	SessionGC()
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}

	provides[name] = provider
}

type Manager struct {
	provider Provider

	CookieName     string
	CookieLifeTime int
	Gclifetime     int64
	Maxlifetime    int64
}

var Instance *Manager
var instanceLock sync.Mutex = sync.Mutex{}

func NewManager(provideName string) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("type %s no register, do you forget import ?", provideName)
	}
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if Instance == nil {
		defaultTime := int64(60)
		provider.SessionInit(defaultTime)

		Instance = &Manager{
			provider:       provider,
			CookieName:     "sessionId",
			CookieLifeTime: 3600,
			Maxlifetime:    defaultTime,
			Gclifetime:     defaultTime,
		}
	}
	return Instance, nil
}

func (manager *Manager) GetProvider() Provider {
	return manager.provider
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Store, err error) {
	sid, err := manager.getSid(r)
	if err != nil {
		return nil, err
	}

	if sid != "" && manager.provider.SessionExist(sid) {
		return manager.provider.SessionRead(sid)
	}

	sid, err = manager.sessionID()
	if err != nil {
		return nil, err
	}

	session, err = manager.provider.SessionRead(sid)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     manager.CookieName,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	if manager.CookieLifeTime > 0 {
		cookie.MaxAge = manager.CookieLifeTime
		// cookie.Expires = time.Now().Add(time.Duration(manager.CookieLifeTime) * time.Second)
	}

	http.SetCookie(w, cookie)
	return
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	sid := url.QueryEscape(cookie.Value)
	manager.provider.SessionDestroy(sid)

	cookie = &http.Cookie{
		Name:     manager.CookieName,
		Path:     "/",
		HttpOnly: true,
		// Expires:  time.Now(),
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}

// GC Start session gc process.
// it can do gc in times after gc lifetime.
func (manager *Manager) GC() {
	manager.provider.SessionGC()
	time.AfterFunc(time.Duration(manager.Gclifetime)*time.Second, func() { manager.GC() })
}

func (manager *Manager) getSid(r *http.Request) (string, error) {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return "", nil
	}

	return url.QueryEscape(cookie.Value), nil
}

func (manager *Manager) sessionID() (string, error) {
	b := make([]byte, 16)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", fmt.Errorf("Could not successfully read from the system CSPRNG")
	}
	return hex.EncodeToString(b), nil
}
