package collections

import (
	"github.com/jkratz55/collections-go/internal"
)

// OrderedMap is a map implementation that maintains the order keys were inserted
// into the map. Set, Get, and Delete operations remain O(1) and the keys can be
// iterated in the order they were added using ForEach, or in reverse order using
// ForEachReverse.
//
// The zero-value of OrderedMap is not usable. NewOrderedMap should be used to
// create and initialize a new instance of OrderedMap.
type OrderedMap[K comparable, V any] struct {
	keys internal.KeyList[K, V]
	data map[K]*internal.Element[K, V]
}

// NewOrderedMap creates and initializes a new OrderedMap
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		data: make(map[K]*internal.Element[K, V]),
		keys: internal.KeyList[K, V]{},
	}
}

// Set inserts a new key/value into the map or replaces the value for an existing
// key. If the key didn't exist in the map and was inserted true is returned.
// Otherwise, if they key was already existing returns false.
func (m *OrderedMap[K, V]) Set(key K, val V) bool {
	if _, exists := m.data[key]; !exists {
		element := m.keys.PushBack(key, val)
		m.data[key] = element
		return true
	}
	m.data[key].Value = val
	return false
}

// Contains returns true if the given key exists in the OrderedMap, otherwise
// returns false.
func (m *OrderedMap[K, V]) Contains(key K) bool {
	_, exists := m.data[key]
	return exists
}

// Get retrieves the value for a key. It follows the same idioms of the built-in
// map. If the key doesn't exist the zero value and false value are returned.
func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := m.data[key]
	if !exists {
		var zero V
		return zero, false
	}
	return val.Value, true
}

// GetOrDefault retrieves the value for a key and if it doesn't exist returns the
// provided default value.
func (m *OrderedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	val, exists := m.data[key]
	if !exists {
		return defaultValue
	}
	return val.Value
}

// Delete removes a key/value entry from the OrderedMap. If the key didn't exist
// returns false so, otherwise returns true if the entry was deleted.
func (m *OrderedMap[K, V]) Delete(key K) bool {
	elem, exists := m.data[key]
	if !exists {
		return false
	}
	m.keys.Remove(elem)
	delete(m.data, key)
	return true
}

// Size returns the number of entries in the OrderedMap
func (m *OrderedMap[K, V]) Size() int {
	return len(m.data)
}

// Keys returns the keys in the order they were inserted.
func (m *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for e := m.keys.Front(); e != nil; e = e.Next() {
		keys = append(keys, e.Key)
	}
	return keys
}

// ForEach iterates through the map entries passing the key/value pair to the
// provided function/closure in the order they were inserted.
func (m *OrderedMap[K, V]) ForEach(fn func(key K, val V)) {
	for e := m.keys.Front(); e != nil; e = e.Next() {
		fn(e.Key, e.Value)
	}
}

// ForEachReverse iterates through the map entries passing the key/value pairs to
// the provided function/closure in the reverse order they were inserted (most
// recently inserted to least recently)
func (m *OrderedMap[K, V]) ForEachReverse(fn func(key K, val V)) {
	for e := m.keys.Back(); e != nil; e = e.Prev() {
		fn(e.Key, e.Value)
	}
}
