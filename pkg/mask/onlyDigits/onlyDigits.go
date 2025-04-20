package onlydigits

import (
	"strings"
	"unicode"
)

func SetOnlyDigits(text string) string {
	numericText := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, text)

	return numericText
}
