package collections

// Queue represents an unbounded first-in-first-out queue of elements.
type Queue[T any] struct {
	data []T
}

// Offer inserts/adds an element at the end of the queue.
func (q *Queue[T]) Offer(val T) {
	q.data = append(q.data, val)
}

// Poll retrieves the first element in the queue, returning the value, proceeding
// by removing it from the queue. If the queue is empty the zero value will be
// returned along with an ok value of false.
func (q *Queue[T]) Poll() (val T, ok bool) {
	if len(q.data) == 0 {
		return
	}
	val = q.data[0]
	ok = true
	q.data = q.data[1:]
	return
}

// Peek retrieves and returns the first element in the queue. Unlike Poll, the
// element is not removed. If the queue is empty the zero value will be returned
// along with an ok value of false.
func (q *Queue[T]) Peek() (val T, ok bool) {
	if len(q.data) == 0 {
		return
	}
	return q.data[0], true
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// Empty returns a boolean value indicating if the Queue is empty, true if it is
// empty, otherwise false.
func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}
