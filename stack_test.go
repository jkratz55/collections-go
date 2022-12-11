package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[string]()
	assert.NotNil(t, stack)
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("World")
	stack.Push("Hello")

	expected := []string{"World", "Hello"}
	assert.Equal(t, expected, stack.data)
	assert.Equal(t, 2, len(stack.data))
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[string]()
	stack.data = append(stack.data, "World", "Hello")
	expected := []string{"Hello", "World"}
	actual := make([]string, 0)

	val, err := stack.Pop()
	assert.NoError(t, err)
	actual = append(actual, val)

	val, err = stack.Pop()
	assert.NoError(t, err)
	actual = append(actual, val)

	val, err = stack.Pop()
	assert.ErrorIs(t, err, ErrEmptyStack)

	assert.Equal(t, expected, actual)
	assert.Equal(t, 0, len(stack.data))
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[string]()
	stack.data = append(stack.data, "World", "Hello")

	val, ok := stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, "Hello", val)
	assert.Equal(t, 2, len(stack.data))

	val, ok = stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, "Hello", val)
	assert.Equal(t, 2, len(stack.data))

	stack.data = make([]string, 0)
	val, ok = stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, 0, len(stack.data))
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[string]()
	assert.Equal(t, 0, stack.Size())
	stack.data = append(stack.data, "test")
	assert.Equal(t, 1, stack.Size())
	stack.data = append(stack.data, "test")
	assert.Equal(t, 2, stack.Size())
}

func TestStack_Empty(t *testing.T) {
	stack := NewStack[string]()
	assert.True(t, stack.Empty())
	stack.data = append(stack.data, "test")
	assert.False(t, stack.Empty())
}

func TestStack_PopEach(t *testing.T) {
	actual := make([]string, 0)
	fn := func(val string) {
		actual = append(actual, val)
	}

	stack := NewStack[string]()
	stack.data = append(stack.data, "World", "Hello")

	stack.PopEach(fn)

	assert.Equal(t, 0, len(stack.data))
	assert.Equal(t, []string{"Hello", "World"}, actual)
}
