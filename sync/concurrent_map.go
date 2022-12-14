package sync

import (
	"sync"
)

const DefaultShards = 16

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Hasher is a function type that accepts a key that satisfies comparable and
// returns a hashcode.
//
// The Hasher is how ConcurrentMap decides how to shard the data.
type Hasher[T comparable] func(key T) uint32

// StringHasher returns a default Hasher for handling string keys.
func StringHasher() Hasher[string] {
	return func(key string) uint32 {
		hash := uint32(2166136261)
		const prime32 = uint32(16777619)
		keyLength := len(key)
		for i := 0; i < keyLength; i++ {
			hash *= prime32
			hash ^= uint32(key[i])
		}
		return hash
	}
}

type mapShard[K comparable, V any] struct {
	data map[K]V
	sync.RWMutex
}

type ConcurrentMap[K comparable, V any] struct {
	shards     []*mapShard[K, V]
	hasher     Hasher[K]
	shardCount uint
}

func NewConcurrentMap[K comparable, V any](shards int, hasher Hasher[K]) ConcurrentMap[K, V] {
	if shards < 1 {
		shards = DefaultShards
	}
	if hasher == nil {
		panic("illegal use of API, cannot use ConcurrentMap with nil Hasher")
	}
	mapShards := make([]*mapShard[K, V], shards)
	for i := range mapShards {
		mapShards[i] = &mapShard[K, V]{
			data:    make(map[K]V, 0),
			RWMutex: sync.RWMutex{},
		}
	}
	return ConcurrentMap[K, V]{
		shards:     mapShards,
		hasher:     hasher,
		shardCount: uint(shards),
	}
}

func (m ConcurrentMap[K, V]) Get(key K) (V, bool) {
	shard := m.getShard(key)
	shard.RLock()
	val, ok := shard.data[key]
	shard.RUnlock()
	return val, ok
}

// MGet fetches multiple keys and returns the values as a slice.
//
// MGet takes a pessimistic approach and acquires a read lock on all shards in
// the ConcurrentMap before fetching keys. After all the keys have been fetched
// the locks are released. This might lead to lock contention in cases where
// there are heavy writes.
func (m ConcurrentMap[K, V]) MGet(keys ...K) []V {
	// Since we don't know which shards the keys might reside in without iterating
	// over all the keys we pessimistically acquire a read lock on all shards.
	// Notes: Many IDEs and linters may complain about defer in a for loop but in
	// this case we want to defer unlocking until we are completely done fetching
	// all keys.
	for i := range m.shards {
		m.shards[i].RLock()
		defer m.shards[i].RUnlock()
	}
	values := make([]V, 0, len(keys))
	for _, key := range keys {
		shard := m.getShard(key)
		if val, ok := shard.data[key]; ok {
			values = append(values, val)
		}
	}
	return values
}

func (m ConcurrentMap[K, V]) Contains(key K) bool {
	shard := m.getShard(key)
	shard.RLock()
	_, ok := shard.data[key]
	shard.RUnlock()
	return ok
}

func (m ConcurrentMap[K, V]) Set(key K, val V) {
	shard := m.getShard(key)
	shard.Lock()
	shard.data[key] = val
	shard.Unlock()
}

func (m ConcurrentMap[K, V]) SetIfPresent(key K, val V) bool {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	if _, ok := shard.data[key]; ok {
		shard.data[key] = val
		return true
	}
	return false
}

func (m ConcurrentMap[K, V]) SetIfAbsent(key K, val V) bool {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	if _, ok := shard.data[key]; !ok {
		shard.data[key] = val
		return true
	}
	return false
}

func (m ConcurrentMap[K, V]) MSet(data map[K]V) {
	for key, val := range data {
		shard := m.getShard(key)
		shard.Lock()
		shard.data[key] = val
		shard.Unlock()
	}
}

func (m ConcurrentMap[K, V]) Delete(key K) bool {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	_, ok := shard.data[key]
	if !ok {
		return false
	}
	delete(shard.data, key)
	return true
}

func (m ConcurrentMap[K, V]) Pop(key K) (V, bool) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	val, ok := shard.data[key]
	if ok {
		delete(shard.data, key)
	}
	return val, ok
}

func (m ConcurrentMap[K, V]) Size() uint64 {
	size := uint64(0)
	for _, shard := range m.shards {
		shard.RLock()
		size = size + uint64(len(shard.data))
		shard.RUnlock()
	}
	return size
}

func (m ConcurrentMap[K, V]) ShardStats() map[int]int {
	stats := make(map[int]int)
	for i := range m.shards {
		shard := m.shards[i]
		shard.RLock()
		stats[i] = len(shard.data)
		shard.RUnlock()
	}
	return stats
}

func (m ConcurrentMap[K, V]) Keys() []K {
	keys := make([]K, 0)
	for i := range m.shards {
		m.shards[i].RLock()
		defer m.shards[i].RUnlock()
	}

	for i := range m.shards {
		shard := m.shards[i]
		for key := range shard.data {
			keys = append(keys, key)
		}
	}
	return keys
}

func (m ConcurrentMap[K, V]) Iterator() *ConcurrentMapIterator[K, V] {
	return &ConcurrentMapIterator[K, V]{
		data:    m.snapshot(),
		current: -1,
	}
}

func (m ConcurrentMap[K, V]) getShard(key K) *mapShard[K, V] {
	return m.shards[uint(m.hasher(key))%m.shardCount]
}

func (m ConcurrentMap[K, V]) snapshot() []Entry[K, V] {
	data := make([]Entry[K, V], 0)
	for i := range m.shards {
		shard := m.shards[i]
		shard.RLock()
		for key, val := range shard.data {
			data = append(data, Entry[K, V]{
				Key:   key,
				Value: val,
			})
		}
		shard.RUnlock()
	}
	return data
}

type ConcurrentMapIterator[K comparable, V any] struct {
	current int
	data    []Entry[K, V]
}

func (cmi *ConcurrentMapIterator[K, V]) Next() bool {
	cmi.current++
	if cmi.current >= len(cmi.data) {
		return false
	}
	return true
}

func (cmi *ConcurrentMapIterator[K, V]) Value() Entry[K, V] {
	return cmi.data[cmi.current]
}
