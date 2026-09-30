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

	return toInt(rune(doc[8])) == calcIEDigit(doc, iePRFirstWeights) &&
		toInt(rune(doc[9])) == calcIEDigit(doc, iePRSecondWeights)
}
