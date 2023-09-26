package brdoc

// IsVoterID verifies if the given string is a valid voter registration document
func IsVoterID(doc string) bool {
	if !allDigit(doc) || len(doc) != 12 {
		return false
	}

	sumA := 0
	for i, digit := range doc[:len(doc)-4] {
		sumA += toInt(digit) * (i + 2)
	}
	dv1 := dvMod11(sumA)

	sumB := toInt(rune(doc[8]))*7 + toInt(rune(doc[9]))*8 + dv1*9
	dv2 := dvMod11(sumB)

	return dv1 == toInt(rune(doc[10])) && dv2 == toInt(rune(doc[11]))
}

func dvMod11(num int) int {
	mod := num % 11
	if mod == 10 || mod == 11 {
		return 0
	}
	return mod
}
