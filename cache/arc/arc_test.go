package arc

import "testing"

func TestARC(t *testing.T) {
	cache := NewARC(2)

	cache.Put("a")
	cache.Put("b")

	val := cache.Get("a")
	if val != nil {
		println("Value for 'a':", val)
	} else {
		println("'a' not found")
	}

	cache.Put("c")

	val = cache.Get("b")
	if val != nil {
		println("Value for 'b':", val)
	} else {
		println("'b' not found")
	}
}
