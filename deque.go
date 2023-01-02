package collections

import (
	"fmt"
)

const minDequeCapacity = 16

// Deque is a linear collection of elements that supports insertion and deletion
// at both ends. Deque is short for double ended queue.
//
// Deque efficiently handles adding and removing elements at either end with O(1)
// performance. This is accomplished using a ring buffer to store the data which
// reduces pressure on garbage collector and more efficiently uses memory when
// compared to implementations that use linked lists or slices.
//
// Deque can be used as both a queue or a stack.
//
//	Stack
//		PushBack - Push element to top of stack
//		PopBack - Pop top element off the stack
//		Back - Peek at the top of the stack
//
//	Queue
//		PushBack- Add element to end of the queue
//		PopFront - Poll the first element at the head of the queue
//		Front - Peek at the next element in the queue
//
// This package also provides Queue and Stack types which wraps Deque in a simpler
// API.
//
// The zero-value of Deque is not usable. New instances of Deque should be created
// and initialized using NewDeque function.
type Deque[T any] struct {
	data  []T
	head  int
	tail  int
	count int
}

// NewDeque creates and initializes a new empty Deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		data:  make([]T, minDequeCapacity),
		head:  0,
		tail:  0,
		count: 0,
	}
}

// PushFront pushes a new element at the front of the Deque
func (q *Deque[T]) PushFront(elem T) {
	q.lazyGrow()
	q.head = q.prev(q.head)
	q.data[q.head] = elem
	q.count++
}

// PushBack pushes a new element to the back of the Deque
func (q *Deque[T]) PushBack(elem T) {
	q.lazyGrow()
	q.data[q.tail] = elem
	q.tail = q.next(q.tail)
	q.count++
}

// PopFront pops the front/first element of the Deque removing it and returning
// the value. If the Deque is empty the zero-value will be returned with a bool
// value of false.
func (q *Deque[T]) PopFront() (T, bool) {
	var zero T
	if q.count <= 0 {
		return zero, false
	}

	val := q.data[q.head]
	q.data[q.head] = zero
	q.head = q.next(q.head)
	q.count--

	q.lazyShrink()
	return val, true
}

// PopBack pops the back/last element of the Deque removing it and returning
// the value. If the Deque is empty the zero-value will be returned with a bool
// value of false.
func (q *Deque[T]) PopBack() (T, bool) {
	var zero T
	if q.count <= 0 {
		return zero, false
	}

	q.tail = q.prev(q.tail)

	val := q.data[q.tail]
	q.data[q.tail] = zero
	q.count--

	q.lazyShrink()
	return val, true
}

// Front returns the first element of the Deque. If the Deque is empty the
// zero-value will be returned with a bool value of false.
func (q *Deque[T]) Front() (T, bool) {
	if q.count <= 0 {
		var zero T
		return zero, false
	}
	return q.data[q.head], true
}

// Back returns the last element of the Deque. If the Deque is empty the
// zero-value will be returned with a bool value of false.
func (q *Deque[T]) Back() (T, bool) {
	if q.count <= 0 {
		var zero T
		return zero, false
	}
	return q.data[q.prev(q.tail)], true
}

// GetAt returns the element at the provided index.
//
// If the index is not valid this will panic.
func (q *Deque[T]) GetAt(idx int) T {
	if idx < 0 || idx >= q.count {
		panic(fmt.Sprintf("index %d is out of bounds", idx))
	}
	// bitwise modulus
	return q.data[(q.head+idx)&(len(q.data)-1)]
}

// Clear removes all the elements from the Deque but retains the current capacity.
func (q *Deque[T]) Clear() {
	var zero T
	modBits := len(q.data) - 1
	h := q.head
	for i := 0; i < q.Len(); i++ {
		q.data[(h+i)&modBits] = zero
	}
	q.head = 0
	q.tail = 0
	q.count = 0
}

// Capacity returns the size of the underlying ring-buffer holding the elements.
func (q *Deque[T]) Capacity() int {
	return cap(q.data)
}

// Len returns the number of elements in the Deque.
func (q *Deque[T]) Len() int {
	return q.count
}

func (q *Deque[T]) IsEmpty() bool {
	return q.count == 0
}

// AsSlice returns an array containing all the elements of the Deque in order
// of front/head to back/tail.
func (q *Deque[T]) AsSlice() []T {
	res := make([]T, 0)
	for i := 0; i < q.count; i++ {
		res = append(res, q.data[(q.head+i)&(len(q.data)-1)])
	}
	return res
}

// prev returns the previous buffer position
func (q *Deque[T]) prev(idx int) int {
	return (idx - 1) & (len(q.data) - 1)
}

// next returns the next buffer position
func (q *Deque[T]) next(idx int) int {
	return (idx + 1) & (len(q.data) - 1)
}

// lazyGrow inspects the size of the buffer and resizes it to expand if needed.
// Otherwise, this simply returns early.
func (q *Deque[T]) lazyGrow() {
	// If the buffer doesn't need to grow skip and return
	if q.count != len(q.data) {
		return
	}
	if len(q.data) == 0 {
		q.data = make([]T, minDequeCapacity)
	}
	q.resize()
}

// lazyShrink inspects the size of the buffer and resizes it shrinking it if the
// buffer is 1/4 full or less.
func (q *Deque[T]) lazyShrink() {
	if len(q.data) > minDequeCapacity && (q.count<<2) == len(q.data) {
		q.resize()
	}
}

// resize internally resizes the buffer to fit exactly twice it's current size.
// This is used to grow and shrink the buffer.
func (q *Deque[T]) resize() {
	newData := make([]T, q.count<<1)
	// If the buffer is not wrapping around simply copy head position to tail
	// position to new larger buffer.
	if q.tail > q.head {
		copy(newData, q.data[q.head:q.tail])
	} else {
		// Otherwise the buffer is wrapping around so the deque from head position
		// until the end needs to be copied first, and then from the beginning of
		// the buffer until that tail
		temp := copy(newData, q.data[q.head:])
		copy(newData[temp:], q.data[:q.tail])
	}
	q.head = 0
	q.tail = q.count
	q.data = newData
}
