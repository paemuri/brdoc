package brdoc

import (
	"regexp"
)

var (
	ieSPRegexp        = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}\.?\d{3}$`)
	ieSPRuralRegexp   = regexp.MustCompile(`^P-?\d{8}\.?\d/?\d{3}$`)
	ieSPFirstWeights  = []int{1, 3, 4, 5, 6, 7, 8, 10}
	ieSPSecondWeights = []int{3, 2, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

// isIESP verifies if `doc` is a valid IE of SP, either of industrials and
// merchants (000.000.000.000) or of rural producers (P-00000000.0/000).
func isIESP(doc string) bool {
	// The rules are published by [1].
	// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_SP.html.

	if ieSPRuralRegexp.MatchString(doc) {
		cleanNonDigits(&doc)
		return toInt(rune(doc[8])) == calcIESPDigit(doc, ieSPFirstWeights)
	}

	if ieSPRegexp.MatchString(doc) {
		cleanNonDigits(&doc)
		return toInt(rune(doc[8])) == calcIESPDigit(doc, ieSPFirstWeights) &&
			toInt(rune(doc[11])) == calcIESPDigit(doc, ieSPSecondWeights)
	}

	return false
}

// calcIESPDigit returns the rightmost digit of the remainder of the weighted
// sum of `doc` divided by 11.
func calcIESPDigit(doc string, weights []int) int {
	sum := 0
	for i, weight := range weights {
		sum += toInt(rune(doc[i])) * weight
	}
	return sum % 11 % 10
}
