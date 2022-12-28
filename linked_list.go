package collections

import (
	"encoding/json"

	"github.com/vmihailenco/msgpack/v5"
)

// Node represents an element in a LinkedList.
//
// The zero-value of Node is not usable and is not attended to be instantiated
// outside this package.
type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

func (n Node[T]) Next() *Node[T] {
	return n.next
}

func (n Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n Node[T]) Value() T {
	return n.val
}

func (n Node[T]) HasNext() bool {
	return n.next == nil
}

func (n Node[T]) HasPrev() bool {
	return n.prev == nil
}

// LinkedList is a doubly linked list implementation.
//
// LinkedList is inspired by Java's LinkedList implementation and Deque interface.
// LinkedList is a linear double linked list that supports insertion and removal
// at both ends. LinkedList can be used as a list, stack, or queue.
//
// When using LinkedList as a stack the front/head represents the top of the stack.
// Push pushes elements to the front of the list and pop/peek reads from the front
// of the LinkedList.
//
// When using LinkedList as a queue the front/head represents the beginning of the
// queue. Offer pushes elements to the back of the LinkedList and Poll/Peek read
// from the front of the LinkedList.
//
// For obvious reasons the same instance of a LinkedList cannot be used as both a
// stack and queue, or a list. The following below shows the API mappings for the
// stack and queue API methods.
//
//		Stack API:
//	  - Push -> PushFront
//	  - Pop -> PopFront
//	  - Peek -> Front
//
//		Queue API:
//	  - Offer -> PushBack
//	  - Poll -> PopFront
//	  - Peek -> Front
//
// Note: When working with indexes LinkedList following the same rules as slices
// and arrays in Go. If you provide an invalid index a panic will occur.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// PushFront adds the element at the beginning of the list
func (l *LinkedList[T]) PushFront(val T) {
	// If head is nil the list is empty, so we can simply call PushBack to add
	// element to end of the list
	if l.head == nil {
		l.PushBack(val)
		return
	}
	// Otherwise we need to create a new node and set it has the head of the list
	// while updating the linkage between the previous head and new head
	next := l.head
	newNode := &Node[T]{
		val:  val,
		prev: nil,
		next: next,
	}
	next.prev = newNode
	l.head = newNode
	l.size++
}

// PushBack adds the element at the end of the list
func (l *LinkedList[T]) PushBack(val T) {
	// LinkedList is empty
	if l.head == nil {
		l.head = &Node[T]{
			prev: nil,
			next: nil,
			val:  val,
		}
		l.tail = l.head
		l.size++
		return
	}

	// Otherwise LinkedList isn't empty and we need to add new element at the tail
	prevTail := l.tail
	newNode := &Node[T]{
		prev: prevTail,
		next: nil,
		val:  val,
	}
	prevTail.next = newNode
	l.tail = newNode
	l.size++
}

// Front returns the first (head) element of the LinkedList.
//
// If the list is empty the zero value is returned and a boolean value of false.
func (l *LinkedList[T]) Front() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.val, true
}

// Back returns the last (tail) element of the LinkedList
//
// If the list is empty the zero value is returned and a boolean value of false.
func (l *LinkedList[T]) Back() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.val, true
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	prevHead := l.head

	// If head doesn't have a next value then popping head will result in the
	// LinkedList being empty.
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		l.size--
		return prevHead.val, true
	} else {
		// Otherwise there are more than one element in the LinkedList and the pointer
		// to head needs to be updated.
		l.head = l.head.next
		l.head.prev = nil
		l.size--
		return prevHead.val, true
	}
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	prevTail := l.tail

	// If tail doesn't have a prev value then popping tail will result in the
	// LinkedList being empty.
	if l.tail.prev == nil {
		l.tail = nil
		l.head = nil
		l.size--
		return prevTail.val, true
	} else {
		// Otherwise there are more than one element in the LinkedList and the pointer
		// to tail needs to be updated.
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
		return prevTail.val, true
	}
}

func (l *LinkedList[T]) InsertBefore(idx int, val T) {
	if idx > l.size-1 || idx < 0 {
		panic("illegal index: index out of bounds")
	}
	newNode := &Node[T]{
		val: val,
	}
	// Handle case of inserting at head
	if idx == 0 {
		// If LinkedList is empty we can just use the Add method to add element
		if l.head == nil {
			l.Add(val)
			return
		}
		prevHead := l.head
		l.head = newNode
		newNode.next = prevHead
		l.size++
		return
	}
	current := l.getAt(idx)
	prevNode := current.prev
	newNode.next = current
	prevNode.next = newNode
	l.size++
}

func (l *LinkedList[T]) InsertAfter(idx int, val T) {

}

func (l *LinkedList[T]) GetAt(idx int) T {
	return l.getAt(idx).val
}

func (l *LinkedList[T]) RemoveAt(idx int) {
	node := l.getAt(idx)
	if node.prev == nil {

	}
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Empty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) AsSlice() []T {
	data := make([]T, 0, l.size)
	for i := l.Iterator(); i.Next(); {
		data = append(data, i.Value())
	}
	return data
}

// ----------------------------------------------------------------------------
// Stack API
// ----------------------------------------------------------------------------

// Push pushes an element at the end/top of the LinkedList
//
// Push is the equivalent of PushFront but provides a familiar style API for working
// with stacks.
func (l *LinkedList[T]) Push(val T) {
	l.PushFront(val)
}

// Pop fetches the last element at the end/top of the LinkedList and then removes
// it. If the LinkedList is empty the zero value and ok value of false is returned.
//
// Pop is the equivalent of PopBack but provides a familiar style API for working
// with stacks.
func (l *LinkedList[T]) Pop() (val T, ok bool) {
	return l.PopFront()
}

// Peek fetches the last element at the end/top of the LinkedList but does not
// remove it. If the LinkedList is empty the zero value and ok value of false
// is returned.
//
// Peek is the equivalent of Back but provides a familiar style API for working
// with stacks.
func (l *LinkedList[T]) Peek() (val T, ok bool) {
	return l.Front()
}

// ----------------------------------------------------------------------------
// Queue API
// ----------------------------------------------------------------------------

func (l *LinkedList[T]) Offer(val T) {
	l.PushBack(val)
}

func (l *LinkedList[T]) Poll() (val T, ok bool) {
	return l.PopFront()
}

// ----------------------------------------------------------------------------

func (l *LinkedList[T]) Iterator() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		list: l,
	}
}

func (l *LinkedList[T]) ReverseIterator() *ReverseLinkedListIterator[T] {
	return &ReverseLinkedListIterator[T]{
		list: l,
	}
}

func (l *LinkedList[T]) MarshalJSON() ([]byte, error) {
	data := l.AsSlice()
	return json.Marshal(data)
}

func (l *LinkedList[T]) UnmarshalJSON(data []byte) error {
	var raw []T
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for _, val := range raw {
		l.Add(val)
	}
	return nil
}

func (l *LinkedList[T]) MarshalMsgpack() ([]byte, error) {
	data := l.AsSlice()
	return msgpack.Marshal(data)
}

func (l *LinkedList[T]) UnmarshalMsgpack(data []byte) error {
	var raw []T
	if err := msgpack.Unmarshal(data, &raw); err != nil {
		return err
	}
	for _, val := range raw {
		l.Add(val)
	}
	return nil
}

func (l *LinkedList[T]) getAt(idx int) *Node[T] {
	if idx < 0 || idx > l.size-1 {
		panic("illegal index: index out of bounds")
	}
	current := l.head
	if idx == 0 {
		return current
	}
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current
}

type LinkedListIterator[T any] struct {
	list    *LinkedList[T]
	current *Node[T]
}

func (i *LinkedListIterator[T]) Next() bool {
	// The list is empty
	if i.list.size == 0 {
		return false
	}
	// If first time next has been called load head into current to begin iterating
	// starting with head
	if i.current == nil {
		i.current = i.list.head
		return true
	}
	// If there is a next Node move the pointer of current to next and return true.
	if i.current.next != nil {
		i.current = i.current.next
		return true
	}
	// If we've reached this point then there is no more elements in the LinkedList
	return false
}

func (i *LinkedListIterator[T]) Value() T {
	return i.current.val
}

type ReverseLinkedListIterator[T any] struct {
	list    *LinkedList[T]
	current *Node[T]
}

func (i *ReverseLinkedListIterator[T]) Next() bool {
	// The list is empty
	if i.list.size == 0 {
		return false
	}
	// If first time next has been called load head into current to begin iterating
	// starting with head
	if i.current == nil {
		i.current = i.list.tail
		return true
	}
	// If there is a next Node move the pointer of current to next and return true.
	if i.current.prev != nil {
		i.current = i.current.prev
		return true
	}
	// If we've reached this point then there is no more elements in the LinkedList
	return false
}

func (i *ReverseLinkedListIterator[T]) Value() T {
	return i.current.val
}

// ----------- REMOVE BELOW ------------ //

// Add adds the element at the end of the list
//
// Add is effectively an alias for PushBack
func (l *LinkedList[T]) Add(val T) {
	l.PushBack(val)
}

func (l *LinkedList[T]) AddAt(idx int, val T) {
	if idx > l.size-1 || idx < 0 {
		panic("illegal index: index out of bounds")
	}
	newNode := &Node[T]{
		val: val,
	}
	// Handle case of inserting at head
	if idx == 0 {
		// If LinkedList is empty we can just use the Add method to add element
		if l.head == nil {
			l.Add(val)
			return
		}
		prevHead := l.head
		l.head = newNode
		newNode.next = prevHead
		l.size++
		return
	}
	current := l.getAt(idx)
	prevNode := current.prev
	newNode.next = current
	prevNode.next = newNode
	l.size++
}

func (l *LinkedList[T]) GetFirst() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.val, true
}

func (l *LinkedList[T]) GetLast() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.val, true
}

func (l *LinkedList[T]) PopFirst() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	prevHead := l.head

	// If head doesn't have a next value then popping head will result in the
	// LinkedList being empty.
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		l.size--
		return prevHead.val, true
	} else {
		// Otherwise there are more than one element in the LinkedList and the pointer
		// to head needs to be updated.
		l.head = l.head.next
		l.head.prev = nil
		l.size--
		return prevHead.val, true
	}
}

func (l *LinkedList[T]) PopLast() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	prevTail := l.tail

	// If tail doesn't have a prev value then popping tail will result in the
	// LinkedList being empty.
	if l.tail.prev == nil {
		l.tail = nil
		l.head = nil
		l.size--
		return prevTail.val, true
	} else {
		// Otherwise there are more than one element in the LinkedList and the pointer
		// to tail needs to be updated.
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
		return prevTail.val, true
	}
}
