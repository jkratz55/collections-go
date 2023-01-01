package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	// Since Stack is just a wrapper around Deque this is more/less an integration
	// test to ensure Stack is behaving as expected. Since Stack is just delegating
	// out to Deque it is safe to make the assumption Stack itself is working if this
	// passes. No need to inspect the internals of Deque, that is what Deque unit test
	// are for.

	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	assert.False(t, stack.Empty())
	assert.Equal(t, 5, stack.Size())

	val, ok := stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, 5, val)

	for i := 5; i > 0; i-- {
		val, err := stack.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, val)
		assert.Equal(t, i-1, stack.Size())
	}

	assert.True(t, stack.Empty())

	val, err := stack.Pop()
	assert.Error(t, err)

	val, ok = stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, 0, val)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	expected := []int{5, 4, 3, 2, 1}
	actual := make([]int, 0)
	stack.PopEach(func(val int) {
		actual = append(actual, val)
	})
	assert.Equal(t, expected, actual)
	assert.True(t, stack.Empty())
}
