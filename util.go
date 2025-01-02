package brdoc

import (
	"bytes"
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

// cleanNonDigits removes every rune that is not a digit.
func cleanNonDigits(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}
