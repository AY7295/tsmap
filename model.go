package tsmap

import (
	"sync"
	"sync/atomic"
)

// TSMap is a thread-safe map with specific key/value type.
type TSMap[K any, V any] interface {
	// Get returns the value associated with the given key and whether it was
	Get(key K) (V, bool)
	// Set sets the value associated with the given key.
	Set(key K, value V)
	// Delete deletes the value associated with the given key.
	Delete(key K)
	// Len returns the number of items in the map.
	Len() int
	// Range calls fn sequentially for each key and value present in the map.
	//	but if exception exists and return true, fn will not be called.
	Range(fn Operation[K, V], exception ...Exception[K, V])
}

type Exception[K, V any] func(key K, value V) bool
type Operation[K, V any] func(key K, value V)

// New returns a new TSMap[K, V] that uses sync.Map for synchronization.
func New[K comparable, V any]() TSMap[K, V] {
	return NewSync[K, V]()
}

// NewMutex returns a new TSMap[K, V] that uses map[K]V and a mutex for synchronization.
func NewMutex[K comparable, V any]() TSMap[K, V] {
	return &mutexMap[K, V]{
		mp: make(map[K]V),
	}
}

// NewSync returns a new TSMap[K, V] that uses sync.Map for synchronization.
func NewSync[K comparable, V any]() TSMap[K, V] {
	return &syncMap[K, V]{
		mp:   &sync.Map{},
		size: &atomic.Int64{},
	}
}

// WithUntil when first occurrence of exception, stop range.
func WithUntil[K, V any](exception Exception[K, V]) Exception[K, V] {
	stop := false
	return func(key K, value V) bool {
		if stop {
			return true
		}
		if exception(key, value) {
			stop = true
			return true
		}
		return false
	}
}

// defaultException always returns false.
func defaultException[K, V any](key K, value V) bool {
	return false
}
