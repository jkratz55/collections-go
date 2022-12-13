package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		want *Set[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSet(), "NewSet()")
		})
	}
}

func TestSetIterator_Next(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    SetIterator[T]
		want bool
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Next(), "Next()")
		})
	}
}

func TestSetIterator_Value(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    SetIterator[T]
		want T
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Value(), "Value()")
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args[T comparable] struct {
		vals []T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.vals...)
		})
	}
}

func TestSet_AsSlice(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want []T
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.AsSlice(), "AsSlice()")
		})
	}
}

func TestSet_Clone(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want *Set[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Clone(), "Clone()")
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		val T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Contains(tt.args.val), "Contains(%v)", tt.args.val)
		})
	}
}

func TestSet_Equals(t *testing.T) {
	type args[T comparable] struct {
		other *Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Equals(tt.args.other), "Equals(%v)", tt.args.other)
		})
	}
}

func TestSet_ForEach(t *testing.T) {
	type args[T comparable] struct {
		fn func(val T)
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.ForEach(tt.args.fn)
		})
	}
}

func TestSet_Iter(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want *SetIterator[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Iter(), "Iter()")
		})
	}
}

func TestSet_MarshalJSON(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		s       Set[T]
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalJSON()
			if !tt.wantErr(t, err, fmt.Sprintf("MarshalJSON()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "MarshalJSON()")
		})
	}
}

func TestSet_MarshalMsgpack(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		s       Set[T]
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalMsgpack()
			if !tt.wantErr(t, err, fmt.Sprintf("MarshalMsgpack()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "MarshalMsgpack()")
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args[T comparable] struct {
		val T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.val)
		})
	}
}

func TestSet_Size(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want int
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Size(), "Size()")
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args[T comparable] struct {
		other *Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want *Set[T]
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.Union(tt.args.other), "Union(%v)", tt.args.other)
		})
	}
}

func TestSet_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	type testCase[T comparable] struct {
		name    string
		s       Set[T]
		args    args
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.s.UnmarshalJSON(tt.args.data), fmt.Sprintf("UnmarshalJSON(%v)", tt.args.data))
		})
	}
}

func TestSet_UnmarshalMsgpack(t *testing.T) {
	type args struct {
		data []byte
	}
	type testCase[T comparable] struct {
		name    string
		s       Set[T]
		args    args
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.s.UnmarshalMsgpack(tt.args.data), fmt.Sprintf("UnmarshalMsgpack(%v)", tt.args.data))
		})
	}
}
