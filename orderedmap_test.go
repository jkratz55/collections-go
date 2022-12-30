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
