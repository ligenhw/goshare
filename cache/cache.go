package cache

import (
	"container/list"
	"log"
)

type Cache interface {
	Put(k, v interface{})
	Get(k interface{}) interface{}
}

type LruCache struct {
	caches map[interface{}]*list.Element
	list   *list.List
	size   int
}

// like entry<k,v> in java
type Item struct {
	v interface{}
	k interface{}
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		caches: make(map[interface{}]*list.Element),
		list:   list.New(),
		size:   size,
	}
}

var _ Cache = new(LruCache)

func (m *LruCache) Put(k, v interface{}) {
	if e, ok := m.caches[k]; ok {
		m.update(e)
		return
	}

	e := m.list.PushFront(Item{v: v, k: k})
	m.caches[k] = e

	log.Println(m.list.Len())
	if m.list.Len() > m.size {
		m.removeOldestElement()
	}
}

func (m *LruCache) Get(k interface{}) interface{} {
	if e, ok := m.caches[k]; ok {
		m.update(e)
		return e.Value.(Item).v
	} else {
		return nil
	}
}

func (m *LruCache) update(e *list.Element) {
	m.list.MoveToFront(e)
}

func (m *LruCache) removeOldestElement() {
	e := m.list.Back()
	m.list.Remove(e)
	delete(m.caches, e.Value.(Item).k)
}
