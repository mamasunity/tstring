package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter1(t *testing.T) {
	assert.Equal(t, "hello", Center("hello", 4, " "))
}

func TestCenter2(t *testing.T) {
	assert.Equal(t, "  hello   ", Center("hello", 10, " "))
}

func TestCenter3(t *testing.T) {
	assert.Equal(t, "12hello123", Center("hello", 10, "123"))
}

func TestExpandTabs1(t *testing.T) {
	assert.Equal(t, "a    bc    def", ExpandTabs("a\tbc\tdef", 4))
}

func TestExpandTabs2(t *testing.T) {
	assert.Equal(t, "a  bc  def", ExpandTabs("a\tbc\tdef", 2))
}

func TestExpandTabs3(t *testing.T) {
	assert.Equal(t, "中    文", ExpandTabs("中\t文", 4))
}

func TestFirstRuneLower1(t *testing.T) {
	assert.Equal(t, "aBC", FirstRuneToLower("ABC"))
}

func TestFirstRuneLower2(t *testing.T) {
	assert.Equal(t, "abc", FirstRuneToLower("abc"))
}

func TestFirstRuneToUpper1(t *testing.T) {
	assert.Equal(t, "ABC", FirstRuneToUpper("aBC"))
}

func TestFirstRuneToUpper2(t *testing.T) {
	assert.Equal(t, "ABC", FirstRuneToUpper("ABC"))
}
