package brdoc

import (
	"regexp"
)

var (
	cnjRegexp = regexp.MustCompile(`^\d{7}-?\d{2}\.?\d{4}\.?\d\.?\d{2}\.?\d{4}$`)

	cnjCourts = map[byte][][2]int{
		'1': {{0, 0}},                       // STF.
		'2': {{0, 0}},                       // CNJ.
		'3': {{0, 0}},                       // STJ.
		'4': {{1, 6}, {90, 90}},             // TRFs and CJF.
		'5': {{0, 24}, {90, 90}},            // TST, TRTs and CSJT.
		'6': {{0, 27}},                      // TSE and TREs.
		'7': {{0, 12}},                      // STM and CJMs.
		'8': {{1, 27}},                      // TJs.
		'9': {{13, 13}, {21, 21}, {26, 26}}, // TJMs of MG, RS and SP.
	}
)

// IsCNJ verifies if the given string is a valid CNJ document, the unique number
// of the processes of the Judiciary.
//
// The format, the check digits algorithm and the courts (TR) of each segment of
// the Judiciary (J) are defined by Article 1 and Annex VIII of
// [Res. CNJ 65/2008].
//
// [Res. CNJ 65/2008]: https://atos.cnj.jus.br/atos/detalhar/119
func IsCNJ(doc string) bool {
	if !cnjRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if !validCNJCourt(doc[13], toInt(rune(doc[14]))*10+toInt(rune(doc[15]))) {
		return false
	}

	// The check digits (DD) are moved to the end, and the number modulo 97 must
	// be 1. The remainder is calculated digit by digit, as the number is too
	// large for an int.
	rem := 0
	for _, r := range doc[:7] + doc[9:] + doc[7:9] {
		rem = (rem*10 + toInt(r)) % 97
	}

	return rem == 1
}

// validCNJCourt verifies if the court (TR) exists in the segment of the
// Judiciary (J).
func validCNJCourt(segment byte, court int) bool {
	for _, r := range cnjCourts[segment] {
		if r[0] <= court && court <= r[1] {
			return true
		}
	}

	return false
}
