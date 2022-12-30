package collections

import (
	"errors"
)

const minDequeCapacity = 16

var (
	// ErrInvalidIndex is a sentinel error value indicating the index provided
	// for an operation wasn't valid or in bounds.
	ErrInvalidIndex = errors.New("invalid index")
)

type Deque[T any] struct {
	data  []T
	head  int
	tail  int
	count int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		data:  make([]T, minDequeCapacity),
		head:  0,
		tail:  0,
		count: 0,
	}
}

func (q *Deque[T]) PushFront(elem T) {
	q.lazyGrow()
	q.head = q.prev(q.head)
	q.data[q.head] = elem
	q.count++
}

func (q *Deque[T]) PushBack(elem T) {
	q.lazyGrow()
	q.data[q.tail] = elem
	q.tail = q.next(q.tail)
	q.count++
}

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

func (q *Deque[T]) Front() (T, bool) {
	if q.count <= 0 {
		var zero T
		return zero, false
	}
	return q.data[q.head], true
}

func (q *Deque[T]) Back() (T, bool) {
	if q.count <= 0 {
		var zero T
		return zero, false
	}
	return q.data[q.prev(q.tail)], true
}

func (q *Deque[T]) GetAt(idx int) (T, error) {
	if idx < 0 || idx >= q.count {
		var zero T
		return zero, ErrInvalidIndex
	}
	// bitwise modulus
	return q.data[(q.head+idx)&(len(q.data)-1)], nil
}

func (q *Deque[T]) RemoveAt(idx int) error {
	return nil
}

func (q *Deque[T]) InsertBefore(idx int, elem T) {

}

func (q *Deque[T]) InsertAfter(idx int, elem T) {

}

func (q *Deque[T]) Push(elem T) {
	q.PushFront(elem)
}

func (q *Deque[T]) Pop() (T, bool) {
	return q.PopFront()
}

func (q *Deque[T]) Peek() (T, bool) {
	return q.Front()
}

func (q *Deque[T]) Offer(elem T) {
	q.PushBack(elem)
}

func (q *Deque[T]) Poll() (T, bool) {
	return q.PopFront()
}

func (q *Deque[T]) Capacity() int {
	if q == nil {
		return 0
	}
	return cap(q.data)
}

func (q *Deque[T]) Len() int {
	if q == nil {
		return 0
	}
	return q.count
}

func (q *Deque[T]) AsSlice() []T {
	if q == nil || q.count == 0 {
		return make([]T, 0)
	}
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
		// Otherwise the buffer is wrapping around so the data from head position
		// until the end needs to be copied first, and then from the beginning of
		// the buffer until that tail
		temp := copy(newData, q.data[q.head:])
		copy(newData[temp:], q.data[:q.tail])
	}
	q.head = 0
	q.tail = q.count
	q.data = newData
}
