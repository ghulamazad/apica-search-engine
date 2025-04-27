package utils

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	text = strings.ToLower(text)
	text = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return r
		}
		return ' ' // Replace bad characters with space
	}, text)

	return strings.Fields(text)
}
