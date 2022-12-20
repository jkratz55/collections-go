package sync

import (
	"fmt"
	"testing"
)

func TestMap_Load(t *testing.T) {
	m := Map[string, string]{}
	m.Store("hello", "world")
	m.Store("awesome", "sauce")

	fmt.Println(m.Load("hello"))
}
