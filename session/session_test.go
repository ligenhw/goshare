package session

import (
	"testing"
)

func TestStore(t *testing.T) {
	globalSession, _ := NewManager("mem", "")
	go globalSession.GC()

	sid, _ := globalSession.sessionID()
	store, _ := globalSession.provider.SessionRead(sid)
	store.Set("userID", 1)

	r := store.Get("userID")
	t.Log(r.(int))
}
