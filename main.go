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

		return strings.Repeat(" ", leftPadding) + str + strings.Repeat(" ", rightPadding)
	}
}
