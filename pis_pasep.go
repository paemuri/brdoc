package brdoc

import (
	"regexp"
)

const pisLength = 11

// Regexp pattern for PIS.
var (
	PISRegexp = regexp.MustCompile(`^\d{3}\.?\d{3,5}\.?\d{2,4}-?\d$`)

	pisWeights = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsPIS verifies if the given string is a valid PIS number.
func IsPIS(doc string) bool {
	if !PISRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if len(doc) != pisLength {
		return false
	}

	var sum int
	for i, weight := range pisWeights {
		sum += toInt(rune(doc[i])) * weight
	}
	mod := sum % pisLength

	digit := 0
	if mod > 1 {
		digit = pisLength - mod
	}

	return toInt(rune(doc[10])) == digit
}
