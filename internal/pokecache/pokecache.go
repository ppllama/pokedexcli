package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt	time.Time
	val			[]byte
}

type Cache struct{
	cache	map[string]cacheEntry
	mu		sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: 	time.Now().UTC(),
		val:		val,
	}
	c.cache[key] = newEntry
	fmt.Printf("added %s to cache\n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	if !ok {
		return nil, ok
	}

	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C{
		c.reap(time.Now().UTC(), interval)
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: 	make(map[string]cacheEntry),
		mu: 	sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    for url, data := range c.cache {
        if data.createdAt.Before(now.Add(-interval)) {
            delete(c.cache, url)
        }
    }
}
