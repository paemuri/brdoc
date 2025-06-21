package brdoc

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

// "Cadastro" will be used internally to define both CPF and CNPJ.

var (
	cpfRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	cnpjRegexp = regexp.MustCompile(`^[0-9A-Z]{2}\.?[0-9A-Z]{3}\.?[0-9A-Z]{3}/?[0-9A-Z]{4}-?[0-9]{2}$`)
)

// IsCPF verifies if the given string is a valid CPF document.
func IsCPF(doc string) bool {
	const (
		size = 9
		pos  = 10
	)

	return isCadastro(doc, cpfRegexp, size, pos)
}

// IsCNPJ verifies if the given string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	const (
		size = 12
		pos  = 5
	)

	return isCadastro(doc, cnpjRegexp, size, pos)
}

// isCadastro generates the digits for a given CPF or CNPJ and compares it with
// the original digits.
func isCadastro(
	doc string,
	pattern *regexp.Regexp,
	size int,
	position int,
) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanCadastro(&doc)

	// Invalidates documents with all digits equal.
	if allEq(doc) {
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
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}
