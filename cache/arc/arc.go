package arc

import (
	"container/list"
	"sync"
)

type ARC struct {
	mtx      sync.Mutex
	capacity int
	t1       *list.List
	b1       *list.List
	t2       *list.List
	b2       *list.List
	cache    map[interface{}]*list.Element
}

func NewARC(capacity int) *ARC {
	return &ARC{
		capacity: capacity,
		t1:       list.New(),
		b1:       list.New(),
		t2:       list.New(),
		b2:       list.New(),
		cache:    make(map[interface{}]*list.Element),
	}
}

// Get returns the item from the cache.
func (c *ARC) Get(item any) any {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if elem, found := c.cache[item]; found {
		c.t2.MoveToFront(elem)
		return elem.Value
	}

	return nil
}

func (c *ARC) Put(item any) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.capacity == 0 {
		return
	}

	if elem, found := c.cache[item]; found {
		elem.Value = item
		c.t2.MoveToFront(elem)
		return
	}

	// not found, so add it to the cache
	if c.t1.Len()+c.t2.Len() == c.capacity {
		if c.t1.Len() == c.capacity {
			c.removeLast(c.b1)
		} else {
			c.removeLast(c.t1)
		}
	} else if c.t1.Len()+c.b1.Len()+c.t2.Len()+c.b2.Len() >= c.capacity {
		if c.t1.Len()+c.b1.Len()+c.t2.Len()+c.b2.Len() == 2*c.capacity {
			c.removeLast(c.b2)
		} else {
			c.removeLast(c.t2)
		}
	}

	// add the new item to the cache
	elem := c.t1.PushFront(item)
	c.cache[item] = elem
}

// removeLast removes the last element from the list.
func (c *ARC) removeLast(l *list.List) {
	if l.Len() == 0 {
		return
	}

	elem := l.Back()
	l.Remove(elem)
	delete(c.cache, elem)
}
