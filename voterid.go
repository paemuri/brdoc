package brdoc

import (
	"fmt"
	"regexp"
	"strconv"
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

	docRune := []rune(doc)
	docUF := fmt.Sprintf("%d%d", toInt(docRune[8]), toInt(docRune[9]))
	docUFInt, _ := strconv.Atoi(docUF)
	if docUFInt < 1 || docUFInt > 28 {
		return false
	}

	sumA := 0
	for i, digit := range doc[:len(doc)-4] {
		sumA += toInt(digit) * (i + 2)
	}
	dv1 := voterIDMod11(sumA)

	sumB := toInt(docRune[8])*7 + toInt(docRune[9])*8 + dv1*9
	dv2 := voterIDMod11(sumB)

	return dv1 == toInt(docRune[10]) && dv2 == toInt(docRune[11])
}

func voterIDMod11(num int) int {
	mod := num % 11
	if mod == 10 {
		return 0
	}
	return mod
}
