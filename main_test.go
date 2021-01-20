package main

import (
	"testing"
)

func TestCenter1(t *testing.T) {
	expect := "hello"
	actual := Center("hello", 4, " ")

	if expect != actual {
		t.Errorf("\"%s \"!=\" %s\"", expect, actual)
	}
}

func TestCenter2(t *testing.T) {
	expect := "  hello   "
	actual := Center("hello", 10, " ")

	if expect != actual {
		t.Errorf("\"%s \"!=\" %s\"", expect, actual)
	}
}
