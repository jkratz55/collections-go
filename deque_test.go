package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {
	deque := NewDeque[string]()
	deque.PushBack("hello")
	deque.PushBack("world")
	deque.PushBack("end")

	val, ok := deque.PopFront()
	assert.True(t, ok)
	assert.Equal(t, "hello", val)

	val, ok = deque.PopFront()
	assert.True(t, ok)
	assert.Equal(t, "world", val)

	val, ok = deque.PopFront()
	assert.True(t, ok)
	assert.Equal(t, "end", val)

	val, ok = deque.PopFront()
	assert.False(t, ok)
}
