package sync

import (
	"sync"
)

// Map is a very thin layer over sync.Map to make it type-safe.
//
// This implementation is intended as a simple stop gap until the standard library
// adds a generic version of sync.Map.
type Map[K comparable, V any] struct {
	internal sync.Map
}

// Load returns the value stored in the map for a key, or nil if no value is
// present. The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (V, bool) {
	val, ok := m.internal.Load(key)
	return val.(V), ok
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, val V) {
	m.internal.Store(key, val)
}

// LoadOrStore returns the existing value for the key if present. Otherwise,
// it stores and returns the given value. The loaded result is true if the
// value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key any, value any) (V, bool) {
	val, ok := m.internal.LoadOrStore(key, value)
	return val.(V), ok
}

// LoadAndDelete deletes the value for a key, returning the previous value
// if any. The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key any) (V, bool) {
	val, ok := m.internal.LoadAndDelete(key)
	return val.(V), ok
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key any) {
	m.internal.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of
// the Map's contents: no key will be visited more than once, but if the
// value for any key is stored or deleted concurrently (including by f),
// Range may reflect any mapping for that key from any point during the
// Range call. Range does not block other methods on the receiver; even
// f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f
// returns false after a constant number of calls.
func (m *Map[K, V]) Range(fn func(key K, val V) bool) {
	m.internal.Range(func(k, v any) bool {
		return fn(k.(K), v.(V))
	})
}
