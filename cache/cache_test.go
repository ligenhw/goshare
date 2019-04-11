package cache

import (
	"log"
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewLruCache(10)
	for i := 0; i < 100; i++ {
		c.Put(i, "item "+strconv.Itoa(i))
	}

	c.Get(95)
	c.Get(95)

	for i := 0; i < 8; i++ {
		c.Put(i, "item - "+strconv.Itoa(i))
	}

	for i := 0; i < 100; i++ {
		v := c.Get(i)
		log.Println(i, v)
	}
}
