package brdoc

import (
	"regexp"
)

var (
	ieACRegexp        = regexp.MustCompile(`^01\.?\d{3}\.?\d{3}/?\d{3}-?\d{2}$`)
	ieACFirstWeights  = []int{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	ieACSecondWeights = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// isIEAC verifies if `doc` is a valid IE of AC (01.000.000/000-00).
func isIEAC(doc string) bool {
	// The rules are published by [1].
	// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_AC.html.

	if !ieACRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	return toInt(rune(doc[11])) == calcIEDigit(doc, ieACFirstWeights) &&
		toInt(rune(doc[12])) == calcIEDigit(doc, ieACSecondWeights)
}
