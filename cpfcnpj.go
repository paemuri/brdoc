package brdoc

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

// IsCPF verifies if the string is a valid CPF document.
func IsCPF(doc string) bool {
	return isCPFOrCNPJ(
		doc,
		ValidateCPFFormat,
		9,
		10, 11,
	)
}

// IsCNPJ verifies if the string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	return isCPFOrCNPJ(
		doc,
		ValidateCNPJFormat,
		12,
		5, 6,
	)
}

func isCPFOrCNPJ(doc string, validate func(string) bool, size int, pos1, pos2 int) bool {

	if !validate(doc) {
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
	digit := calculateDigit(d, pos1)

	// Calculates the second digit.
	d = d + digit
	digit = calculateDigit(d, pos2)

	return doc == d+digit
}

// ValidateCPFFormat verifies if the CPF has a valid format.
func ValidateCPFFormat(doc string) bool {

	const (
		pattern = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
	)

	return regexp.MustCompile(pattern).MatchString(doc)
}

// ValidateCNPJFormat verifies if the CNPJ has a valid format.
func ValidateCNPJFormat(doc string) bool {

	const (
		pattern = `^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`
	)

	return regexp.MustCompile(pattern).MatchString(doc)
}

func clean(doc *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

func allEq(doc string) bool {

	if len(doc) == 0 {
		return true
	}

	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

func calculateDigit(doc string, positions int) string {
	sum := 0

	// Sums all the digits in the document.
	// Ex.
	//   3    4    2    6    1    8    7    1    0
	// x10   x9   x8   x7   x6   x5   x4   x3   x2
	//  30 + 36 + 16 + 42 +  6 + 40 + 28 +  3 +  0 = 201
	for i := 0; i < len(doc); i++ {
		digit, _ := strconv.Atoi(string(doc[i]))

		sum += int(digit) * positions
		positions--

		if positions < 2 {
			positions = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}
	return strconv.Itoa(11 - sum)
}
