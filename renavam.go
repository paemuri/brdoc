package brdoc

import (
	"strconv"
)

var (
	renavamSequence = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsRENAVAM verifies if the given string is a valid RENAVAM document.
func IsRENAVAM(doc string) bool {

	runes := []rune(doc)
	if len(runes) != 11 {
		return false
	}

	vSum := 0
	for i := 0; i <= 9; i++ {
		vNumber, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}
		vSum = vSum + (vNumber * renavamSequence[i])
	}

	vDV := (vSum * 10) % 11
	if vDV == 10 {
		vDV = 0
	}

	return (string(runes[10]) == strconv.Itoa(vDV))
}
