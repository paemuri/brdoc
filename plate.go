package brdoc

import (
	"regexp"
)

var (
	plateNationalRegexp = regexp.MustCompile(`^[A-Z]{3}-?\d{4}$`)
	plateMercosulRegexp = regexp.MustCompile(`^[A-Z]{3}\d[A-Z]\d{2}$`)
)

// IsPlate verifies if the given string is a valid license plate. It can be
// either in the old national format or the new Mercosul one.
func IsPlate(doc string) bool {
	return IsNationalPlate(doc) || IsMercosulPlate(doc)
}

// IsNationalPlate verifies if the given string is a valid license plate in the
// old national format.
//
// The format is defined by Article 2 of [Res. CONTRAN 969/2022].
//
// [Res. CONTRAN 969/2022]: https://www.gov.br/transportes/pt-br/assuntos/transito/conteudo-contran/resolucoes/resolucao9692022.pdf
func IsNationalPlate(doc string) bool {
	return plateNationalRegexp.MatchString(doc)
}

// IsMercosulPlate verifies if the given string is a valid license plate in the
// new Mercosul format.
//
// The format is defined by item 1.2 of Annex I of [Res. CONTRAN 969/2022].
//
// [Res. CONTRAN 969/2022]: https://www.gov.br/transportes/pt-br/assuntos/transito/conteudo-contran/resolucoes/resolucao9692022anexos.pdf
func IsMercosulPlate(doc string) bool {
	return plateMercosulRegexp.MatchString(doc)
}
