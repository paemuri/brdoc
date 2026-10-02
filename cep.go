package brdoc

import (
	"regexp"
	"strconv"
)

var (
	cepRegexp = regexp.MustCompile(`^\d{5}-?\d{3}$`)
)

// cepRanges maps the ranges of the first 3 digits of the CEP to their UFs.
var cepRanges = []struct {
	first, last int
	uf          UF
}{
	{10, 199, SP},
	{200, 289, RJ},
	{290, 299, ES},
	{300, 399, MG},
	{400, 489, BA},
	{490, 499, SE},
	{500, 569, PE},
	{570, 579, AL},
	{580, 589, PB},
	{590, 599, RN},
	{600, 639, CE},
	{640, 649, PI},
	{650, 659, MA},
	{660, 688, PA},
	{689, 689, AP},
	{690, 692, AM},
	{693, 693, RR},
	{694, 698, AM},
	{699, 699, AC},
	{700, 727, DF},
	{728, 729, GO},
	{730, 736, DF},
	{737, 767, GO},
	{768, 769, RO},
	{770, 779, TO},
	{780, 788, MT},
	{790, 799, MS},
	{800, 879, PR},
	{880, 899, SC},
	{900, 999, RS},
}

// IsCEPFrom verifies if `doc` is a valid CEP. Also, it validates if its related
// UF is part of the given options. If none is provided, it validates the
// document for any state/district. This function is a wrapper around [IsCEP].
func IsCEPFrom(doc string, ufs ...UF) bool {
	valid, uf := IsCEP(doc)
	if !valid {
		return false
	}

	return isFrom(uf, ufs)
}

// IsCEP verifies if `doc` is a valid CEP and returns its related UF.
//
// The CEP format is described by [Correios], and the ranges of each UF are
// listed at [Busca CEP].
//
// [Correios]: https://www.correios.com.br/enviar/precisa-de-ajuda/tudo-sobre-cep
// [Busca CEP]: https://buscacepinter.correios.com.br/app/faixa_cep_uf_localidade/index.php
func IsCEP(doc string) (valid bool, uf UF) {
	if !cepRegexp.MatchString(doc) {
		return false, ""
	}

	prefix, _ := strconv.Atoi(doc[:3])
	for _, r := range cepRanges {
		if r.first <= prefix && prefix <= r.last {
			return true, r.uf
		}
	}

	return false, ""
}
