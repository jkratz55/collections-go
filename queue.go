package collections

// Queue represents an unbounded first-in-first-out queue of elements.
//
// Under the hood Queue simply wraps Deque to provide a more friendly API for
// working with queues.
type Queue[T any] struct {
	deque *Deque[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		deque: NewDeque[T](),
	}
}

// Offer inserts/adds an element at the end of the queue.
func (q *Queue[T]) Offer(val T) {
	q.deque.PushBack(val)
}

// Poll retrieves the first element in the queue, returning the value, proceeding
// by removing it from the queue. If the queue is empty the zero value will be
// returned along with an ok value of false.
func (q *Queue[T]) Poll() (val T, ok bool) {
	return q.deque.PopFront()
}

// Peek retrieves and returns the first element in the queue. Unlike Poll, the
// element is not removed. If the queue is empty the zero value will be returned
// along with an ok value of false.
func (q *Queue[T]) Peek() (val T, ok bool) {
	return q.deque.Front()
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return q.deque.Len()
}

// Empty returns a boolean value indicating if the Queue is empty, true if it is
// empty, otherwise false.
func (q *Queue[T]) Empty() bool {
	return q.deque.Len() == 0
}
