package brdoc

import (
	"github.com/Nhanderu/tuyo/convert"
	"regexp"
)

// IsCPF verifies if the string is a valid CPF document.
func IsCPF(doc string) bool {
	doc = clean(doc)
	if !validateCPFFormat(doc) {
		return false
	}

	// Calculates the first digit.
	d := doc[:9]
	digit := calculateDigit(d, 10)

	// Calculates the second digit.
	d = d + digit
	digit = calculateDigit(d, 11)

	return doc == d+digit
}

// IsCNPJ verifies if the string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	doc = clean(doc)
	if !validateCNPJFormat(doc) {
		return false
	}

	// Calculates the first digit.
	d := doc[:12]
	digit := calculateDigit(d, 5)

	// Calculates the second digit.
	d = d + digit
	digit = calculateDigit(d, 6)

	return doc == d+digit
}

func validateCPFFormat(doc string) bool {
	return validateFormat("^\\d{3}\\.?\\d{3}\\.?\\d{3}\\-?\\d{2}$", doc)
}

func validateCNPJFormat(doc string) bool {
	return validateFormat("^\\d{2}\\.?\\d{3}\\.?\\d{3}\\/?\\d{4}\\-?\\d{2}$", doc)
}

func validateFormat(doc, pattern string) bool {
	matched, err := regexp.MatchString(pattern, doc)
	if err != nil {
		return false
	}
	return matched
}

func clean(doc string) string {
	re, err := regexp.Compile("\\D")
	if err != nil {
		return ""
	}
	return re.ReplaceAllString(doc, "")
}

func calculateDigit(doc string, positions int) string {
	sum := 0

	// Sums all the digits in the document.
	// Ex.
	//   3    4    2    6    1    8    7    1    0
	// x10   x9   x8   x7   x6   x5   x4   x3   x2
	//  30 + 36 + 16 + 42 +  6 + 40 + 28 +  3 +  0 = 201
	for i := 0; i < len(doc); i++ {
		digit, _ := convert.ToInt(doc[i])

		sum += digit * positions
		positions--

		if positions < 2 {
			positions = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	} else {
		return convert.ToString(11 - sum)
	}
}
