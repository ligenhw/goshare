package session

import (
	"container/list"
	"log"
	"sync"
	"time"
)

var mempder = &MemProvider{list: list.New(), sessions: make(map[string]*list.Element)}

type MemSessionStore struct {
	sid          string                      //session id
	timeAccessed time.Time                   //last access time
	value        map[interface{}]interface{} //session store
	lock         sync.RWMutex
}

// Set value to memory session
func (st *MemSessionStore) Set(key, value interface{}) error {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.value[key] = value
	return nil
}

// Get value from memory session by key
func (st *MemSessionStore) Get(key interface{}) interface{} {
	st.lock.RLock()
	defer st.lock.RUnlock()
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete in memory session by key
func (st *MemSessionStore) Delete(key interface{}) error {
	st.lock.Lock()
	defer st.lock.Unlock()
	delete(st.value, key)
	return nil
}

// SessionID get this id of memory session store
func (st *MemSessionStore) SessionID() string {
	return st.sid
}

type MemProvider struct {
	lock        sync.RWMutex             // locker
	sessions    map[string]*list.Element // map in memory
	list        *list.List               // for gc
	maxlifetime int64
}

// SessionInit init memory session
func (pder *MemProvider) SessionInit(maxlifetime int64) error {
	pder.maxlifetime = maxlifetime
	return nil
}

func (pder *MemProvider) SessionRead(sid string) (session Store, err error) {
	pder.lock.RLock()
	if element, ok := pder.sessions[sid]; ok {
		go pder.SessionUpdte(sid)
		pder.lock.RUnlock()
		return element.Value.(*MemSessionStore), nil
	}
	pder.lock.RUnlock()

	pder.lock.Lock()
	newsess := &MemSessionStore{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        make(map[interface{}]interface{}),
	}
	element := pder.list.PushFront(newsess)
	pder.sessions[sid] = element
	pder.lock.Unlock()
	return newsess, nil
}

func (pder *MemProvider) SessionExist(sid string) bool {
	pder.lock.RLock()
	defer pder.lock.RUnlock()

	if _, ok := pder.sessions[sid]; ok {
		return true
	}

	return false
}

func (pder *MemProvider) SessionDestroy(sid string) error {
	pder.lock.Lock()
	pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		pder.list.Remove(element)
		delete(pder.sessions, sid)
	}

	return nil
}

func (pder *MemProvider) SessionGC() {
	pder.lock.RLock()
	for {
		element := pder.list.Back()
		if element == nil {
			break
		}

		if element.Value.(*MemSessionStore).timeAccessed.Unix()+pder.maxlifetime < time.Now().Unix() {
			pder.lock.RUnlock()
			pder.lock.Lock()
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*MemSessionStore).sid)
			log.Println("session gc : " + element.Value.(*MemSessionStore).sid)
			pder.lock.Unlock()
			pder.lock.RLock()
		} else {
			break
		}
	}
	pder.lock.RUnlock()
}

func (pder *MemProvider) SessionUpdte(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*MemSessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	Register("mem", mempder)
}
