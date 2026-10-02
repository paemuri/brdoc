package brdoc

import (
	"bytes"
	"regexp"
	"strconv"
)

// "Tax ID" will be used internally to define both CPF and CNPJ.

var (
	cpfRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	cnpjRegexp = regexp.MustCompile(`^[0-9A-Z]{2}\.?[0-9A-Z]{3}\.?[0-9A-Z]{3}/?[0-9A-Z]{4}-?[0-9]{2}$`)
)

type taxIDType int

const (
	taxIDCPF taxIDType = iota
	taxIDCNPJ
)

// IsCPF verifies if the given string is a valid CPF document.
//
// There is no official source for the check digits algorithm or the mask.
// Documents with all digits equal are rejected, as listed by the [DJE layout].
//
// [DJE layout]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=36307
func IsCPF(doc string) bool {
	return isTaxID(doc, taxIDCPF)
}

// IsCNPJ verifies if the given string is a valid CNPJ document, either numeric
// or alphanumeric.
//
// The format and the check digits algorithm are defined by Annex XV of
// [IN RFB 2.119/2022], and the mask by the [Receita Federal FAQ]. Documents
// with order number 0000 are rejected, as listed by the [DJE layout]. It also
// lists base numbers 11.111.111 to 99.999.999 as invalid, but they are
// accepted, as some were issued (e.g. 66.666.666/0001-91).
//
// [IN RFB 2.119/2022]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=76204
// [Receita Federal FAQ]: https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/perguntas-e-respostas/cnpj/cnpj-alfanumerico.pdf
// [DJE layout]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=36307
func IsCNPJ(doc string) bool {
	return isTaxID(doc, taxIDCNPJ)
}

// isTaxID generates the digits for a given CPF or CNPJ and compares it with the
// original digits.
func isTaxID(doc string, idType taxIDType) bool {
	var (
		pattern  *regexp.Regexp
		size     int
		position int
	)

	switch idType {
	case taxIDCPF:
		pattern, size, position = cpfRegexp, 9, 10
	case taxIDCNPJ:
		pattern, size, position = cnpjRegexp, 12, 5
	default:
		return false
	}

	if !pattern.MatchString(doc) {
		return false
	}

	cleanTaxID(&doc)

	// Invalid documents, as documented by `IsCPF` and `IsCNPJ`.
	switch idType {
	case taxIDCPF:
		if allEq(doc) {
			return false
		}
	case taxIDCNPJ:
		if doc[8:12] == "0000" {
			return false
		}
	}

	d := doc[:size]
	digit := calcTaxIDDigit(d, position)

	d = d + digit
	digit = calcTaxIDDigit(d, position+1)

	return doc == d+digit
}

// calcTaxIDDigit calculates the next digit for the given document.
func calcTaxIDDigit(doc string, position int) string {
	var sum int
	for _, r := range doc {
		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

// cleanTaxID removes every rune that is not a digit or an uppercase letter.
func cleanTaxID(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if isDigit(r) || ('A' <= r && r <= 'Z') {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}
