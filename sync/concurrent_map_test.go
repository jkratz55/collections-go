package sync

import (
	"fmt"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := NewConcurrentMap[string, int](16, StringHasher())
	m.Set("test", 1)
	m.Set("Awkbar!", 343)

	fmt.Println(m.Get("test"))
}
