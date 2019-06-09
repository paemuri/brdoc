package brdoc

// IsPlate verifies if the given string is a valid
// license plate, either in the old national format
// or the new Mercosul one.
func IsPlate(doc string) bool {
	return IsNationalPlate(doc) || IsMercosulPlate(doc)
}

// IsNationalPlate verifies if the given string is a
// valid license plate in the old national format.
func IsNationalPlate(doc string) bool {
	return false
}

// IsMercosulPlate verifies if the given string is a
// valid license plate in the new Mercosul format.
func IsMercosulPlate(doc string) bool {
	return false
}
