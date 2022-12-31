package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	// Since Queue is just a wrapper around Deque this is more/less an integration
	// test to ensure Queue is behaving as expected. Since Queue is just delegating
	// out to Deque it is safe to make the assumption Queue itself is working if this
	// passes. No need to inspect the internals of Deque, that is what Deque unit test
	// are for.

	q := NewQueue[int]()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	q.Offer(4)
	q.Offer(5)

	assert.False(t, q.Empty())
	assert.Equal(t, 5, q.Size())

	val, ok := q.Peek()
	assert.Equal(t, 1, val)
	assert.True(t, ok)

	for i := 1; i <= 5; i++ {
		val, ok = q.Poll()
		assert.True(t, ok)
		assert.Equal(t, i, val)
		assert.Equal(t, 5-i, q.Size())
	}

	assert.True(t, q.Empty())

	val, ok = q.Peek()
	assert.False(t, ok)
	assert.Equal(t, 0, val)

	val, ok = q.Poll()
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}
