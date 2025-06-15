package brdoc

import (
	"regexp"
	"strconv"
)

// "Cadastro" will be used internally to define both CPF and CNPJ.

var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

// IsCPF verifies if the given string is a valid CPF document.
func IsCPF(doc string) bool {
	const (
		size = 9
		pos  = 10
	)

	return isCadastro(doc, CPFRegexp, size, pos)
}

// IsCNPJ verifies if the given string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	const (
		size = 12
		pos  = 5
	)

	return isCadastro(doc, CNPJRegexp, size, pos)
}

// isCadastro generates the digits for a given CPF or CNPJ and compares it with the original digits.
func isCadastro(doc string, pattern *regexp.Regexp, size int, position int) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

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
