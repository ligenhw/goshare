package redis

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

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

// SessionRelease save to redis and set the expire time.
func (rs *SessionStore) SessionRelease() {
	b, err := session.EncodeGob(rs.values)
	if err != nil {
		return
	}
	c := rs.p.Get()
	defer c.Close()
	if _, err := redis.Int(c.Do("SETEX", rs.sid, rs.maxlifetime, string(b))); err != nil {
		log.Println(err)
	}
}

// Provider redis session provider
type Provider struct {
	maxlifetime int64
	savePath    string
	poolsize    int
	password    string
	dbNum       int
	poollist    *redis.Pool
}

// SessionInit init redis session
// savepath like redis server addr,pool size,IdleTimeout second
// e.g. 127.0.0.1:6379,10,60
func (rp *Provider) SessionInit(maxlifetime int64, savePath string) error {
	rp.maxlifetime = maxlifetime
	configs := strings.Split(savePath, ",")
	if len(configs) > 0 {
		rp.savePath = configs[0]
	}
	if len(configs) > 1 {
		poolsize, err := strconv.Atoi(configs[1])
		if err != nil || poolsize < 0 {
			rp.poolsize = MaxPoolSize
		} else {
			rp.poolsize = poolsize
		}
	} else {
		rp.poolsize = MaxPoolSize
	}
	var idleTimeout time.Duration
	if len(configs) > 2 {
		timeout, err := strconv.Atoi(configs[2])
		if err == nil && timeout > 0 {
			idleTimeout = time.Duration(timeout) * time.Second
		}
	}
	rp.poollist = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rp.savePath)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		MaxIdle: rp.poolsize,
	}

	rp.poollist.IdleTimeout = idleTimeout

	return rp.poollist.Get().Err()
}

// SessionRead read redis session by sid
func (rp *Provider) SessionRead(sid string) (session.Store, error) {
	c := rp.poollist.Get()
	defer c.Close()

	var kv map[interface{}]interface{}

	kvs, err := redis.String(c.Do("GET", sid))
	if err != nil && err != redis.ErrNil {
		return nil, err
	}
	if len(kvs) == 0 {
		kv = make(map[interface{}]interface{})
	} else {
		if kv, err = session.DecodeGob([]byte(kvs)); err != nil {
			return nil, err
		}
	}

	rs := &SessionStore{p: rp.poollist, sid: sid, values: kv, maxlifetime: rp.maxlifetime}
	return rs, nil
}

// SessionExist check redis session exist by sid
func (rp *Provider) SessionExist(sid string) bool {
	c := rp.poollist.Get()
	defer c.Close()

	if existed, err := redis.Int(c.Do("EXISTS", sid)); err != nil || existed == 0 {
		return false
	}
	return true
}

// SessionDestroy delete redis session by id
func (rp *Provider) SessionDestroy(sid string) error {
	c := rp.poollist.Get()
	defer c.Close()

	c.Do("DEL", sid)
	return nil
}

// SessionGC Impelment method, no used.
func (rp *Provider) SessionGC() {
}

func init() {
	session.Register("redis", redispder)
}
