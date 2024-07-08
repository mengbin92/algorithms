package lru

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	mtx      sync.Mutex            // protects the cache
	capacity int                   // capacity of the cache
	cache    map[any]*list.Element // nearly O(1) lookups
	list     *list.List            // O(1) insert, update, delete
}

// NewLRUCache creates a new LRU cache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[any]*list.Element),
		list:     list.New(),
	}
}

// Contains checks if the given item is in the cache.
// This function is safe for concurrent access.
func (c *LRUCache) Contains(item any) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	node, exists := c.cache[item]
	if exists {
		c.list.MoveToFront(node)
	}
	return exists
}

// Get returns the item from the cache.
// This function is safe for concurrent access.
func (c *LRUCache) Get(item any) any {
	node, exists := c.cache[item]
	if exists {
		c.mtx.Lock()
		c.list.MoveToFront(node)
		c.mtx.Unlock()
		return node.Value
	} else {
		c.Put(item)
		return item
	}
}

// Add adds the given item to the cache.
// This function is safe for concurrent access.
func (c *LRUCache) Put(item any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// if capacity is 0, nothing can be added, so just return
	if c.capacity == 0 {
		return
	}

	// check if the item is already in the cache
	if node, exists := c.cache[item]; exists {
		c.list.MoveToFront(node)
		return
	}

	// if the cache is full, remove the last element
	if c.list.Len() == c.capacity {
		last := c.list.Back()
		c.list.Remove(last)
		delete(c.cache, last.Value)
	}

	// add the new item to the front of the list
	node := c.list.PushFront(item)
	c.cache[item] = node
}

// Delete removes the given item from the cache if exists.
// This function is safe for concurrent access.
func (c *LRUCache) Delete(item any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// check if the item is already in the cache
	if node, exists := c.cache[item]; exists {
		c.list.Remove(node)
		delete(c.cache, item)
	}
}

// Len returns the number of items in the cache.
// This function is safe for concurrent access.
func (c *LRUCache) Len() int {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.list.Len()
}
