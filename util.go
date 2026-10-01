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

// isDigit is a simpler version of unicode.IsDigit: verifies whether a rune is a
// single digit number.
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// calcMod11Digit returns 11 minus the remainder of `calcMod11Sum` of `doc`
// divided by 11, or 0 if it is 10 or 11.
func calcMod11Digit(doc string) int {
	digit := 11 - calcMod11Sum(doc)%11
	if digit >= 10 {
		return 0
	}

	return digit
}

// calcMod11Sum returns the sum of the digits of `doc` multiplied by the weights
// from 2 to 9, from right to left, restarting at 2 after 9. Letters are valued
// as their ASCII code minus 48.
func calcMod11Sum(doc string) int {
	sum := 0
	weight := 2
	for i := len(doc) - 1; i >= 0; i-- {
		sum += toInt(rune(doc[i])) * weight

		weight++
		if weight > 9 {
			weight = 2
		}
	}

	return sum
}

// calcMod10Digit returns 10 minus the remainder of the sum of the digits of the
// products of `doc` divided by 10, or 0 if it is 10. The weights alternate
// between 2 and 1, from right to left.
func calcMod10Digit(doc string) int {
	sum := 0
	weight := 2
	for i := len(doc) - 1; i >= 0; i-- {
		product := toInt(rune(doc[i])) * weight
		sum += product/10 + product%10

		weight = 3 - weight
	}

	return (10 - sum%10) % 10
}
