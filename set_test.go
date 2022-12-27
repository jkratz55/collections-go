package collections

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

func TestNewSet(t *testing.T) {
	set := NewSet[string]()
	assert.NotNil(t, set)
}

func TestSet_Add(t *testing.T) {
	set := NewSet[string]()

	data := []string{"billy", "bob", "the", "great"}
	set.Add(data...)

	assert.Equal(t, 4, len(set.data))
	for _, val := range data {
		_, ok := set.data[val]
		assert.True(t, ok)
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("apple", "orange", "banana", "grape")

	s2 := NewSet[string]()
	s2.Add("pizza", "hamburger", "pie", "apple")

	s3 := s1.Union(s2)

	expected := map[string]struct{}{
		"apple":     {},
		"orange":    {},
		"banana":    {},
		"grape":     {},
		"pizza":     {},
		"hamburger": {},
		"pie":       {},
	}
	assert.Equal(t, expected, s3.data)
}

func TestSet_Remove(t *testing.T) {
	set := NewSet[string]()
	set.Remove("hello") // noop
	set.data["hello"] = struct{}{}
	set.data["hola"] = struct{}{}

	set.Remove("hello")
	_, ok := set.data["hello"]
	assert.False(t, ok)
}

func TestSet_Contains(t *testing.T) {
	set := NewSet[string]()
	set.Add("hello", "world", "billy bob")

	assert.False(t, set.Contains("batman"))
	assert.True(t, set.Contains("hello"))
	assert.True(t, set.Contains("world"))
	assert.False(t, set.Contains("billy"))
}

func TestSet_Size(t *testing.T) {
	set := NewSet[string]()
	assert.Equal(t, len(set.data), set.Size())
	assert.Equal(t, 0, set.Size())

	set.Add("hello", "world", "billy bob")

	assert.Equal(t, len(set.data), set.Size())
	assert.Equal(t, 3, set.Size())
}

func TestSet_Equals(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	s2 := NewSet[string]()
	s2.Add("pizza", "tacos", "hamburger")

	assert.True(t, s1.Equals(s2))

	s3 := NewSet[string]()
	s3.Add("hamburger", "tacos", "pizza")
	assert.True(t, s1.Equals(s3))

	s4 := NewSet[string]()
	s4.Add("pizza", "pineapples")
	assert.False(t, s1.Equals(s4))
}

func TestSet_Clone(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	s2 := s1.Clone()
	assert.Equal(t, s1.data, s2.data)

	s2.Add("hello")
	assert.NotEqual(t, s1.data, s2.data)
}

func TestSet_AsSlice(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	data := s1.AsSlice()
	assert.ElementsMatch(t, []string{"pizza", "tacos", "hamburger"}, data)
}

func TestSet_ForEach(t *testing.T) {
	var actual []string

	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	s1.ForEach(func(val string) {
		actual = append(actual, val)
	})

	assert.Equal(t, 3, len(actual))
	assert.ElementsMatch(t, []string{"pizza", "tacos", "hamburger"}, actual)
}

func TestSet_Iter(t *testing.T) {
	var actual []string

	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	for iter := s1.Iter(); iter.Next(); {
		actual = append(actual, iter.Value())
	}

	assert.Equal(t, 3, len(actual))
	assert.ElementsMatch(t, []string{"pizza", "tacos", "hamburger"}, actual)
}

func TestSet_MarshalJSON(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	_, err := json.Marshal(s1)
	assert.NoError(t, err)
}

func TestSet_UnmarshalJSON(t *testing.T) {
	input := "[\"pizza\", \"tacos\", \"hamburger\"]"
	set := NewSet[string]()

	err := json.Unmarshal([]byte(input), set)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"pizza", "tacos", "hamburger"}, set.AsSlice())
}

func TestSet_MarshalMsgpack(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("pizza", "tacos", "hamburger")

	_, err := msgpack.Marshal(s1)
	assert.NoError(t, err)
}

func TestSet_UnmarshalMsgpack(t *testing.T) {
	input := []byte{147, 169, 104, 97, 109, 98, 117, 114, 103, 101, 114, 165, 112, 105, 122, 122, 97, 165, 116, 97, 99, 111, 115}
	set := NewSet[string]()

	err := msgpack.Unmarshal(input, set)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"pizza", "tacos", "hamburger"}, set.AsSlice())
}
