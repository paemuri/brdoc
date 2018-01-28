package brdoc

import (
	"regexp"
	"strconv"
)

// IsCPF verifies if the string is a valid CPF document.
func IsCPF(doc string) bool {
	return isCPFOrCNPJ(
		doc,
		validateCPFFormat,
		9,
		10, 11,
	)
}

// IsCNPJ verifies if the string is a valid CNPJ document.
func IsCNPJ(doc string) bool {
	return isCPFOrCNPJ(
		doc,
		validateCNPJFormat,
		12,
		5, 6,
	)
}

func isCPFOrCNPJ(doc string, validate func(string) bool, size int, pos1, pos2 int) bool {

	if !validate(doc) {
		return false
	}

	doc = clean(doc)

	// Calculates the first digit.
	d := doc[:size]
	digit := calculateDigit(d, pos1)

	// Calculates the second digit.
	d = d + digit
	digit = calculateDigit(d, pos2)

	return doc == d+digit
}

func validateCPFFormat(doc string) bool {
	return validateFormat("^\\d{3}\\.?\\d{3}\\.?\\d{3}-?\\d{2}$", doc,
		"^000\\.?000\\.?000-?00$",
		"^111\\.?111\\.?111-?11$",
		"^222\\.?222\\.?222-?22$",
		"^333\\.?333\\.?333-?33$",
		"^444\\.?444\\.?444-?44$",
		"^555\\.?555\\.?555-?55$",
		"^666\\.?666\\.?666-?66$",
		"^777\\.?777\\.?777-?77$",
		"^888\\.?888\\.?888-?88$",
		"^999\\.?999\\.?999-?99$")
}

func validateCNPJFormat(doc string) bool {
	return validateFormat("^\\d{2}\\.?\\d{3}\\.?\\d{3}\\/?\\d{4}-?\\d{2}$", doc,
		"^\\d{2}\\.?\\d{3}\\.?\\d{3}\\/?0000-?\\d{2}$",
		"^11\\.?111\\.?111\\/?1111-?11$",
		"^22\\.?222\\.?222\\/?2222-?22$",
		"^33\\.?333\\.?333\\/?3333-?33$",
		"^44\\.?444\\.?444\\/?4444-?44$",
		"^55\\.?555\\.?555\\/?5555-?55$",
		"^66\\.?666\\.?666\\/?6666-?66$",
		"^77\\.?777\\.?777\\/?7777-?77$",
		"^88\\.?888\\.?888\\/?8888-?88$",
		"^99\\.?999\\.?999\\/?9999-?99$")
}

func validateFormat(pattern, doc string, invalid ...string) bool {
	for _, p := range invalid {
		if ok, err := regexp.MatchString(p, doc); err != nil || ok {
			return false
		}
	}
	ok, err := regexp.MatchString(pattern, doc)
	return err == nil && ok
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
		digit, _ := strconv.ParseInt(string(doc[i]), 10, 0)

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
	return strconv.FormatInt(int64(11-sum), 10)
}
