package brdoc

import (
	"regexp"
)

// Regexp pattern for license plates patterns.
var (
	NationalPlateRegexp = regexp.MustCompile(`^[A-Z]{3}-?\d{4}$`)
	MercosulPlateRegexp = regexp.MustCompile(`^[A-Z]{3}\d[A-Z]\d{2}$`)
)

// IsPlate verifies if the given string is a valid license plate.IsPlate
// It can be either in the old national format or the new Mercosul one.
func IsPlate(doc string) bool {
	return IsNationalPlate(doc) || IsMercosulPlate(doc)
}

// IsNationalPlate verifies if the given string is a valid license plate in the old national format.
func IsNationalPlate(doc string) bool {
	return NationalPlateRegexp.MatchString(doc)
}

// IsMercosulPlate verifies if the given string is a valid license plate in the new Mercosul format.
func IsMercosulPlate(doc string) bool {
	return MercosulPlateRegexp.MatchString(doc)
}
