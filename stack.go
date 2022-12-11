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
// The zero-value of Stack is not usable. In order to properly initialize and use
// Stack use the NewStack function.
type Stack[T any] struct {
	data []T
}

// NewStack creates and initializes a new empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Push pushes an element to the top of the stack.
func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

// Pop removes and returns the element at the top of the stack. In the event the
// Stack is empty a non-nil error value of ErrEmptyStack is returned.
func (s *Stack[T]) Pop() (val T, err error) {
	if len(s.data) == 0 {
		return val, ErrEmptyStack
	}
	val = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

// Peek returns the element at the top of the stack but doesn't remove it. Peek
// additionally returns a boolean indicating if a value was returned. If the stack
// was empty false we be returned.
func (s *Stack[T]) Peek() (val T, ok bool) {
	if len(s.data) == 0 {
		return val, false
	}
	val = s.data[len(s.data)-1]
	return val, true
}

// Size returns the current size (number of elements) on the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Empty returns a boolean indicating if the stack is empty.
func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

// PopEach will pop each element off the stack one by one until the stack is empty.
// As each element is popped off the stack the provided function is executed with
// the popped off value.
func (s *Stack[T]) PopEach(fn func(val T)) {
	if len(s.data) == 0 {
		return
	}
	for {
		val, err := s.Pop()
		if err != nil {
			return
		}
		fn(val)
	}
}
