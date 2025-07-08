package pokecache

import (
	"maps"
	"sync"
	"time"
)

type Cache struct {
	lock        *sync.RWMutex
	entrys      map[string]CacheEntry
	duration    time.Duration
	lastReaping time.Time
	nextReaping time.Time
}

func NewCache(duration time.Duration) Cache {
	t := time.Now()
	return Cache{
		lock:        &sync.RWMutex{},
		entrys:      make(map[string]CacheEntry),
		duration:    duration,
		lastReaping: t,
		nextReaping: t.Add(duration),
	}
}

func (c *Cache) reaping(t time.Time) {

	maps.DeleteFunc(c.entrys, func(_ string, entry CacheEntry) bool { return entry.created.Before(c.lastReaping) })

	c.lastReaping = t
	c.nextReaping = t.Add(c.duration)
}

func (c *Cache) Add(key string, val []byte) {

	c.lock.Lock()
	defer c.lock.Unlock()

	t := time.Now()

	if t.After(c.nextReaping) {
		c.reaping(t)
	}

	c.entrys[key] = CacheEntry{
		val:     val,
		created: t,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	entry, ok := c.entrys[key]
	return entry.val, ok
}
