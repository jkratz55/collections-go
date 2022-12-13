package sync

import (
	"sync"
)

const DefaultShards = 16

type Entry[T any] struct {
	Key   string
	Value T
}

type mapShard[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

type ConcurrentMap[K comparable, V any] struct {
	shards []mapShard[K, V]
}

func NewMap[K comparable, V any](shards int) ConcurrentMap[K, V] {
	if shards < 1 {
		shards = DefaultShards
	}
	mapShards := make([]mapShard[K, V], shards)
	for i := range mapShards {
		mapShards[i] = mapShard[K, V]{
			data:  make(map[K]V, 0),
			mutex: sync.RWMutex{},
		}
	}
	return ConcurrentMap[K, V]{
		shards: mapShards,
	}
}

func (m ConcurrentMap[K, V]) Get(key string) (V, bool) {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Contains(key string) bool {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Set(key string, val V) {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) SetIfPresent(key string, val V) bool {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) SetIfAbsent(key string, val V) bool {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Delete(key string) bool {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Pop(key string) (V, bool) {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Size() uint64 {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) ShardStats() map[int]uint64 {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Keys() []K {
	panic("not implemented")
}

func (m ConcurrentMap[K, V]) Iter() *MapIterator[V] {
	return &MapIterator[V]{}
}

type MapIterator[T any] struct {
}

func (mi MapIterator[T]) Next() bool {
	panic("not implemented")
}

func (mi MapIterator[T]) Value() Entry[T] {
	panic("not implemented")
}
