package collections

type Set[T comparable] struct {
	data map[T]struct{}
}
