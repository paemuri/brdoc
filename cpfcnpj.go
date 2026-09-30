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

// IsCPF verifies if the given string is a valid CPF document.
func IsCPF(doc string) bool {
	// No official source for the check digits algorithm or the mask.

	const (
		size = 9
		pos  = 10
	)

	return isCadastro(doc, cpfRegexp, size, pos, isInvalidCPFBase)
}

// IsCNPJ verifies if the given string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	// The format, including alphanumeric characters, and the check digits
	// algorithm are defined by Annex XV of [1]. The mask is defined by [2].
	// [1]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=76204.
	// [2]: https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/perguntas-e-respostas/cnpj/cnpj-alfanumerico.pdf.

	const (
		size = 12
		pos  = 5
	)

	return isCadastro(doc, cnpjRegexp, size, pos, isInvalidCNPJBase)
}

// isInvalidCPFBase rejects CPF documents with all digits equal, as listed by
// [1].
// [1]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=36307.
func isInvalidCPFBase(doc string) bool {
	return allEq(doc)
}

// isInvalidCNPJBase rejects CNPJ documents with base number 11.111.111 to
// 99.999.999, or with order number 0000, as listed by [1].
// [1]: http://normas.receita.fazenda.gov.br/sijut2consulta/anexoOutros.action?idArquivoBinario=36307.
func isInvalidCNPJBase(doc string) bool {
	base := doc[:8]
	order := doc[8:12]

	return (base[0] != '0' && isDigit(rune(base[0])) && allEq(base)) ||
		order == "0000"
}

// isCadastro generates the digits for a given CPF or CNPJ and compares it with
// the original digits.
func isCadastro(
	doc string,
	pattern *regexp.Regexp,
	size int,
	position int,
	isInvalidBase func(doc string) bool,
) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanCadastro(&doc)

	if isInvalidBase(doc) {
		return false
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
