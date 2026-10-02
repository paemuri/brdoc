package brdoc

import (
	"regexp"
)

var (
	voterIDRegexp = regexp.MustCompile(`^\d{4} ?\d{4} ?\d{4}$`)
)

// IsVoterID verifies if the given string is a valid voter ID document.
//
// The number format is defined by Article 36 of [Res. TSE 23.659/2021].
//
// [Res. TSE 23.659/2021]: https://www.tse.jus.br/legislacao/compilada/res/2021/resolucao-no-23-659-de-26-de-outubro-de-2021
func IsVoterID(doc string) bool {
	if !voterIDRegexp.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	uf := toInt(rune(doc[8]))*10 + toInt(rune(doc[9]))
	if uf < 1 || uf > 28 {
		return false
	}

	sum := 0
	for i, r := range doc[:8] {
		sum += toInt(r) * (i + 2)
	}
	digit1 := voterIDMod11(sum)

	sum = toInt(rune(doc[8]))*7 + toInt(rune(doc[9]))*8 + digit1*9
	digit2 := voterIDMod11(sum)

	return toInt(rune(doc[10])) == digit1 && toInt(rune(doc[11])) == digit2
}

// voterIDMod11 returns the remainder of `num` divided by 11, or 0 if it is 10.
func voterIDMod11(num int) int {
	mod := num % 11
	if mod == 10 {
		return 0
	}
	return mod
}
