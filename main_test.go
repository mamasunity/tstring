package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	assert.Equal(t, "hello", Center("hello", 4, " "))
	assert.Equal(t, "  hello   ", Center("hello", 10, " "))
	assert.Equal(t, "12hello123", Center("hello", 10, "123"))
}

func TestExpandTabs(t *testing.T) {
	assert.Equal(t, "a    bc    def", ExpandTabs("a\tbc\tdef", 4))
	assert.Equal(t, "a  bc  def", ExpandTabs("a\tbc\tdef", 2))
	assert.Equal(t, "中    文", ExpandTabs("中\t文", 4))
}

func TestFirstRuneLower(t *testing.T) {
	assert.Equal(t, "aBC", FirstRuneToLower("ABC"))
	assert.Equal(t, "abc", FirstRuneToLower("abc"))
}

func TestFirstRuneToUpper(t *testing.T) {
	assert.Equal(t, "ABC", FirstRuneToUpper("aBC"))
	assert.Equal(t, "ABC", FirstRuneToUpper("ABC"))
}

func TestInsert(t *testing.T) {
	assert.Equal(t, "Aabc", Insert("abc", "A", 0))
	assert.Equal(t, "aAbc", Insert("abc", "A", 1))

	assert.Panics(t, func() { Insert("abc", "A", 10) }, "Index must be included range.")
}
