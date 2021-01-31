package tstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	cases := []struct {
		expect string
		str    string
		length int
		pad    string
	}{
		{expect: "hello", str: "hello", length: 4, pad: " "},
		{expect: "  hello   ", str: "hello", length: 10, pad: " "},
		{expect: "12hello123", str: "hello", length: 10, pad: "123"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, Center(c.str, c.length, c.pad))
	}
}

func TestExpandTabs(t *testing.T) {
	cases := []struct {
		expect  string
		str     string
		tabSize int
	}{
		{expect: "a    bc    def", str: "a\tbc\tdef", tabSize: 4},
		{expect: "a  bc  def", str: "a\tbc\tdef", tabSize: 2},
		{expect: "中    文", str: "中\t文", tabSize: 4},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, ExpandTabs(c.str, c.tabSize))
	}
}

func TestFirstRuneLower(t *testing.T) {
	cases := []struct {
		expect string
		str    string
	}{
		{expect: "aBC", str: "ABC"},
		{expect: "abc", str: "abc"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, FirstRuneToLower(c.str))
	}
}

func TestFirstRuneToUpper(t *testing.T) {
	cases := []struct {
		expect string
		str    string
	}{
		{expect: "ABC", str: "aBC"},
		{expect: "ABC", str: "ABC"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, FirstRuneToUpper(c.str))
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		expect string
		dst    string
		src    string
		index  int
	}{
		{expect: "Aabc", dst: "abc", src: "A", index: 0},
		{expect: "aAbc", dst: "abc", src: "A", index: 1},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, Insert(c.dst, c.src, c.index))
	}

	assert.Panics(t, func() { Insert("abc", "A", 10) }, "Index must be included range.")
}

func TestLastPartition(t *testing.T) {
	cases := []struct {
		expectedHead  string
		expectedMatch string
		expectedTail  string
		str           string
		sep           string
	}{
		{expectedHead: "hel", expectedMatch: "l", expectedTail: "o", str: "hello", sep: "l"},
		{expectedHead: "hello", expectedMatch: "", expectedTail: "", str: "hello", sep: "x"},
	}

	for _, c := range cases {
		head, match, tail := LastPartition(c.str, c.sep)
		assert.Equal(t, c.expectedHead, head)
		assert.Equal(t, c.expectedMatch, match)
		assert.Equal(t, c.expectedTail, tail)
	}
}

func TestLeftJustify(t *testing.T) {
	cases := []struct {
		expect string
		str    string
		length int
		pad    string
	}{
		{expect: "hello", str: "hello", length: 4, pad: " "},
		{expect: "hello     ", str: "hello", length: 10, pad: " "},
		{expect: "hello12312", str: "hello", length: 10, pad: "123"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, LeftJustify(c.str, c.length, c.pad))
	}
}

func TestPartition(t *testing.T) {
	cases := []struct {
		expectedHead  string
		expectedMatch string
		expectedTail  string
		str           string
		sep           string
	}{
		{expectedHead: "he", expectedMatch: "l", expectedTail: "lo", str: "hello", sep: "l"},
		{expectedHead: "hello", expectedMatch: "", expectedTail: "", str: "hello", sep: "x"},
	}

	for _, c := range cases {
		head, match, tail := Partition(c.str, c.sep)
		assert.Equal(t, c.expectedHead, head)
		assert.Equal(t, c.expectedMatch, match)
		assert.Equal(t, c.expectedTail, tail)
	}
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "desserts", Reverse("stressed"))
}

func TestRightJustify(t *testing.T) {
	cases := []struct {
		expect string
		str    string
		length int
		pad    string
	}{
		{expect: "hello", str: "hello", length: 4, pad: " "},
		{expect: "     hello", str: "hello", length: 10, pad: " "},
		{expect: "12312hello", str: "hello", length: 10, pad: "123"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, RightJustify(c.str, c.length, c.pad))
	}
}

func TestSwapCase(t *testing.T) {
	assert.Equal(t, "Ab", SwapCase("aB"))
}
