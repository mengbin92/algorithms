package lfu

import (
	"container/list"
	"sync"
)

type entry struct {
	key   any
	value any
	freq  int
}

type LFUCache struct {
	mtx       sync.Mutex // protects the cache
	capacity  int
	size      int
	minFreq   int
	cache     map[any]*list.Element
	frequency map[int]*list.List
}

// NewLFUCache creates a new LFU cache
func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity:  capacity,
		cache:     make(map[any]*list.Element),
		frequency: make(map[int]*list.List),
	}
}

// Get retrieves a value from the cache
func (c *LFUCache) Get(key any) any {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if elem, found := c.cache[key]; found {
		c.incrementFrequency(elem)
		return elem.Value.(*entry).value
	}
	return nil
}

// Put inserts or updates a value in the cache
func (c *LFUCache) Put(key, value any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.capacity == 0 {
		return
	}

	if elem, found := c.cache[key]; found {
		elem.Value.(*entry).value = value
		c.incrementFrequency(elem)
	} else {
		if c.size == c.capacity {
			c.evict()
		}
		newEntry := &entry{key, value, 1}
		if c.frequency[1] == nil {
			c.frequency[1] = list.New()
		}
		elem := c.frequency[1].PushFront(newEntry)
		c.cache[key] = elem
		c.minFreq = 1
		c.size++
	}
}

// incrementFrequency increases the frequency of a cache entry
func (c *LFUCache) incrementFrequency(elem *list.Element) {
	e := elem.Value.(*entry)
	oldFreq := e.freq
	e.freq++

	c.frequency[oldFreq].Remove(elem)
	if c.frequency[oldFreq].Len() == 0 {
		delete(c.frequency, oldFreq)
		if c.minFreq == oldFreq {
			c.minFreq++
		}
	}

	if c.frequency[e.freq] == nil {
		c.frequency[e.freq] = list.New()
	}
	newElem := c.frequency[e.freq].PushFront(e)
    c.cache[e.key] = newElem
}

// evict removes the least frequently used cache entry
func (c *LFUCache) evict() {
	list := c.frequency[c.minFreq]
	elem := list.Back()
	if elem != nil {
		list.Remove(elem)
		delete(c.cache, elem.Value.(*entry).key)
		c.size--
	}
}
