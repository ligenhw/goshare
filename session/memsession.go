package session

import (
	"errors"
	"sync"
)

type MemSession struct {
	lock   sync.RWMutex
	Values map[string]interface{}
}

func (s *MemSession) Get(k string) (interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if v, ok := s.Values[k]; ok {
		return v, nil
	}

	return nil, errors.New("can not find k : " + k)
}

func (s *MemSession) Put(k string, v interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Values[k] = v
	return nil
}
