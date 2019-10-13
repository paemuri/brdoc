package brdoc

import (
	"strconv"
)

var (
	vSum      = 0
	vDV       = 0
	vNumber   = 0
	vSequence = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsRenavam verifies if the given string is a valid Renavam document.
func IsRenavam(doc string) bool {

	runes := []rune(doc)
	if len(runes) != 11 {
		return false
	}

	vNumber = 0
	vSum = 0
	vDV = 0
	for i := 0; i <= 9; i++ {
		vNumber, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}
		vSum = vSum + (vNumber * vSequence[i])
	}

	vDV = (vSum * 10) % 11
	if vDV == 10 {
		vDV = 0
	}

	return (string(runes[10]) == strconv.Itoa(vDV))
}
