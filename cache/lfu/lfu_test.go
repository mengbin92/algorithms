package lfu

import (
	"fmt"
	"testing"
)

func TestLFU(t *testing.T) {
	lfu := NewLFUCache(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1)) // 输出: 1
	lfu.Put(3, 3)           // 淘汰键 2
	fmt.Println(lfu.Get(2)) // 输出: nil (未命中缓存)
	fmt.Println(lfu.Get(3)) // 输出: 3
	lfu.Put(4, 4)           // 淘汰键 1
	fmt.Println(lfu.Get(1)) // 输出: nil (未命中缓存)
	fmt.Println(lfu.Get(3)) // 输出: 3
	fmt.Println(lfu.Get(4)) // 输出: 4
}
