package main

import (
	"strings"
)

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

func ExpandTabs(str string, tabSize int) string {
	return strings.ReplaceAll(str, "\t", strings.Repeat(" ", tabSize))
}
