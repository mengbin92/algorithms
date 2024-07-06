package lru

import (
	"testing"
)

func TestCache(t *testing.T) {
	// Create fake data to be used in the test.
	numNonces := 10
	nonces := make([]int, 0, numNonces)
	for i := 0; i < numNonces*2; i++ {
		nonces = append(nonces, i)
	}

	tests := []struct {
		name     string
		capacity int
	}{
		{"capacity-0", 0},
		{"capacity-5", 5},
		{"capacity all available", numNonces},
		{"capacity not enough", numNonces * 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new LRU cache.
			lru := NewLRUCache(tt.capacity)

			// if tt.capacity == 0 {
			// 	return
			// }
			// Add all the nonces to the cache.
			for i := 0; i < tt.capacity; i++ {
				lru.Add(nonces[i])
			}

			// Check that all the nonces are in the cache.
			for i := 0; i < tt.capacity; i++ {
				if !lru.Contains(nonces[i]) {
					t.Errorf("expected nonce %d to be in the cache", nonces[i])
				}
				value := lru.Get(nonces[i])
				if value != nonces[i] {
					t.Errorf("expected nonce %d to have value %d, got %d", nonces[i], nonces[i], value)
				}
			}

			// Get a nonce that is not in the cache.
			value := lru.Get(tt.capacity + 1)
			if value != tt.capacity+1 {
				t.Errorf("expected nonce %d to not be in the cache", tt.capacity+1)
			}

			// Check that the cache is of the expected size.
			if lru.Len() != tt.capacity {
				t.Errorf("expected cache size to be %d, got %d", tt.capacity, lru.Len())
			}

			for i := 0; i < lru.Len(); i++ {
				// Delete the first nonce in the cache.
				lru.Delete(nonces[i])
				// Check that the nonce is not in the cache.
				if lru.Contains(nonces[i]) {
					t.Errorf("expected nonce %d to be removed from the cache", nonces[i])
				}
			}
		})
	}
}

// BenchmarkCache performs basic benchmarks on LRU cache.
func BenchmarkCache(b *testing.B) {
	// Create fake data to be used in the benchmark.
	b.StopTimer()
	numNonces := 1000000
	nonces := make([]int, 0, numNonces)
	for i := 0; i < numNonces; i++ {
		nonces = append(nonces, i)
	}

	// Run the benchmark.
	b.StartTimer()
	capacity := 2000
	lru := NewLRUCache(capacity)
	for i := 0; i < b.N; i++ {
		lru.Add(nonces[i%numNonces])
	}
}
