package brdoc

import (
	"bytes"
	"errors"
	"regexp"
	"unicode"
)

var (
	rgRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}-?[0-9xX]$`)
)

// IsRG verifies if `doc` is a valid RG.
// `uf` represents the UFs the RG belongs to. Currently, the only implemented
// UFs are the following:
//   - SP
//   - RJ
//
// All the remaining UFs will return an error.
func IsRG(doc string, uf UF) (isValid bool, err error) {
	if uf != SP && uf != RJ {
		err = errors.New("federative unit not implemented")
		return
	}

	if !rgRegexp.MatchString(doc) {
		isValid = false
		return
	}

	cleanRG(&doc)

	isValid = getRGDigit(doc) == calcRGDigit(doc)
	return
}

func getRGDigit(doc string) int {
	switch d := rune(doc[len(doc)-1]); d {
	case 'X', 'x':
		return 10
	case '0':
		return 11
	default:
		return toInt(d)
	}
}

func calcRGDigit(doc string) int {
	sum := 0
	for i, digit := range doc[:len(doc)-1] {
		a := toInt(digit) * (i + 2)
		sum += a
	}

	return 11 - (sum % 11)
}

// cleanRG removes every rune that is not valid for RG.
func cleanRG(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) || r == 'X' || r == 'x' {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}
