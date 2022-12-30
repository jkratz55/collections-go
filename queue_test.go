package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Offer(t *testing.T) {
	q := Queue[int]{}
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)

	assert.Equal(t, 1, q.data[0])
	assert.Equal(t, 2, q.data[1])
	assert.Equal(t, 3, q.data[2])
}

func TestQueue_Peek(t *testing.T) {
	q := Queue[int]{}
	val, ok := q.Peek()
	assert.Equal(t, 0, val)
	assert.False(t, ok)

	q.data = append(q.data, 1)
	q.data = append(q.data, 2)

	val, ok = q.Peek()
	assert.Equal(t, 1, val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, 1, val)
	assert.True(t, ok)
}

func TestQueue_Poll(t *testing.T) {
	q := Queue[int]{}
	val, ok := q.Poll()
	assert.Equal(t, 0, val)
	assert.False(t, ok)

	q.data = append(q.data, 1)
	q.data = append(q.data, 2)
	assert.Equal(t, 2, len(q.data))

	val, ok = q.Poll()
	assert.Equal(t, 1, val)
	assert.True(t, ok)

	assert.Equal(t, 1, len(q.data))
	assert.Equal(t, 2, q.data[0])
}

func TestQueue_Size(t *testing.T) {
	q := Queue[int]{}
	assert.Equal(t, 0, q.Size())
	q.data = append(q.data, 1)
	q.data = append(q.data, 2)
	assert.Equal(t, 2, len(q.data))
	assert.Equal(t, 2, q.Size())
}

func TestQueue_Empty(t *testing.T) {
	q := Queue[int]{}
	assert.True(t, q.Empty())
	q.data = append(q.data, 1)
	q.data = append(q.data, 2)
	assert.False(t, q.Empty())
}
