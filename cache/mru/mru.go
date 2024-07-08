package mru

import (
	"container/list"
	"sync"
)

// MRUCache represents a Most Recently Used (MRU) cache.
type MRUCache struct {
	mtx      sync.Mutex
	capacity int
	cache    map[any]*list.Element
	list     *list.List
}

// NewMRUCache creates a new MRUCache with the given capacity.
func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[any]*list.Element),
		list:     list.New(),
	}
}

// Get returns the item from the cache.
// This function is safe for concurrent access.
func (c *MRUCache) Get(item any) any {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	node, exists := c.cache[item]
	if exists {
		return node.Value
	} else {
		return nil
	}
}

// Add adds the given item to the cache.
// This function is safe for concurrent access.
func (c *MRUCache) Put(item any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// if capacity is 0, nothing can be added, so just return
	if c.capacity == 0 {
		return
	}

	// check if the item is already in the cache
	if node, exists := c.cache[item]; exists {
		node.Value = item
		c.list.MoveToFront(node)
		return
	}

	// if the cache is full, remove the front element
	if c.list.Len() == c.capacity {
		elem := c.list.Front()
		c.list.Remove(elem)
		delete(c.cache, elem.Value)
	}

	// add the new item to the back of the list
	node := c.list.PushFront(item)
	c.cache[item] = node
}
