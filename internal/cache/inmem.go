package cache

import (
	"sync"
	"time"
)

type item struct {
	value      string
	expiration int64 // unix ms
}

type InMemCache struct {
	data map[string]item
	mu   sync.RWMutex
	ttl  time.Duration
}

func NewInMemCache(ttl time.Duration) *InMemCache {
	c := &InMemCache{
		data: make(map[string]item),
		ttl:  ttl,
	}

	// auto cleanup goroutine
	go c.cleanup()

	return c
}

func (c *InMemCache) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = item{
		value:      value,
		expiration: time.Now().Add(c.ttl).UnixMilli(),
	}
}

func (c *InMemCache) Get(key string) (string, bool) {
	c.mu.RLock()
	itm, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		return "", false
	}

	if time.Now().UnixMilli() > itm.expiration {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return "", false
	}

	return itm.value, true
}

func (c *InMemCache) cleanup() {
	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		now := time.Now().UnixMilli()

		c.mu.Lock()
		for k, v := range c.data {
			if now > v.expiration {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}
