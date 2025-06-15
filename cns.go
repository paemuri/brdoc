package brdoc

import (
	"regexp"
)

var (
	cnsRegexp = regexp.MustCompile(
		`^([12]\d{2}(\s?\d{4}){2}\s?00[01]\d|[789]\d{2}(\s?\d{4}){3})$`,
	)
)

// IsCNS verifies if the given string is a valid CNS document.
func IsCNS(doc string) bool {
	if !cnsRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	sum := 0
	for i, r := range doc {
		sum += toInt(r) * (15 - i)
	}

	return sum%11 == 0
}
