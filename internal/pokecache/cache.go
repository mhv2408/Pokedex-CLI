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
	entry_map map[string]cacheEntry
	mu        *sync.Mutex
}

func NewCache(duration time.Duration) Cache {

	new_cache := Cache{
		entry_map: make(map[string]cacheEntry),
		mu:        &sync.Mutex{},
	}
	go new_cache.reapLoop(duration)
	return new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cache_entry := cacheEntry{
		createdAt: time.Time{},
		val:       val,
	}
	c.entry_map[key] = cache_entry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if res, ok := c.entry_map[key]; ok {
		return res.val, true
	}

	return nil, false
}
func (c *Cache) reapLoop(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			c.mu.Lock()
			for k, v := range c.entry_map {
				if time.Since(v.createdAt) > interval {
					delete(c.entry_map, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
