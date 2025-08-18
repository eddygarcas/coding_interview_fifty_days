// Package main implements a simple LRU (Least Recently Used) cache in Go.
// An LRU cache is a data structure that maintains a fixed number of items,
// discarding the least recently used items when the cache reaches its capacity.
package main

import (
	"container/list"
	"fmt"
	"time"
)

// LRUCache represents a Least Recently Used cache with a fixed capacity.
// It stores key-value pairs and evicts the least recently used entry when full.
// Implementation uses Go's container/list package for O(1) operations.
type LRUCache struct {
	capacity int                   // Maximum number of entries the cache can hold
	cache    *list.List            // Doubly linked list to track usage order
	items    map[int]*list.Element // Map from key to list element for O(1) lookups
}

// entry represents a key-value pair stored in the LRUCache.
// Each entry is stored as a value in the linked list element.
type entry struct {
	key   int // The key used to access this entry
	value int // The value associated with the key
}

// NewLRUCache creates and returns a new LRUCache with the specified capacity.
// Parameters:
//   - capacity: The maximum number of key-value pairs the cache can hold
//
// Returns:
//   - A pointer to a newly initialized LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    list.New(),
		items:    make(map[int]*list.Element, capacity), // Define map capacity, this wil avoid rehash and improve performance
	}
}

// Put inserts or updates a key-value pair in the cache.
// If the key already exists, it updates the value and marks the entry as most recently used.
// If the key doesn't exist and the cache is at capacity, it evicts the least recently used entry.
//
// Parameters:
//   - key: The key to store
//   - value: The value to associate with the key
//
// Time Complexity: O(1)
func (c LRUCache) Put(key, value int) {
	// If key exists, update its value and move it to front (most recently used)
	if val, ok := c.items[key]; ok {
		c.cache.MoveToFront(val)
		val.Value.(*entry).value = value
	} else {
		// If cache is at capacity, remove least recently used item (from back)
		if c.cache.Len() == c.capacity {
			lru := c.cache.Back()
			if lru != nil {
				c.cache.Remove(lru)
				delete(c.items, lru.Value.(*entry).key)
			}
		}
		// Add new entry to front of list (most recently used)
		item := &entry{value: value, key: key}
		c.items[key] = c.cache.PushFront(item)
	}
}

// Get retrieves the value for a given key and marks it as most recently used.
// Returns the value and true if found; 0 and false otherwise.
//
// Parameters:
//   - key: The key to look up
//
// Returns:
//   - value: The value associated with the key (0 if not found)
//   - ok: Boolean indicating whether the key was found
//
// Time Complexity: O(1)
func (c LRUCache) Get(key int) (value int, ok bool) {
	if val, ok := c.items[key]; ok {
		// Move accessed item to front (most recently used)
		c.cache.MoveToFront(val)
		return val.Value.(*entry).value, ok
	}
	return 0, false
}

// main demonstrates usage of the LRUCache by inserting and accessing elements.
// It shows how the cache behaves when items are added, updated, and accessed,
// including eviction of least recently used items when capacity is reached.
func main() {
	// Create a new LRU cache with capacity of 4
	cache := NewLRUCache(4)
	start := time.Now()

	// Add initial items to the cache
	cache.Put(1, 100)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(2, 200)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(3, 300)
	fmt.Printf("LRU Cache: %v\n", cache)

	// Update an existing item
	cache.Put(3, 3000)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(4, 400)
	fmt.Printf("LRU Cache: %v\n", cache)

	// Multiple updates to same key (demonstrates MRU behavior)
	cache.Put(4, 400)
	fmt.Printf("LRU Cache: %v\n", cache)
	cache.Put(4, 400)
	fmt.Printf("LRU Cache: %v\n", cache)
	cache.Put(4, 400)
	fmt.Printf("LRU Cache: %v\n", cache)
	cache.Put(4, 400)
	fmt.Printf("LRU Cache: %v\n", cache)

	// Add a new item when cache is full (should evict least recently used)
	cache.Put(5, 500)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(6, 600)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(6, 6000)
	fmt.Printf("LRU Cache: %v\n", cache)

	cache.Put(7, 7000)
	fmt.Printf("LRU Cache: %v\n", cache)

	// Try to access an item that may have been evicted
	if value, ok := cache.Get(1); ok {
		fmt.Printf("LRU Cache value: %v\n", value)
	} else {
		fmt.Printf("LRU Cache %d not found\n", 1)
	}
	fmt.Printf("LRU Cache: %v\n", cache)

	// Print elapsed time for performance measurement
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
