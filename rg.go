package brdoc

import (
	"bytes"
	"errors"
	"regexp"
	"unicode"
)

var (
	rgRegexp = map[FederativeUnit]*regexp.Regexp{
		SP: regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}-?[0-9xX]$`),
		RJ: regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}-?[0-9xX]$`),
	}
	errNotImplementedUF = errors.New("federative unit not implemented")
)

// IsRG verifies if `doc` is a valid RG.
// `uf` represents the Federative Unit the RG belongs to.
// Currently, the only implemented UFs are the following (the remaining will return an error):
//   - SP
//   - RJ
func IsRG(doc string, uf FederativeUnit) (isValid bool, err error) {
	if uf != SP && uf != RJ {
		return false, errNotImplementedUF
	}

	if !rgRegexp[uf].MatchString(doc) {
		return false, nil
	}

	checkDigit := getRGCheckDigit(doc)
	calculatedCheckDigit := calculateRGCheckDigit(doc)

	return calculatedCheckDigit == checkDigit, nil

}

func calculateRGCheckDigit(doc string) int {
	doc = cleanNonRGDigits(doc)
	digitsSum := 0
	for i, digit := range doc[:len(doc)-1] {
		a := toInt(digit) * (i + 2)
		digitsSum += a
	}

	return 11 - (digitsSum % 11)
}

func getRGCheckDigit(doc string) int {
	switch lastDigit := rune(doc[len(doc)-1]); lastDigit {
	case 'X', 'x':
		return 10
	case '0':
		return 11
	default:
		return toInt(lastDigit)
	}
}

// cleanNonDigits removes every rune that is not a valid RG digit.
func cleanNonRGDigits(doc string) string {

	buf := bytes.NewBufferString("")
	for _, r := range doc {
		if isRGValidDigit(r) {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

func isRGValidDigit(r rune) bool {
	return unicode.IsDigit(r) || r == 'X' || r == 'x'
}
