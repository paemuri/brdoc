package brdoc

import (
	"bytes"
)

// allDigit checks if every rune in a given string is a digit.
func allDigit(doc string) bool {
	for _, r := range doc {
		if !isDigit(r) {
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
		if isDigit(r) {
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

// isDigit is a simpler version of unicode.IsDigit: verifies whether a rune is
// a single digit number.
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
