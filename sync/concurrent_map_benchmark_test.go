package sync

import (
	"fmt"
	"testing"
)

func BenchmarkConcurrentMapWrites(b *testing.B) {
	b.ReportAllocs()
	m := NewConcurrentMap[string, int](DefaultShards, StringHasher())

	for i := 0; i < 10000; i++ {
		m.Set(fmt.Sprintf("%d", i), i)
	}
}

func BenchmarkMapWrites(b *testing.B) {
	b.ReportAllocs()
	m := Map[string, int]{}

	for i := 0; i < 10000; i++ {
		m.Store(fmt.Sprintf("%d", i), i)
	}
}

func BenchmarkConcurrentMapReads(b *testing.B) {
	m := NewConcurrentMap[string, int](DefaultShards, StringHasher())

	for i := 0; i < 100000; i++ {
		m.Set(fmt.Sprintf("%d", i), i)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < 100000; i++ {
		_, _ = m.Get(fmt.Sprintf("%d", i))
	}
}

func BenchmarkMapReads(b *testing.B) {
	m := Map[string, int]{}

	for i := 0; i < 100000; i++ {
		m.Store(fmt.Sprintf("%d", i), i)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < 100000; i++ {
		_, _ = m.Load(fmt.Sprintf("%d", i))
	}
}
