package brdoc

// Official name for RENAVAM seems to be all upper case, not "Renavam" or
// something similar.

// IsRENAVAM verifies if the given string is a valid RENAVAM document.
//
// The number format and the check digit method are defined by Article 1 of
// [Portaria DENATRAN 27/2013].
//
// [Portaria DENATRAN 27/2013]: https://www.gov.br/transportes/pt-br/assuntos/transito/arquivos-senatran/portarias/2013/portaria0272013.pdf
func IsRENAVAM(doc string) bool {
	if len(doc) != 11 {
		return false
	}
	if !allDigit(doc) {
		return false
	}

	return toInt(rune(doc[len(doc)-1])) == calcRENAVAMDigit(doc)
}

func calcRENAVAMDigit(doc string) int {
	// Not official logic: exact weights and remainder 10 rule.
	var (
		weights = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	)

	var sum int
	for i, weight := range weights {
		sum += toInt(rune(doc[i])) * weight
	}

	digit := (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}

	return digit
}
