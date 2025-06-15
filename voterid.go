package brdoc

import (
	"fmt"
	"strconv"
)

// IsVoterID verifies if the given string is a valid voter ID document.
func IsVoterID(doc string) bool {
	// This function was based on the logic from [1]. It seems to be a bit
	// sketchy, but it works for now.
	// [1]: http://ghiorzi.org/DVnew.htm#e.

	if len(doc) != 12 {
		return false
	}
	if !allDigit(doc) {
		return false
	}

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
	if mod == 10 || mod == 11 {
		return 0
	}
	return mod
}
