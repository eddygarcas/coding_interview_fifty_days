// Package main implements a simple LRU (Least Recently Used) cache in Go.
package main

import (
	"fmt"
	"time"
)

// LRUCache represents a Least Recently Used cache with a fixed capacity.
// It stores key-value pairs and evicts the least recently used entry when full.
type LRUCache struct {
	capacity int           // Maximum number of entries the cache can hold
	cache    map[int]entry // Map from key to cache entry
}

// entry represents a value stored in the LRUCache, along with its usage index.
type entry struct {
	value int // The value associated with the key
	index int // The usage index (0 = most recently used; higher = less recent)
}

// NewLRUCache creates and returns a new LRUCache with the specified capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]entry, capacity),
	}
}

// Put inserts or updates a key-value pair in the cache.
// It marks the entry as most recently used (index 0),
// increments the index of all other entries, and evicts the least recently used if needed.
func (c LRUCache) Put(key int, value int) {
	for k, e := range c.cache {
		if e.index < c.capacity {
			e.index++
			c.cache[k] = e
		} else {
			delete(c.cache, k)
		}
	}
	c.cache[key] = entry{value, 0}
}

// Get retrieves the value for a given key and marks it as most recently used.
// Returns (value, true) if found; (0, false) otherwise.
func (c LRUCache) Get(key int) (value int, ok bool) {
	if e, ok := c.cache[key]; ok {
		c.Put(key, e.value)
		return e.value, ok
	}
	return 0, false
}

// main demonstrates usage of the LRUCache by inserting and accessing elements.
func main() {
	cache := NewLRUCache(4)
	start := time.Now()
	cache.Put(1, 100)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(2, 200)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(3, 300)
	fmt.Printf("LRU Cache: %v\n", cache)
	cache.Put(3, 3000)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(4, 400)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(5, 500)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(6, 600)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(6, 6000)
	fmt.Printf("LRU Cache: %v\n", cache)
	cache.Put(7, 7000)
	fmt.Printf("LRU Cache: %v\n", cache)
	if value, ok := cache.Get(1); ok {
		fmt.Printf("LRU Cache value: %v\n", value)
	} else {
		fmt.Printf("LRU Cache %d not found\n", 1)
	}
	fmt.Printf("LRU Cache: %v\n", cache)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
