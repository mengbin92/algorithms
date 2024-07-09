package fifo

import (
	"fmt"
	"testing"
)

func TestFIFOCache(t *testing.T) {
	fifo := NewFIFOCache(2)
	fifo.Put(1)
	fifo.Put(2)
	value := fifo.Get(1)
	if value != nil {
		fmt.Println(value) // 输出: 1
	} else {
		fmt.Println("Not found")
	}

	// 淘汰键 1
	fifo.Put(3)
	value = fifo.Get(1)
	// 输出: Not found
	if value != nil {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}
	// 输出: 2
	value = fifo.Get(2)
	if value != nil {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	// 输出: 3
	value = fifo.Get(3)
	if value != nil {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}
}
