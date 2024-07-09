package fifo

import (
	"container/list"
	"sync"
)

type FIFOCache struct {
	mtx      sync.Mutex
	capacity int
	queue    *list.List
	cache    map[any]*list.Element
}

func NewFIFOCache(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity: capacity,
		queue:    list.New(),
		cache:    make(map[any]*list.Element),
	}
}

// Get returns the item from the cache.
func (c *FIFOCache) Get(item any) any {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if node, exists := c.cache[item]; exists {
		return node.Value
	}
	return nil
}

// Put adds the given item to the cache.
func (c *FIFOCache) Put(item any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// if capacity is 0, nothing can be added, so just return
	if c.capacity == 0 {
		return
	}

	// check if the item is already in the cache
	if elem, found := c.cache[item]; found {
		c.queue.MoveToBack(elem)
		return
	}

	// if the cache is full, remove the front element in queue
	if c.queue.Len() == c.capacity {
		c.evict()
	}

	// add the new item to the back of the list
	node := c.queue.PushBack(item)
	c.cache[item] = node
}

// evict removes the oldest entry from the cache
func (c *FIFOCache) evict() {
	elem := c.queue.Front()
	if elem != nil {
		c.queue.Remove(elem)
		delete(c.cache, elem.Value)
	}
}
