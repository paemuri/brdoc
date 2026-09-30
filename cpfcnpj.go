package brdoc

import (
	"bytes"
	"regexp"
	"strconv"
)

// "Cadastro" will be used internally to define both CPF and CNPJ.

var (
	cpfRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	cnpjRegexp = regexp.MustCompile(`^[0-9A-Z]{2}\.?[0-9A-Z]{3}\.?[0-9A-Z]{3}/?[0-9A-Z]{4}-?[0-9]{2}$`)
)

type docType int

const (
	docCPF docType = iota
	docCNPJ
)

// IsCPF verifies if the given string is a valid CPF document.
func IsCPF(doc string) bool {
	// No official source for the check digits algorithm or the mask.

	return isCadastro(doc, docCPF)
}

// IsCNPJ verifies if the given string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	// The format, including alphanumeric characters, and the check digits
	// algorithm are defined by Annex XV of [1]. The mask is defined by [2].
	// [1]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=76204.
	// [2]: https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/perguntas-e-respostas/cnpj/cnpj-alfanumerico.pdf.

	return isCadastro(doc, docCNPJ)
}

// isCadastro generates the digits for a given CPF or CNPJ and compares it with
// the original digits.
func isCadastro(doc string, docType docType) bool {
	var (
		pattern  *regexp.Regexp
		size     int
		position int
	)

	switch docType {
	case docCPF:
		pattern, size, position = cpfRegexp, 9, 10
	case docCNPJ:
		pattern, size, position = cnpjRegexp, 12, 5
	default:
		return false
	}

	if !pattern.MatchString(doc) {
		return false
	}

	cleanCadastro(&doc)

	// The invalid documents are listed by [1]. It also lists CNPJ base numbers
	// 11.111.111 to 99.999.999 as invalid, but they are not rejected, as some
	// were issued (e.g. 66.666.666/0001-91).
	// [1]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=36307.
	switch docType {
	case docCPF:
		if allEq(doc) {
			return false
		}
	case docCNPJ:
		if doc[8:12] == "0000" {
			return false
		}
	}

	d := doc[:size]
	digit := calcCadastroDigit(d, position)

	d = d + digit
	digit = calcCadastroDigit(d, position+1)

	return doc == d+digit
}

// calcCadastroDigit calculates the next digit for the given document.
func calcCadastroDigit(doc string, position int) string {
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

func cleanCadastro(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if isDigit(r) || ('A' <= r && r <= 'Z') {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}
