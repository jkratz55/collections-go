package collections

import (
	"errors"
)

var (
	// ErrEmptyStack is a sentinel error value indicating the operation cannot
	// be performed on an empty stack.
	ErrEmptyStack = errors.New("stack is empty")
)

// Stack represents a last-in-first-out (LIFO) stack of elements.
//
// Under the hood Stack simply wraps Deque to provide a more friendly API for
// working with stacks.
//
// The zero-value of Stack is not usable. In order to properly initialize and use
// Stack use the NewStack function.
type Stack[T any] struct {
	data *Deque[T]
}

// NewStack creates and initializes a new empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: NewDeque[T](),
	}
}

// Push pushes an element to the top of the stack.
func (s *Stack[T]) Push(val T) {
	s.data.PushBack(val)
}

// Pop removes and returns the element at the top of the stack. In the event the
// Stack is empty a non-nil error value of ErrEmptyStack is returned.
func (s *Stack[T]) Pop() (val T, err error) {
	val, ok := s.data.PopBack()
	if !ok {
		var zero T
		return zero, ErrEmptyStack
	}
	return val, nil
}

// Peek returns the element at the top of the stack but doesn't remove it. Peek
// additionally returns a boolean indicating if a value was returned. If the stack
// was empty false we be returned.
func (s *Stack[T]) Peek() (val T, ok bool) {
	return s.data.Back()
}

// Size returns the current size (number of elements) on the stack.
func (s *Stack[T]) Size() int {
	return s.data.Len()
}

// Empty returns a boolean indicating if the stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.data.IsEmpty()
}

// PopEach will pop each element off the stack one by one until the stack is empty.
// As each element is popped off the stack the provided function is executed with
// the popped off value.
func (s *Stack[T]) PopEach(fn func(val T)) {
	for val, ok := s.data.PopBack(); ok; val, ok = s.data.PopBack() {
		fn(val)
	}
}
