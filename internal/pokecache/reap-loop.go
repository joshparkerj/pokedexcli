package pokecache

import (
	"time"
)

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, entry := range c.entries {
		if time.Since(entry.createdAt) > c.interval {
			delete(c.entries, k)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap()
	}
}
