package brdoc

import (
	"regexp"
)

var (
	iePRRegexp        = regexp.MustCompile(`^\d{3}\.?\d{5}-?\d{2}$`)
	iePRFirstWeights  = []int{3, 2, 7, 6, 5, 4, 3, 2}
	iePRSecondWeights = []int{4, 3, 2, 7, 6, 5, 4, 3, 2}
)

// isIEPR verifies if `doc` is a valid IE of PR (000.00000-00).
func isIEPR(doc string) bool {
	// The format is described by [1], and the rules are published by [2].
	// [1]: https://www.fazenda.pr.gov.br/servicos/Empresa/Cadastro-de-Contribuintes-do-ICMS/Saber-como-se-calcula-o-digito-verificador-da-inscricao-estadual-kZrX1Bol.
	// [2]: http://www.sintegra.gov.br/Cad_Estados/cad_PR.html.

	if !iePRRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	return toInt(rune(doc[8])) == calcIEPRDigit(doc, iePRFirstWeights) &&
		toInt(rune(doc[9])) == calcIEPRDigit(doc, iePRSecondWeights)
}

// calcIEPRDigit returns 11 minus the remainder of the weighted sum of `doc`
// divided by 11, or 0 if it is 10 or 11.
func calcIEPRDigit(doc string, weights []int) int {
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
