package brdoc

import (
	"regexp"
)

var (
	cnoRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{5}/?\d{2}$`)
)

// IsCNO verifies if the given string is a valid CNO document, the number of a
// civil construction work. It also works for the CEI numbers of works, which
// were kept by the CNO.
//
// The number of 12 digits is defined by the [CNO layout] of the Receita
// Federal. There is no official source for the check digit algorithm or the
// mask, but the algorithm was verified with all the numbers in the open data
// of the CNO.
//
// [CNO layout]: https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/cno/arquivos/cno-cadastro-nacional-de-obras-layout.pdf
func IsCNO(doc string) bool {
	if !cnoRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	// Not official logic: reject documents with all digits equal.
	if allEq(doc) {
		return false
	}

	return toInt(rune(doc[11])) == calcCNODigit(doc[:11])
}

// calcCNODigit returns 10 minus the rightmost digit of the sum of the last 2
// digits of the weighted sum of `doc`, or 0 if it is 10.
func calcCNODigit(doc string) int {
	sum := 0
	for i, weight := range []int{7, 4, 1, 8, 5, 2, 1, 6, 3, 7, 4} {
		sum += toInt(rune(doc[i])) * weight
	}

	return (10 - (sum/10%10+sum%10)%10) % 10
}
