package tstring

import (
	"strings"
	"unicode"
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

// LastPartition partition by separator.
func LastPartition(str string, sep string) (head string, match string, tail string) {
	index := strings.LastIndex(str, sep)

	if index == -1 {
		return str, "", ""
	}

	return str[0:index], sep, str[index+1:]
}

// LeftJustify justify left side.
func LeftJustify(str string, length int, pad string) string {
	if len(str) >= length {
		return str
	}

	return str + strings.Repeat(pad, length)[:length-len(str)]
}

// Partition partition by separator.
func Partition(str string, sep string) (head string, match string, tail string) {
	index := strings.Index(str, sep)

	if index == -1 {
		return str, "", ""
	}

	return str[0:index], sep, str[index+1:]
}

// Reverse reverse string.
func Reverse(str string) string {
	reversedString := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversedString += string(str[i])
	}

	return reversedString
}

// RightJustify justify left side.
func RightJustify(str string, length int, pad string) string {
	if len(str) >= length {
		return str
	}

	return strings.Repeat(pad, length)[:length-len(str)] + str
}

// SwapCase convert from lower to upper or upper or lower.
func SwapCase(str string) string {
	runeStr := []rune(str)
	returnedStr := ""

	for i := 0; i < len(runeStr); i++ {
		if unicode.IsUpper(runeStr[i]) {
			returnedStr += strings.ToLower(string(runeStr[i]))
		} else if unicode.IsLower(runeStr[i]) {
			returnedStr += strings.ToUpper(string(runeStr[i]))
		} else {
			returnedStr += string(runeStr[i])
		}
	}

	return returnedStr
}
