package sync

import (
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

}

func TestConcurrentMap_Set(t *testing.T) {

}

func TestConcurrentMap_SetIfAbsent(t *testing.T) {

}

func TestConcurrentMap_SetIfPresent(t *testing.T) {

}

func TestConcurrentMap_Size(t *testing.T) {

}

func TestConcurrentMap_SizeByShard(t *testing.T) {

}

func TestConcurrentMap_MSet(t *testing.T) {

}

func TestConcurrentMap_Contains(t *testing.T) {

}

func TestConcurrentMap_Delete(t *testing.T) {

}

func TestConcurrentMap_Keys(t *testing.T) {

}

func TestConcurrentMap_Iterator(t *testing.T) {

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
