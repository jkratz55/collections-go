package sync

import (
	"sync"
)

// Map is a very thin layer over sync.Map to make it type-safe.
type Map[K comparable, V any] struct {
	internal sync.Map
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	val, ok := m.internal.Load(key)
	return val.(V), ok
}

func (m *Map[K, V]) Store(key K, val V) {
	m.internal.Store(key, val)
}

func (m *Map[K, V]) LoadOrStore(key any, value any) (V, bool) {
	val, ok := m.internal.LoadOrStore(key, value)
	return val.(V), ok
}

func (m *Map[K, V]) LoadAndDelete(key any) (V, bool) {
	val, ok := m.internal.LoadAndDelete(key)
	return val.(V), ok
}

func (m *Map[K, V]) Delete(key any) {
	m.internal.Delete(key)
}

func (m *Map[K, V]) Range(fn func(key K, val V) bool) {
	m.internal.Range(func(k, v any) bool {
		return fn(k.(K), v.(V))
	})
}
