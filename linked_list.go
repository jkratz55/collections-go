package collections

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}

type LinkedList[T any] struct {
}
