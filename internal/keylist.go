package internal

type Element[K comparable, V any] struct {
	next  *Element[K, V]
	prev  *Element[K, V]
	Key   K
	Value V
}

func (e *Element[K, V]) Next() *Element[K, V] {
	return e.next
}

func (e *Element[K, V]) Prev() *Element[K, V] {
	return e.prev
}

// KeyList is a double linked list specialized for use by OrderedMap
// to store and maintain order.
type KeyList[K comparable, V any] struct {
	root Element[K, V]
}

func (l *KeyList[K, V]) Empty() bool {
	return l.root.next == nil
}

func (l *KeyList[K, V]) Front() *Element[K, V] {
	return l.root.next
}

func (l *KeyList[K, V]) Back() *Element[K, V] {
	return l.root.prev
}

func (l *KeyList[K, V]) Remove(e *Element[K, V]) {
	if e.prev == nil {
		l.root.next = e.next
	} else {
		e.prev.next = e.next
	}
	if e.next == nil {
		l.root.prev = e.prev
	} else {
		e.next.prev = e.prev
	}
	e.next = nil
	e.prev = nil
}

func (l *KeyList[K, V]) PushFront(key K, val V) *Element[K, V] {
	e := &Element[K, V]{Key: key, Value: val}
	if l.root.next == nil {
		l.root.next = e
		l.root.prev = e
		return e
	}

	e.next = l.root.next
	l.root.next.prev = e
	l.root.next = e
	return e
}

func (l *KeyList[K, V]) PushBack(key K, val V) *Element[K, V] {
	e := &Element[K, V]{Key: key, Value: val}
	if l.root.prev == nil {
		// It's the first element
		l.root.next = e
		l.root.prev = e
		return e
	}

	e.prev = l.root.prev
	l.root.prev.next = e
	l.root.prev = e
	return e
}
