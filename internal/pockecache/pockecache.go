package pockecache

import (
	"sync"
	"time"
)

type Cache struct {
	values   map[string]cacheEntry
	duration time.Duration
	mu       *sync.RWMutex
}

func NewCache(duration time.Duration) Cache {
	ticker := time.Tick(duration)

	new_cache := Cache{
		duration: duration,
		values:   map[string]cacheEntry{},
		mu:       &sync.RWMutex{},
	}
	go new_cache.reapLoop(ticker)

	return new_cache
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (ch Cache) Add(key string, slice []byte) {
	new_entry := cacheEntry{
		createdAt: time.Now(),
		val:       slice,
	}
	ch.values[key] = new_entry
}

func (ch Cache) Get(key string) ([]byte, bool) {
	entry, find := ch.values[key]

	return entry.val, find
}

func (ch Cache) reapLoop(timeTicker <-chan time.Time) error {
	<-timeTicker
	ch.mu.Lock()
	for key, val := range ch.values {
		if time.Now().Sub(val.createdAt) >= ch.duration {
			delete(ch.values, key)
		}
	}
	ch.mu.Unlock()
	return nil
}
