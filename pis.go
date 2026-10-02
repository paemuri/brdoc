package brdoc

import (
	"regexp"
)

var (
	pisRegexp = regexp.MustCompile(`^\d{3}\.?\d{5}\.?\d{2}-?\d$`)
)

// IsPIS verifies if the given string is a valid PIS number. It also works for
// NIS and NIT numbers.
//
// There is no official source for any of this logic. [Caixa] states that NIS
// and PIS numbers are the same, but there is no official source for NIT
// numbers.
//
// [Caixa]: https://www.caixa.gov.br/servicos/nis/Paginas/default.aspx
func IsPIS(doc string) bool {
	if !pisRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if allEq(doc) {
		return false
	}

	return toInt(rune(doc[len(doc)-1])) == calcPISDigit(doc)
}

// calcPISDigit returns 11 minus the remainder of the weighted sum of the first
// 10 digits of `doc` divided by 11, or 0 if the remainder is 0 or 1.
func calcPISDigit(doc string) int {
	var (
		weights = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	)

	var sum int
	for i, weight := range weights {
		sum += toInt(rune(doc[i])) * weight
	}

	mod := sum % 11
	digit := 0
	if mod > 1 {
		digit = 11 - mod
	}

	return digit
}
