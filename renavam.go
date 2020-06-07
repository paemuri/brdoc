package brdoc

import (
	"strconv"
)

var (
	renavamAcc = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsRENAVAM verifies if the given string is a valid RENAVAM document.
func IsRENAVAM(doc string) bool {

	if len(doc) != 11 {
		return false
	}
	if !allDigit(doc) {
		return false
	}

	var sum int
	for i, r := range doc[:len(doc)-1] {
		sum += toInt(r) * renavamAcc[i]
	}

	digit := (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}

	return string(doc[len(doc)-1]) == strconv.Itoa(digit)
}
