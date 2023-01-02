package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap(t *testing.T) {

	om := NewOrderedMap[string, string]()
	assert.True(t, om.Set("Hello", "World"))
	assert.True(t, om.Set("Billy", "Bob"))
	assert.True(t, om.Set("John", "Doe"))
	assert.True(t, om.Set("Jane", "Doe"))
	assert.True(t, om.Set("Sleepy", "Joe"))
	assert.True(t, om.Set("Orange", "Man"))
	assert.False(t, om.Set("Orange", "Man"))

	val, ok := om.Get("Billy")
	assert.True(t, ok)
	assert.Equal(t, "Bob", val)

	val, ok = om.Get("Hello")
	assert.True(t, ok)
	assert.Equal(t, "World", val)

	val, ok = om.Get("AHHHHH")
	assert.False(t, ok)

	val = om.GetOrDefault("AHHHHH", "default")
	assert.Equal(t, "default", val)

	assert.Equal(t, "Man", om.GetOrDefault("Orange", "NOOOOO"))

	assert.True(t, om.Contains("Orange"))
	assert.False(t, om.Contains("AHHHHH"))

	assert.Equal(t, 6, om.Size())

	assert.True(t, om.Delete("Jane"))
	assert.False(t, om.Delete("Jane"))

	assert.Equal(t, 5, om.Size())

	expectedKeys := []string{"Hello", "Billy", "John", "Sleepy", "Orange"}
	expectedValues := []string{"World", "Bob", "Doe", "Joe", "Man"}

	var actualKeys []string
	var actualValues []string

	om.ForEach(func(key string, val string) {
		actualKeys = append(actualKeys, key)
		actualValues = append(actualValues, val)
	})
	assert.Equal(t, expectedKeys, actualKeys)
	assert.Equal(t, expectedValues, actualValues)

	assert.Equal(t, expectedKeys, om.Keys())

	expectedKeys = []string{"Orange", "Sleepy", "John", "Billy", "Hello"}
	expectedValues = []string{"Man", "Joe", "Doe", "Bob", "World"}
	actualKeys = make([]string, 0)
	actualValues = make([]string, 0)

	om.ForEachReverse(func(key string, val string) {
		actualKeys = append(actualKeys, key)
		actualValues = append(actualValues, val)
	})
	assert.Equal(t, expectedKeys, actualKeys)
	assert.Equal(t, expectedValues, actualValues)
}

func TestOrderedMap_Set(t *testing.T) {
	om := NewOrderedMap[string, string]()
	assert.True(t, om.Set("Hello", "World"))

	val, ok := om.data["Hello"]
	assert.True(t, ok)
	assert.Equal(t, "World", val.Value)
	assert.Equal(t, "Hello", om.keys.Front().Key)
	assert.Equal(t, "World", om.keys.Front().Value)

	assert.False(t, om.Set("Hello", "Cow"))
	assert.Equal(t, "Cow", val.Value)
	assert.Equal(t, "Hello", om.keys.Front().Key)
	assert.Equal(t, "Cow", om.keys.Front().Value)

	assert.True(t, om.Set("Mellow", "Yellow"))
	elem := om.keys.Front().Next()
	assert.Equal(t, "Mellow", elem.Key)
	assert.Equal(t, "Yellow", elem.Value)
}

func TestOrderedMap_Contains(t *testing.T) {
	om := NewOrderedMap[string, string]()
	assert.False(t, om.Contains("hello"))

	om.Set("hello", "world")
	assert.True(t, om.Contains("hello"))
}

func TestOrderedMap_Get(t *testing.T) {
	om := NewOrderedMap[string, string]()

	val, ok := om.Get("hello")
	assert.False(t, ok)
	assert.Equal(t, "", val)

	assert.True(t, om.Set("hello", "world"))
	val, ok = om.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, "world", val)
}

func TestOrderedMap_GetOrDefault(t *testing.T) {
	om := NewOrderedMap[string, string]()

	val := om.GetOrDefault("hello", "DEFAULT")
	assert.Equal(t, "DEFAULT", val)

	assert.True(t, om.Set("hello", "world"))
	val = om.GetOrDefault("hello", "DEFAULT")
	assert.Equal(t, "world", val)
}

func TestOrderedMap_Delete(t *testing.T) {
	om := NewOrderedMap[string, string]()

	assert.True(t, om.Set("1", "hello1"))
	assert.True(t, om.Set("2", "hello2"))
	assert.True(t, om.Set("3", "hello3"))

	assert.True(t, om.Delete("2"))
	assert.Equal(t, 2, len(om.data))

	root := om.keys.Front()
	last := root.Next()

	assert.Equal(t, root, om.data["1"])
	assert.Equal(t, last, om.data["3"])
	assert.Equal(t, last, root.Next())
	assert.Equal(t, root, last.Prev())
}

func TestOrderedMap_Size(t *testing.T) {
	om := NewOrderedMap[string, string]()
	assert.Equal(t, 0, om.Size())

	assert.True(t, om.Set("1", "hello1"))
	assert.True(t, om.Set("2", "hello2"))
	assert.True(t, om.Set("3", "hello3"))

	assert.Equal(t, 3, om.Size())
}

func TestOrderedMap_Keys(t *testing.T) {
	om := NewOrderedMap[string, string]()

	assert.True(t, om.Set("1", "hello1"))
	assert.True(t, om.Set("2", "hello2"))
	assert.True(t, om.Set("3", "hello3"))

	assert.Equal(t, []string{"1", "2", "3"}, om.Keys())
}
