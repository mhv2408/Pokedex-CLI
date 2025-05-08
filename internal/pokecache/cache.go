package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(duration time.Duration) Cache {

	new_cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go new_cache.reapLoop(duration)
	return new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Time{},
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if res, ok := c.cache[key]; ok {
		return res.val, true
	}

	return nil, false
}
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) > interval {
				delete(c.cache, k)
			}
		}
	}
}
