package collections

import (
	"fmt"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	l := LinkedList[string]{}
	l.Add("hello")
	l.Add("world")
	l.Add("cool")

	fmt.Println(l)

	for i := l.Iterator(); i.Next(); {
		fmt.Println(i.Value())
	}

	for i := l.ReverseIterator(); i.Next(); {
		fmt.Println(i.Value())
	}
}
