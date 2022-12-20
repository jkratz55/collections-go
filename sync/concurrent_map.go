package sync

import (
	"sync"
)

// DefaultShards is the default shards a ConcurrentMap will use.
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

// mapShard is a shard of data in a ConcurrentMap. It contains the underlying
// data as map[K]V and a RWMutex to protect that data.
type mapShard[K comparable, V any] struct {
	data map[K]V
	sync.RWMutex
}

// ConcurrentMap is a thread safe sharded map implementation. ConcurrentMap shards
// data to reduce lock contention.
//
// The zero-value of ConcurrentMap is not usable. Instead, NewConcurrentMap function
// should be used to crete and initialize a new ConcurrentMap.
type ConcurrentMap[K comparable, V any] struct {
	shards     []*mapShard[K, V]
	hasher     Hasher[K]
	shardCount uint
}

// NewConcurrentMap creates and initializes a new empty ConcurrentMap. NewConcurrentMap
// accepts two required parameters, number of shards, and the Hasher to hash keys.
// If value for shards < 1 than DefaultShards will be used. If a nil Hasher is provided
// this function will panic.
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

// Get retrieves a single element from the ConcurrentMap. Get follows the same
// semantics of the built-in map returning the value and a boolean indicating
// if the key exists.
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

// Contains returns a boolean indicating if the key exists.
func (m ConcurrentMap[K, V]) Contains(key K) bool {
	shard := m.getShard(key)
	shard.RLock()
	_, ok := shard.data[key]
	shard.RUnlock()
	return ok
}

// Set inserts or updates ConcurrentMap by setting the key value pair. If the key
// already exists its value is overridden.
func (m ConcurrentMap[K, V]) Set(key K, val V) {
	shard := m.getShard(key)
	shard.Lock()
	shard.data[key] = val
	shard.Unlock()
}

// SetIfPresent sets the value for a given key only if they key already exists
// in the ConcurrentMap. This is essentially an update only operation.
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

// SetIfAbsent set the value for a given key only if the key doesn't already
// exist in the ConcurrentMap. This is essentially a insert only operation.
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

// MSet performs a Set operation on multiple key-value paris supplied
// as a map.
func (m ConcurrentMap[K, V]) MSet(data map[K]V) {
	for key, val := range data {
		shard := m.getShard(key)
		shard.Lock()
		shard.data[key] = val
		shard.Unlock()
	}
}

// Delete deletes a single key/value from the ConcurrentMap returning
// a boolean indicating if the key was present or not.
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

// Pop fetching the value for a given key and if the key was found deletes that
// key from the ConcurrentMap
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

// Size returns the approx size (number of elements) in the ConcurrentMap.
// The size is approximated due to the nature of how the data is sharded.
// To prevent lock contention each shard is processed and shards already
// processed may have undergone changes by the time this function returns.
func (m ConcurrentMap[K, V]) Size() uint64 {
	size := uint64(0)
	for _, shard := range m.shards {
		shard.RLock()
		size = size + uint64(len(shard.data))
		shard.RUnlock()
	}
	return size
}

// SizeByShard returns the approx size (number of elements) per shard in the
// ConcurrentMap. The returned values represent the size of the shard at the
// time a read lock was acquired on it. The returned values may not be exact
// as the shards may have been modified after determining their size.
func (m ConcurrentMap[K, V]) SizeByShard() map[int]int {
	stats := make(map[int]int)
	for i := range m.shards {
		shard := m.shards[i]
		shard.RLock()
		stats[i] = len(shard.data)
		shard.RUnlock()
	}
	return stats
}

// Keys returns all the keys in the ConcurrentMap. Keys is an atomic operation
// across all shards. A read lock is acquired on all shards before retrieving
// the keys. Keys can be an expensive operation and is not recommended to be
// called often.
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

// Iterator returns an iterator to iterate through all the entries in the ConcurrentMap.
//
// Internally Iterator takes a snapshot of each shard one by one, acquiring a read
// lock as it snapshots each shard. This is done for performance reasons as to not
// hold locks as the caller iterates through the ConcurrentMap. However, this means
// that the data being iterated through could potentially be stale.
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
