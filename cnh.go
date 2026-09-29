package brdoc

// IsCNH verifies if the given string is a valid CNH document.
func IsCNH(doc string) bool {
	// The number format is defined by Article 4 of [1].
	// [1]: https://www.legisweb.com.br/legislacao/?id=425117.

	if len(doc) != 11 {
		return false
	}
	if !allDigit(doc) {
		return false
	}
	// Not official logic: reject documents with all digits equal.
	if allEq(doc) {
		return false
	}

	// Not official logic: check digits calculation.
	sum := 0
	acc := 9
	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc--
	}

	base := 0
	digit1 := sum % 11
	if digit1 == 10 {
		base = -2
	}
	if digit1 > 9 {
		digit1 = 0
	}

	sum = 0
	acc = 1
	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc++
	}

	var digit2 int
	if (sum%11)+base < 0 {
		digit2 = 11 + (sum % 11) + base
	}
	if (sum%11)+base >= 0 {
		digit2 = (sum % 11) + base
	}
	if digit2 > 9 {
		digit2 = 0
	}

	return toInt(rune(doc[len(doc)-2])) == digit1 &&
		toInt(rune(doc[len(doc)-1])) == digit2
}
