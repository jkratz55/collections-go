package collections

import (
	"encoding/json"
	"reflect"

	"github.com/vmihailenco/msgpack/v5"
)

// Set is a collection that contains no duplicate elements. More formally, sets
// contain no pair of elements e1 and e2 such that e1 == e2. As implied by its name,
// Set models the mathematical set abstraction.
//
// Set makes no guarantees as to the iteration order of the set; in particular, it
// does not guarantee that the order will remain constant over time.
//
// The zero-value of Set is not usable. Set should be created and initialized using
// NewSet function.
//
// Set supports marshaling/unmarshalling for json and msgpack out of the box.
//
// Note: Set is not thread safe.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates and initializes a Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// Add adds the elements into the Set. If an element already exists it is effectively
// a no-op.
func (s *Set[T]) Add(vals ...T) {
	for _, val := range vals {
		s.data[val] = struct{}{}
	}
}

// Remove removes/deletes an element from the Set. If the element doesn't exist
// this is a no-op.
func (s *Set[T]) Remove(val T) {
	delete(s.data, val)
}

// Contains reads through the Set and returns a boolean value indicating if the
// provided value is in the Set.
func (s *Set[T]) Contains(val T) bool {
	_, ok := s.data[val]
	return ok
}

// Size returns the current number of elements in the Set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Equals returns a boolean indicating if the Set is equal to the provided Set.
func (s *Set[T]) Equals(other *Set[T]) bool {
	return reflect.DeepEqual(s.data, other.data)
}

// ForEach iterates through the Set passing the value to the provided function.
func (s *Set[T]) ForEach(fn func(val T)) {
	for key, _ := range s.data {
		fn(key)
	}
}

// Iter returns a stateful iterator for iterator over a Set.
//
// Note: Internally AsSlice is called to populate the SetIterator. In the vast
// majority of use cases this is unlikely to matter. However, if the Set contains
// a vast amount of data it may be more efficient to use ForEach with a closure.
func (s *Set[T]) Iter() *SetIterator[T] {
	return &SetIterator[T]{
		current: -1,
		data:    s.AsSlice(),
	}
}

// AsSlice converts a Set to a built-in Go slice.
func (s *Set[T]) AsSlice() []T {
	vals := make([]T, 0, len(s.data))
	for key, _ := range s.data {
		vals = append(vals, key)
	}
	return vals
}

// Clone does a deep copy and returns a new Set with the same elements/deque.
func (s *Set[T]) Clone() *Set[T] {
	other := NewSet[T]()
	for key, _ := range s.data {
		other.Add(key)
	}
	return other
}

// Union returns a new Set which contains all the elements in both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := s.Clone()
	for key, _ := range other.data {
		unionSet.Add(key)
	}
	return unionSet
}

// MarshalJSON marshals a Set into binary JSON representation
func (s *Set[T]) MarshalJSON() ([]byte, error) {
	elements := s.AsSlice()
	return json.Marshal(elements)
}

// UnmarshalJSON unmarshalls binary JSON representation of a Set into this instance
// of Set.
func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var raw []T
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.Add(raw...)
	return nil
}

// MarshalMsgpack marshals a Set into binary msgpack representation.
func (s *Set[T]) MarshalMsgpack() ([]byte, error) {
	elements := s.AsSlice()
	return msgpack.Marshal(elements)
}

// UnmarshalMsgpack unmarshalls binary msgpack representation of a Set into this
// instance of Set.
func (s *Set[T]) UnmarshalMsgpack(data []byte) error {
	var raw []T
	if err := msgpack.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.Add(raw...)
	return nil
}

// SetIterator is a stateful iterator for iterating through a Set.
type SetIterator[T comparable] struct {
	current int
	data    []T
}

// Next moves the iterator to the next element/value and returns a boolean
// indicating if there is a valid value.
func (s *SetIterator[T]) Next() bool {
	s.current++
	if s.current >= len(s.data) {
		return false
	}
	return true
}

// Value returns the current value.
func (s *SetIterator[T]) Value() T {
	return s.data[s.current]
}
