package brdoc

// IsCNS verifies if the given string is a valid CNS document.
func IsCNS(doc string) bool {

	cleanNonDigits(&doc)
	if len(doc) != 15 {
		return false
	}

	sum := 0

	for i, r := range doc {
		sum += toInt(r) * (15 - i)
	}
	return sum%11 == 0
}
