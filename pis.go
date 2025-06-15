package brdoc

import (
	"regexp"
)

var (
	PISRegexp = regexp.MustCompile(`^\d{3}\.?\d{3,5}\.?\d{2,4}-?\d$`)
)

// IsPIS verifies if the given string is a valid PIS number.
func IsPIS(doc string) bool {
	if !PISRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if len(doc) != 11 {
		return false
	}

	return toInt(rune(doc[len(doc)-1])) == calcPISDigit(doc)
}

func calcPISDigit(doc string) int {
	var (
		weights = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	)

	var sum int
	for i, weight := range weights {
		sum += toInt(rune(doc[i])) * weight
	}

	mod := sum % 11
	digit := 0
	if mod > 1 {
		digit = 11 - mod
	}

	return digit
}
