package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque_Simple(t *testing.T) {
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

func TestDeque_Complex(t *testing.T) {
	deque := NewDeque[int]()

	for i := 0; i < 16; i++ {
		deque.PushBack(i)
	}
	for i := 0; i < 8; i++ {
		deque.PopBack()
	}

	for i := 0; i < 8; i++ {
		deque.PushFront(i)
	}

	expected := []int{7, 6, 5, 4, 3, 2, 1, 0, 0, 1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, deque.AsSlice())

	for i := 0; i < 8; i++ {
		deque.PopFront()
	}

	for i := 0; i < 8; i++ {
		deque.PushBack(i)
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7}, deque.data)
}

func TestDeque_Grow(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < minDequeCapacity+1; i++ {
		deque.PushBack(i)
	}
	assert.Equal(t, 32, len(deque.data))
	assert.Equal(t, 32, deque.Capacity())
	assert.Equal(t, 17, deque.Len())

	val, ok := deque.Front()
	assert.True(t, ok)
	assert.Equal(t, 0, val)

	val, ok = deque.Back()
	assert.True(t, ok)
	assert.Equal(t, 16, val)
}

func TestDeque_Shrink(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 130; i++ {
		deque.PushBack(i)
	}
	assert.Equal(t, 256, len(deque.data))
	assert.Equal(t, 256, deque.Capacity())
	assert.Equal(t, 130, deque.Len())

	for i := 0; i < 96; i++ {
		_, _ = deque.PopBack()
	}

	assert.Equal(t, 34, deque.Len())
	assert.Equal(t, 128, len(deque.data))
	assert.Equal(t, 128, deque.Capacity())

	for i := 0; i < 16; i++ {
		_, _ = deque.PopBack()
	}

	assert.Equal(t, 18, deque.Len())
	assert.Equal(t, 64, len(deque.data))
	assert.Equal(t, 64, deque.Capacity())
}

func TestDeque_PushFront(t *testing.T) {
	deque := NewDeque[int]()
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PushFront(3)

	assert.Equal(t, 3, deque.data[deque.head])
	assert.Equal(t, 1, deque.data[deque.prev(deque.tail)])
}

func TestDeque_PushBack(t *testing.T) {
	deque := NewDeque[int]()
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	assert.Equal(t, 1, deque.data[deque.head])
	assert.Equal(t, 3, deque.data[deque.prev(deque.tail)])
}

func TestDeque_Front(t *testing.T) {
	deque := NewDeque[int]()

	val, ok := deque.Front()
	assert.False(t, ok)

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	val, ok = deque.Front()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
}

func TestDeque_Back(t *testing.T) {
	deque := NewDeque[int]()

	val, ok := deque.Back()
	assert.False(t, ok)

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	val, ok = deque.Back()
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func TestDeque_PopFront(t *testing.T) {
	deque := NewDeque[int]()

	_, ok := deque.PopFront()
	assert.False(t, ok)

	for i := 1; i < 128; i++ {
		deque.PushBack(i)
	}

	for i := 1; i < 128; i++ {
		val, ok := deque.PopFront()
		assert.True(t, ok)
		assert.Equal(t, i, val)
	}
}

func TestDeque_PopBack(t *testing.T) {
	deque := NewDeque[int]()

	_, ok := deque.PopBack()
	assert.False(t, ok)

	for i := 1; i < 128; i++ {
		deque.PushBack(i)
	}

	for i := 128; i <= 1; i-- {
		val, ok := deque.PopBack()
		assert.True(t, ok)
		assert.Equal(t, i, val)
	}
}

func TestDeque_GetAt(t *testing.T) {
	deque := NewDeque[int]()
	for i := 1; i < 128; i++ {
		deque.PushBack(i)
	}

	assert.Equal(t, 1, deque.GetAt(0))
	assert.Equal(t, 2, deque.GetAt(1))
	assert.Equal(t, 3, deque.GetAt(2))

	assert.Panics(t, func() {
		deque.GetAt(-1)
	})

	assert.Panics(t, func() {
		deque.GetAt(128)
	})
}

func TestDeque_Clear(t *testing.T) {
	deque := NewDeque[int]()
	for i := 1; i < 128; i++ {
		deque.PushBack(i)
	}
	deque.Clear()

	assert.Equal(t, 0, deque.count)
	assert.Equal(t, 0, deque.head)
	assert.Equal(t, 0, deque.tail)
}

func TestDeque_AsSlice(t *testing.T) {
	deque := NewDeque[int]()

	assert.Equal(t, []int{}, deque.AsSlice())

	for i := 1; i < 10; i++ {
		deque.PushBack(i)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, deque.AsSlice())
}
