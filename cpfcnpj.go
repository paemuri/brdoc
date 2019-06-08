package brdoc

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

// Regexp pattern for CPF and CNPJ.
var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

// IsCPF verifies if the string is a valid CPF
// document.
func IsCPF(doc string) bool {

	const (
		sizeWithoutDigits = 9
		position          = 10
	)

	return isCPFOrCNPJ(doc, CPFRegexp, sizeWithoutDigits, position)
}

// IsCNPJ verifies if the string is a valid CNPJ
// document.
func IsCNPJ(doc string) bool {

	const (
		sizeWithoutDigits = 12
		position          = 5
	)

	return isCPFOrCNPJ(doc, CNPJRegexp, sizeWithoutDigits, position)
}

// isCPFOrCNPJ generates the digits for a given
// CPF or CNPJ and compares it with the original
// digits.
func isCPFOrCNPJ(doc string, pattern *regexp.Regexp, size int, position int) bool {

	if !pattern.MatchString(doc) {
		return false
	}

	// Removes special characters.
	clean(&doc)

	// Invalidates documents with all
	// digits equal.
	if allEq(doc) {
		return false
	}

	// Calculates the first digit.
	d := doc[:size]
	digit := calculateDigit(d, position)

	// Calculates the second digit.
	d = d + digit
	digit = calculateDigit(d, position+1)

	return doc == d+digit
}

// clean removes every rune that is not a digit.
func clean(doc *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

// allEq checks if every rune in a given string
// is equal.
func allEq(doc string) bool {

	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// calculateDigit calculates the next digit for
// the given document.
func calculateDigit(doc string, position int) string {

	var sum int
	for _, r := range doc {

		sum += int(r-'0') * position
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
