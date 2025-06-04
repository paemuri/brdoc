package brdoc

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Regex patterns for CPF and CNPJ formats.
var (
	CPFRegexp       = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp      = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
	CNPJAlfhaRegexp = regexp.MustCompile(`^(?:[A-Za-z0-9]{2}\.[A-Za-z0-9]{3}\.[A-Za-z0-9]{3}/[A-Za-z0-9]{4}-\d{2}|[A-Za-z0-9]{12}\d{2})$`)
)

// IsCPF checks whether a string is a valid CPF (Cadastro de Pessoas Físicas).
func IsCPF(doc string) bool {

	return validateCPF(doc)
}

// IsCNPJ checks whether a string is a valid CNPJ (Cadastro Nacional da Pessoa Jurídica).
func IsCNPJ(doc string) bool {

	return validateCNPJ(doc)
}

// validateCPF performs full CPF validation including digit verification.
func validateCPF(cpf string) bool {

	if !CPFRegexp.MatchString(cpf) {
		return false
	}

	cleanNonDigits(&cpf)

	// Invalidates documents with all digits equal.
	if allEq(cpf) {
		return false
	}

	d := cpf[:9]
	digit := calculateDigit(d, 10)

	d = d + digit
	digit = calculateDigit(d, 10+1)

	return cpf == d+digit
}

// validateCNPJ performs full CNPJ validation including digit verification.
func validateCNPJ(cnpj string) bool {

	if !CNPJRegexp.MatchString(cnpj) && !CNPJAlfhaRegexp.MatchString(cnpj) {
		return false
	}

	// Remove doc format
	replacer := strings.NewReplacer(".", "", "/", "", "-", "")
	cnpj = replacer.Replace(cnpj)

	if len(cnpj) != 14 || allEq(cnpj[:12]) {
		return false
	}

	base := cnpj[:12]
	dvInf1 := int(cnpj[12] - '0')
	dvInf2 := int(cnpj[13] - '0')

	dvCalc1, ok1 := calculateCNPJDigit(base)
	if !ok1 || dvCalc1 != dvInf1 {
		return false
	}

	dvCalc2, ok2 := calculateCNPJDigit(base + strconv.Itoa(dvCalc1))
	if !ok2 || dvCalc2 != dvInf2 {
		return false
	}

	return true
}

// allEq checks if every rune in a given string is equal.
func allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// calculateDigit computes a check digit for CPF using weighted sum logic.
func calculateDigit(doc string, position int) string {
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

// calculateCNPJDigit computes a CNPJ verification digit using weights.
func calculateCNPJDigit(cnpjBase string) (int, bool) {
	weights := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	n := len(cnpjBase)

	for i := n - 1; i >= 0; i-- {
		r := rune(cnpjBase[i])
		val := valueForCheckDigit(r)
		if val < 0 {
			return 0, false
		}
		weight := weights[(n-1-i)%len(weights)]
		sum += val * weight
	}

	remainder := sum % 11
	if remainder == 0 || remainder == 1 {
		return 0, true
	}
	return 11 - remainder, true
}

// valueForCheckDigit returns the numeric value for a digit or alphanumeric character.
func valueForCheckDigit(r rune) int {
	ascii := int(r)
	if unicode.IsDigit(r) || (r >= 'A' && r <= 'Z') {
		return ascii - 48
	}
	return -1
}
