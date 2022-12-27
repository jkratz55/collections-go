package sync

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConcurrentMap(t *testing.T) {
	assert.NotPanics(t, func() {
		m := NewConcurrentMap[string, int](16, StringHasher())
		assert.Equal(t, 16, len(m.shards))
		assert.Equal(t, uint(16), m.shardCount)
	})

	assert.Panics(t, func() {
		_ = NewConcurrentMap[string, int](0, nil)
	})
}

func TestConcurrentMap_Get(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		expectedValue string
		expectedFound bool
		initFunc      func(m ConcurrentMap[string, string])
	}{
		{
			name:          "Fetch Existing Key",
			key:           "hello",
			expectedValue: "world",
			expectedFound: true,
			initFunc: func(m ConcurrentMap[string, string]) {
				m.Set("hello", "world")
			},
		},
		{
			name:          "Key Doesn't Exist",
			key:           "test",
			expectedValue: "",
			expectedFound: false,
			initFunc: func(m ConcurrentMap[string, string]) {

			},
		},
	}

	for _, test := range tests {
		m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
		test.initFunc(m)
		val, ok := m.Get(test.key)
		assert.Equal(t, test.expectedValue, val)
		assert.Equal(t, ok, test.expectedFound)
	}
}

func TestConcurrentMap_MGet(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	m.Set("hello", "world")
	m.Set("test", "test")
	m.Set("os", "macOS")

	vals := m.MGet("hello", "test", "os", "blah")
	assert.Equal(t, []string{"world", "test", "macOS"}, vals)
}

func TestConcurrentMap_Pop(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		expectedValue string
		expectedFound bool
		init          func(m ConcurrentMap[string, string])
	}{
		{
			name:          "Value doesn't exist",
			key:           "hello",
			expectedValue: "",
			expectedFound: false,
			init: func(m ConcurrentMap[string, string]) {
				// do nothing
			},
		},
		{
			name:          "Value Successfully Popped",
			key:           "hello",
			expectedValue: "world",
			expectedFound: true,
			init: func(m ConcurrentMap[string, string]) {
				m.Set("hello", "world")
			},
		},
	}

	for _, test := range tests {
		m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
		test.init(m)
		val, ok := m.Pop(test.key)
		assert.Equal(t, test.expectedValue, val)
		assert.Equal(t, test.expectedFound, ok)
		if ok {
			_, found := m.Get(test.key)
			assert.False(t, found)
		}
	}
}

func TestConcurrentMap_Set(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	m.Set("hello", "world")
	m.Set("test", "test")

	val, ok := m.Get("hello")
	assert.Equal(t, "world", val)
	assert.True(t, ok)

	val, ok = m.Get("test")
	assert.Equal(t, "test", val)
	assert.True(t, ok)
}

func TestConcurrentMap_SetIfAbsent(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	set := m.SetIfAbsent("hello", "world")
	assert.True(t, set)

	set = m.SetIfAbsent("hello", "billy")
	assert.False(t, set)

	set = m.SetIfAbsent("mellow", "world")
	assert.True(t, set)
}

func TestConcurrentMap_SetIfPresent(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())

	set := m.SetIfPresent("hello", "world")
	assert.False(t, set)

	m.Set("hello", "world")

	set = m.SetIfPresent("hello", "billy")
	assert.True(t, set)

	set = m.SetIfPresent("mellow", "world")
	assert.False(t, set)
}

func TestConcurrentMap_Size(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("%d", i), "hello")
	}
	assert.Equal(t, uint64(1000), m.Size())
}

func TestConcurrentMap_SizeByShard(t *testing.T) {
	// todo: this should be unit tested
}

func TestConcurrentMap_MSet(t *testing.T) {
	data := map[string]string{
		"hello": "world",
		"hi":    "bye",
		"ninja": "turtles",
	}
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	m.MSet(data)

	assert.Equal(t, uint64(3), m.Size())

	val, found := m.Get("hello")
	assert.True(t, found)
	assert.Equal(t, "world", val)

	val, found = m.Get("hi")
	assert.True(t, found)
	assert.Equal(t, "bye", val)

	val, found = m.Get("ninja")
	assert.True(t, found)
	assert.Equal(t, "turtles", val)
}

func TestConcurrentMap_Contains(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	exist := m.Contains("hello")
	assert.False(t, exist)

	m.Set("hello", "world")
	exist = m.Contains("hello")
	assert.True(t, exist)
}

func TestConcurrentMap_Delete(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	deleted := m.Delete("hello")
	assert.False(t, deleted)

	m.Set("hello", "")
	deleted = m.Delete("hello")
	assert.True(t, deleted)
}

func TestConcurrentMap_Keys(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	m.Set("hello", "world")
	m.Set("test", "test")

	keys := m.Keys()
	assert.Equal(t, 2, len(keys))
	assert.ElementsMatch(t, []string{"hello", "test"}, keys)
}

func TestConcurrentMap_Iterator(t *testing.T) {
	m := NewConcurrentMap[string, string](DefaultShards, StringHasher())
	m.Set("hello", "world")
	m.Set("test", "test")

	iter := m.Iterator()
	entries := []Entry[string, string]{}
	for iter.Next() {
		entries = append(entries, iter.Value())
	}

	assert.Equal(t, 2, len(entries))
	assert.ElementsMatch(t, []Entry[string, string]{
		{
			Key:   "hello",
			Value: "world",
		},
		{
			Key:   "test",
			Value: "test",
		},
	}, entries)
}

func TestConcurrentMap_GenericHasher(t *testing.T) {
	type EmployeeID uint

	m := NewConcurrentMap[EmployeeID, string](DefaultShards, NewHasher[EmployeeID]())
	m.Set(EmployeeID(11110000), "Billy Bob")
	m.Set(EmployeeID(22220000), "Jane Doe")
	m.Set(EmployeeID(83243243), "Agent 47")

	val, ok := m.Get(EmployeeID(83243243))
	assert.True(t, ok)
	assert.Equal(t, "Agent 47", val)
}
