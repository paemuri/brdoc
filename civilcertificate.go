package brdoc

import (
	"regexp"
)

var (
	civilCertificateRegexp = regexp.MustCompile(
		`^\d{6}[. ]?\d{2}[. ]?55[. ]?\d{4}[. ]?[1-7][. ]?\d{5}[. ]?\d{3}[. ]?\d{7}[- ]?\d{2}$`,
	)
)

// IsCivilCertificate verifies if the given string is a valid number of a civil
// (birth, marriage or death) certificate.
//
// The format is defined by Article 473 of [Provimento CNJ 149/2023], with the
// code 55 of the civil registry and the type of book from 1 to 7. The check
// digits algorithm is not published by it, and was verified with real
// certificates.
//
// [Provimento CNJ 149/2023]: https://atos.cnj.jus.br/atos/detalhar/5243
func IsCivilCertificate(doc string) bool {
	if !civilCertificateRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	// Not official logic: check digits calculation.
	first := calcCivilCertificateDigit(doc[:30])
	second := calcCivilCertificateDigit(doc[:30] + string(rune('0'+first)))

	return toInt(rune(doc[30])) == first && toInt(rune(doc[31])) == second
}

// calcCivilCertificateDigit returns the remainder of the weighted sum of `doc`
// divided by 11, or 1 if it is 10. The weights start at 32 minus the length of
// `doc`, and increase by 1, restarting at 0 after 10.
func calcCivilCertificateDigit(doc string) int {
	sum := 0
	weight := 32 - len(doc)
	for _, r := range doc {
		sum += toInt(r) * weight

		weight++
		if weight > 10 {
			weight = 0
		}
	}

	digit := sum % 11
	if digit == 10 {
		return 1
	}

	return digit
}
