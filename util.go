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

// allEq checks if every rune in a given string is equal.
func allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// isFrom checks whether the doc's UF is part of the given options.
func isFrom(uf UF, ufs []UF) bool {
	if len(ufs) == 0 {
		return true
	}
	for _, u := range ufs {
		if uf == u {
			return true
		}
	}
	return false
}

// cleanNonDigitsAndNonLetters removes every rune that is not a digit and letters.
func cleanNonDigitsAndNonLetters(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}
