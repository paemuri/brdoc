package brdoc

import (
	"bytes"
	"regexp"
	"unicode"
)

var (
	RGRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}-?[0-9xX]$`)
)

func IsRG(doc string) bool {
	if !RGRegexp.MatchString(doc) {
		return false
	}

	checkDigit := getRGCheckDigit(doc)
	calculatedCheckDigit := calculateRGCheckDigit(doc)

	return calculatedCheckDigit == checkDigit
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
