package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type blah struct {
}

func (b blah) String() string {
	// TODO implement me
	panic("implement me")
}

func TestNewConcurrentMap(t *testing.T) {
	assert.NotPanics(t, func() {
		m := NewConcurrentMap[string, int](16, StringHasher())
		assert.Equal(t, 16, len(m.shards))
		assert.Equal(t, 16, m.shardCount)
	})

	assert.Panics(t, func() {
		_ = NewConcurrentMap[string, int](0, nil)
	})

	m := NewConcurrentMap[blah, string](16, StringerHasher())
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
			name:          "",
			key:           "",
			expectedValue: "",
			expectedFound: false,
			initFunc:      nil,
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
