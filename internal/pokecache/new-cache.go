package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) (cache Cache) {
	cache = Cache{
		mu:       &sync.Mutex{},
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}

	go cache.reapLoop()

	return
}
