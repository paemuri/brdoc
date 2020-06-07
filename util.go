package brdoc

import (
	"unicode"
)

// allDigit checks if every rune in a given string is a digit.
func allDigit(doc string) bool {

	for _, r := range doc {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// toInt converts a rune to an int.
func toInt(r rune) int {
	return int(r - '0')
}
