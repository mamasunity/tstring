package main

import (
	"strings"
)

// Center make string centerd.
func Center(str string, length int, pad string) string {
	if len(str) >= length {
		return str
	} else {
		restPadding := length - len(str)
		leftPadding := restPadding / 2

		rightPadding := restPadding / 2
		if restPadding%2 == 1 {
			rightPadding++
		}

		return strings.Repeat(pad, leftPadding)[:leftPadding] + str + strings.Repeat(pad, rightPadding)[:rightPadding]
	}
}

// ExpandTabs replace from tab to some space.
func ExpandTabs(str string, tabSize int) string {
	if tabSize <= 0 {
		panic("tabSize should greater than 0 or equal 0")
	}

	return strings.ReplaceAll(str, "\t", strings.Repeat(" ", tabSize))
}

// FirstRuneToLower returns first rune lower string when rune is alphabet.
func FirstRuneToLower(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

// FirstRuneToUpper returns first rune upper string when rune is alphabet.
func FirstRuneToUpper(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

// Insert is insertion on index position.
func Insert(dst string, src string, index int) string {
	return dst[:index] + src + dst[index:]
}

func main() {
	// DO NOT IMPLEMENT
}
