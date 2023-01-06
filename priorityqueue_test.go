package collections

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int]()
	pq.Push(0, 1)
	pq.Push(1, 1)
	pq.Push(2, 10)
	pq.Push(3, 5)

	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	pq.Push(4, 30)
	fmt.Println(pq.Peek())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
}
