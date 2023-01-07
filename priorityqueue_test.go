package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_Push(t *testing.T) {
	t.Run("All Same Priority", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Push(1, 0)
		pq.Push(2, 0)
		pq.Push(3, 0)
		pq.Push(4, 0)
		pq.Push(5, 0)

		elements := getElementsFromInternals(pq)
		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, elements)
	})

	t.Run("Insert Higher Priority in Middle", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Push(1, 0)
		pq.Push(2, 0)
		pq.Push(3, 10)
		pq.Push(4, 0)
		pq.Push(5, 0)

		elements := getElementsFromInternals(pq)
		expected := []int{3, 2, 1, 4, 5}
		assert.Equal(t, expected, elements)
	})
}

func TestPriorityQueue_Poll(t *testing.T) {
	t.Run("Poll Empty Queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		val, ok := pq.Poll()
		assert.False(t, ok)
		assert.Equal(t, 0, val)
	})

	t.Run("Poll Populated Queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Push(1, 0)
		pq.Push(2, 2)
		pq.Push(3, 10)
		pq.Push(4, 3)
		pq.Push(5, 5)

		val, ok := pq.Poll()
		assert.True(t, ok)
		assert.Equal(t, 3, val)

		val, ok = pq.Poll()
		assert.True(t, ok)
		assert.Equal(t, 5, val)

		val, ok = pq.Poll()
		assert.True(t, ok)
		assert.Equal(t, 4, val)

		val, ok = pq.Poll()
		assert.True(t, ok)
		assert.Equal(t, 2, val)

		val, ok = pq.Poll()
		assert.True(t, ok)
		assert.Equal(t, 1, val)

		val, ok = pq.Poll()
		assert.False(t, ok)
		assert.Equal(t, 0, val)
	})
}

func TestPriorityQueue_Peek(t *testing.T) {
	t.Run("Empty Queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		val, ok := pq.Peek()
		assert.False(t, ok)
		assert.Equal(t, 0, val)
	})

	t.Run("Populated Queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Push(1, 0)
		pq.Push(2, 2)
		pq.Push(3, 10)
		pq.Push(4, 3)
		pq.Push(5, 5)

		val, ok := pq.Peek()
		assert.True(t, ok)
		assert.Equal(t, 3, val)

		val, ok = pq.Peek()
		assert.True(t, ok)
		assert.Equal(t, 3, val)
	})
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := NewPriorityQueue[int]()
	assert.Equal(t, 0, pq.Len())

	pq.Push(1, 0)
	pq.Push(2, 2)
	pq.Push(3, 10)
	pq.Push(4, 3)
	pq.Push(5, 5)

	assert.Equal(t, 5, pq.Len())

	pq.Poll()
	assert.Equal(t, 4, pq.Len())
}

func TestPriorityQueue_IsEmpty(t *testing.T) {
	pq := NewPriorityQueue[int]()
	assert.True(t, pq.IsEmpty())

	pq.Push(1, 0)
	pq.Push(2, 2)
	pq.Push(3, 10)
	pq.Push(4, 3)
	pq.Push(5, 5)

	assert.False(t, pq.IsEmpty())
}

func getElementsFromInternals[T any](pq *PriorityQueue[T]) []T {
	elements := make([]T, 0)
	for _, element := range pq.internal {
		elements = append(elements, element.value)
	}
	return elements
}
