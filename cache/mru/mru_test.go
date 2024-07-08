package mru

import (
	"fmt"
	"testing"
)

func TestMRU(t *testing.T) {
	cache := NewMRUCache(2)
	cache.Put(1)
	cache.Put(2)
	fmt.Println(cache.Get(1)) // Output: 1
	cache.Put(3)              // Removes key 2
	fmt.Println(cache.Get(2)) // Output: nil
	cache.Put(4)              // Removes key 3
	fmt.Println(cache.Get(1)) // Output: 1
	fmt.Println(cache.Get(3)) // Output: nil
	fmt.Println(cache.Get(4)) // Output: 4
}
