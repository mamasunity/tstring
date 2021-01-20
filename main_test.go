package main

import "testing"

func TestCenter(t *testing.T) {
	expect := "hello"
	actual := Center("hello", 4, " ")

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
