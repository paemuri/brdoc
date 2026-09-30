package brdoc

// IsIE verifies if `doc` is a valid IE (Inscrição Estadual) of the given `uf`.
func IsIE(doc string, uf UF) bool {
	// The rules of each UF are published by [1].
	// [1]: http://www.sintegra.gov.br/insc_est.html.

	switch uf {
	case AC:
		return isIEAC(doc)
	case PR:
		return isIEPR(doc)
	case SP:
		return isIESP(doc)
	default:
		return false
	}
}

// calcIEDigit returns 11 minus the remainder of the weighted sum of `doc`
// divided by 11, or 0 if it is 10 or 11, a calculation shared by several UFs.
func calcIEDigit(doc string, weights []int) int {
	sum := 0
	for i, weight := range weights {
		sum += toInt(rune(doc[i])) * weight
	}

	digit := 11 - sum%11
	if digit >= 10 {
		return 0
	}

	return digit
}
