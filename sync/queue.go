package sync

import (
	"context"
)

// BlockingQueue is an implementation of a bounded queue. BlockingQueue has a finite
// capacity that is defined when initializing an instance of BlockingQueue using the
// function NewBlockingQueue.
//
// BlockingQueue provides a Java-like API providing Offer, TryOffer, Poll, and TryPoll
// methods. Offer and Poll are blocking, while TryOffer and TryPoll are non-blocking.
//
// The zero-value of BlockingQueue is not usable. A BlockingQueue should be created
// using the NewBlockingQueue function.
type BlockingQueue[T any] struct {
	data     chan T
	capacity int
}

// NewBlockingQueue instantiates and initializes a new BlockingQueue. The supplied
// capacity must be a value greater than or equal to 1. A capacity of 0 or negative
// values will result in a panic.
func NewBlockingQueue[T any](capacity int) *BlockingQueue[T] {
	if capacity < 1 {
		panic("capacity cannot be less than 1")
	}
	return &BlockingQueue[T]{
		data:     make(chan T, capacity),
		capacity: capacity,
	}
}

// Offer offers an element to the queue blocking until the element is accepted
// and added to the queue, or the context is done. If the context is done before
// the element is accepted by the queue a non-nil error value will be returned.
func (bq *BlockingQueue[T]) Offer(ctx context.Context, val T) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case bq.data <- val:
		return nil
	}
}

// TryOffer offers an element to the queue without blocking. If the queue is at
// capacity and not accepting elements TryPoll will immediately return false
// indicating the value was not added, returns true if the element was added
// to the queue.
func (bq *BlockingQueue[T]) TryOffer(val T) bool {
	select {
	case bq.data <- val:
		return true
	default:
		return false
	}
}

// Poll retrieves a single item from queue blocking until an element is available
// or the context is done. In the context is done before an element arrives a
// non-nil error value is returned. Otherwise, the value and a nil error value
// are returned.
func (bq *BlockingQueue[T]) Poll(ctx context.Context) (T, error) {
	var defaultVal T
	select {
	case <-ctx.Done():
		return defaultVal, ctx.Err()
	case val := <-bq.data:
		return val, nil
	}
}

// TryPoll retrieves a single element from the queue without blocking. If there
// are no elements available to poll/consume TryPoll returns immediately with
// the zero value of the element and a boolean value of false. If a value was
// polled from the queue it is returned along with a boolean value of true to
// indicate an element was successfully polled.
func (bq *BlockingQueue[T]) TryPoll() (T, bool) {
	var defaultVal T
	select {
	case val := <-bq.data:
		return val, true
	default:
		return defaultVal, false
	}
}

// Clear removes all the elements from the queue.
//
// Clear is blocking and will run as long as elements are in the queue. If elements
// are being offered while the queue is being cleared it may run indefinitely.
func (bq *BlockingQueue[T]) Clear() {
	for len(bq.data) > 0 {
		<-bq.data
	}
}

// Size returns the count of elements in the queue.
func (bq *BlockingQueue[T]) Size() int {
	return len(bq.data)
}

// Empty returns a boolean value indicating if the queue is empty.
func (bq *BlockingQueue[T]) Empty() bool {
	return len(bq.data) == 0
}

// Capacity returns the capacity of the queue.
func (bq *BlockingQueue[T]) Capacity() int {
	return bq.capacity
}

// CapacityRemaining returns the remaining capacity of the queue by taking the
// capacity and subtracting the elements in the queue.
func (bq *BlockingQueue[T]) CapacityRemaining() int {
	return bq.capacity - len(bq.data)
}
