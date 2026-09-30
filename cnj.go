package brdoc

import (
	"regexp"
)

var (
	cnjRegexp = regexp.MustCompile(`^\d{7}-?\d{2}\.?\d{4}\.?\d\.?\d{2}\.?\d{4}$`)
)

// IsCNJ verifies if the given string is a valid CNJ document, the unique number
// of the processes of the Judiciary.
//
// The format and the check digits algorithm are defined by Article 1 and Annex
// VIII of [Res. CNJ 65/2008].
//
// [Res. CNJ 65/2008]: https://atos.cnj.jus.br/atos/detalhar/119
func IsCNJ(doc string) bool {
	if !cnjRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	// The segment of the Judiciary (J) goes from 1 to 9.
	if doc[13] == '0' {
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
