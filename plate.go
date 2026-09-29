package brdoc

import (
	"regexp"
)

var (
	plateNationalRegexp = regexp.MustCompile(`^[A-Z]{3}-?\d{4}$`)
	plateMercosulRegexp = regexp.MustCompile(`^[A-Z]{3}\d[A-Z]\d{2}$`)
)

// IsPlate verifies if the given string is a valid license plate.
// It can be either in the old national format or the new Mercosul one.
func IsPlate(doc string) bool {
	return IsNationalPlate(doc) || IsMercosulPlate(doc)
}

// IsNationalPlate verifies if the given string is a valid license plate in the
// old national format.
func IsNationalPlate(doc string) bool {
	// The format is defined by Article 2 of [1].
	// [1]: https://www.gov.br/transportes/pt-br/assuntos/transito/conteudo-contran/resolucoes/resolucao9692022.pdf.

	return plateNationalRegexp.MatchString(doc)
}

// IsMercosulPlate verifies if the given string is a valid license plate in the
// new Mercosul format.
func IsMercosulPlate(doc string) bool {
	// The format is defined by item 1.2 of Annex I of [1].
	// [1]: https://www.gov.br/transportes/pt-br/assuntos/transito/conteudo-contran/resolucoes/resolucao9692022anexos.pdf.

	return plateMercosulRegexp.MatchString(doc)
}
