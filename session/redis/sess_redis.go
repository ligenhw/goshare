package redis

import (
	"sync"

	"github.com/ligenhw/goshare/session"

	"github.com/gomodule/redigo/redis"
)

var redispder = &Provider{}

// MaxPoolSize redis max pool size
var MaxPoolSize = 100

// SessionStore redis session store
type SessionStore struct {
	p           *redis.Pool
	sid         string
	lock        sync.RWMutex
	values      map[interface{}]interface{} // read and write buffer
	maxlifetime int64
}

// Set value in redis session
func (rs *SessionStore) Set(key, value interface{}) error {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	rs.values[key] = value
	return nil
}

// Get value in redis session
func (rs *SessionStore) Get(key interface{}) interface{} {
	rs.lock.RLock()
	defer rs.lock.RUnlock()
	if v, ok := rs.values[key]; ok {
		return v
	}
	return nil
}

// Delete key in session
func (rs *SessionStore) Delete(key interface{}) error {
	rs.lock.RLock()
	defer rs.lock.RUnlock()
	delete(rs.values, key)
	return nil
}

// SessionID get sid
func (rs *SessionStore) SessionID() string {
	return rs.sid
}

// --- provider ---

type Provider struct {
}

// SessionInit init memory session
func (rp *Provider) SessionInit(maxlifetime int64) error {
	// rp.maxlifetime = maxlifetime
	return nil
}

// SessionInit init memory session
func (rp *Provider) SessionRead(sid string) (session.Store, error) {
	// rp.maxlifetime = maxlifetime
	return nil, nil
}

// SessionExist init memory session
func (rp *Provider) SessionExist(sid string) bool {
	// rp.maxlifetime = maxlifetime
	return false
}

// SessionGC init memory session
func (rp *Provider) SessionDestroy(sid string) error {
	// rp.maxlifetime = maxlifetime
	return nil
}

// SessionGC Impelment method, no used.
func (rp *Provider) SessionGC() {
}

func init() {
	session.Register("redis", redispder)
}
