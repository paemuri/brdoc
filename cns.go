package brdoc

import (
	"regexp"
	"strconv"
)

var (
	cnsRegexp = regexp.MustCompile(
		`^([12]\d{2}(\s?\d{4}){2}\s?00[01]\d|[789]\d{2}(\s?\d{4}){3})$`,
	)
)

// IsCNS verifies if the given string is a valid CNS document.
func IsCNS(doc string) bool {
	// The validation follows the Ministry of Health's algorithm, published at
	// [1].
	// [1]: https://rni-docs.anvisa.gov.br/docs/regras_gerais/validacoes/validacaoCNS/.

	if !cnsRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	// CNS starting with 1 or 2 is generated from a PIS, so its last 4 digits
	// must be exactly the ones generated from the first 11.
	if doc[0] == '1' || doc[0] == '2' {
		return doc == genCNSFromPIS(doc[:11])
	}

	return calcCNSSum(doc)%11 == 0
}

func genCNSFromPIS(pis string) string {
	sum := calcCNSSum(pis)

	digit := 11 - (sum % 11)
	if digit == 11 {
		digit = 0
	}
	if digit == 10 {
		sum += 2
		digit = 11 - (sum % 11)
		return pis + "001" + strconv.Itoa(digit)
	}

	return pis + "000" + strconv.Itoa(digit)
}

func calcCNSSum(doc string) int {
	sum := 0
	for i, r := range doc {
		sum += toInt(r) * (15 - i)
	}

	return sum
}
