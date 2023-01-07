package collections

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type item[T any] struct {
	value    T
	priority int
	index    int
}

type internalPriorityQueue[T any] []*item[T]

func (pq internalPriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq internalPriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq internalPriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *internalPriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *internalPriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// PriorityQueue is a queue data structure that orders elements according to priority.
// When Pop is invoked the item with the highest priority is returned.
//
// The zero-value of PriorityQueue is not usable. Use NewPriorityQueue to create and
// initialize a new PriorityQueue.
type PriorityQueue[T any] struct {
	internal internalPriorityQueue[T]
}

// NewPriorityQueue creates and initializes a new PriorityQueue.
func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		internal: make(internalPriorityQueue[T], 0),
	}
	heap.Init(&pq.internal)
	return pq
}

// Push adds an element to the PriorityQueue with the specified priority.
func (pq *PriorityQueue[T]) Push(val T, priority int) {
	item := &item[T]{
		value:    val,
		priority: priority,
	}
	heap.Push(&pq.internal, item)
}

// Poll retrieves the highest priority item from the queue and removes it from
// the PriorityQueue. If the queue is empty the zero value is returned with a
// boolean value of false.
func (pq *PriorityQueue[T]) Poll() (T, bool) {
	if len(pq.internal) == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(&pq.internal).(*item[T])
	return item.value, true
}

// Peek returns the next element to be polled from the PriorityQueue. If the
// queue is empty the zero value is returned with a boolean value of false.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if len(pq.internal) == 0 {
		var zero T
		return zero, false
	}
	return pq.internal[0].value, true
}

// Len returns the length/size of the PriorityQueue.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.internal)
}

// IsEmpty returns true if the PriorityQueue is empty, otherwise false.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.internal) == 0
}
