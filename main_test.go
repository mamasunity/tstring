package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter1(t *testing.T) {
	assert.Equal(t, Center("hello", 4, " "), "hello")
}

func TestCenter2(t *testing.T) {
	assert.Equal(t, Center("hello", 10, " "), "  hello   ")
}

func TestCenter3(t *testing.T) {
	assert.Equal(t, Center("hello", 10, "123"), "12hello123")
}

func TestExpandTabs1(t *testing.T) {
	assert.Equal(t, ExpandTabs("a\tbc\tdef", 4), "a    bc    def")
}

func TestExpandTabs2(t *testing.T) {
	assert.Equal(t, ExpandTabs("a\tbc\tdef", 2), "a  bc  def")
}

func TestExpandTabs3(t *testing.T) {
	assert.Equal(t, ExpandTabs("中\t文", 4), "中    文")
}

func TestFirstRuneLower1(t *testing.T) {
	assert.Equal(t, FirstRuneToLower("ABC"), "aBC")
}

func TestFirstRuneLower2(t *testing.T) {
	assert.Equal(t, FirstRuneToLower("abcde"), "abcde")
}
