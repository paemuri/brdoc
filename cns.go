package brdoc

import (
	"regexp"
)

// Regexp pattern for CNS.
var (
	CNSRegexp = regexp.MustCompile(`^([12]\d{2}\s?\d{4}\s?\d{4}\s?00[01]\d|[789]\d{2}\s?\d{4}\s?\d{4}\s?\d{4})$`)
)

// IsCNS verifies if the given string is a valid CNS document.
func IsCNS(doc string) bool {

	if !CNSRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	sum := 0
	for i, r := range doc {
		sum += toInt(r) * (15 - i)
	}

	return sum%11 == 0
}
