package brdoc

import (
	"strconv"
)

var (
	vSum = 0
	vDV = 0
	vNumber = 0
	vDigit = 0
	vSequence = []int{3,2,9,8,7,6,5,4,3,2}
)

func IsRenavam(doc string) bool {
	
	if len(doc) < 11 {
		return false
	}
		
	for i := 0; i <= 9; i++ {
		vNumber, err := strconv.Atoi(string(doc[i]))
		if err != nil {
			return false
		}		
		vSum = vSum + (vNumber * vSequence[i])
	}

	vDV = (vSum * 10) % 11
	if vDV == 10 {
		vDV = 0
	}

	vDigit, err := strconv.Atoi(string(doc[10]))
	if err != nil {
		return false
	}		
  
	return (vDV == vDigit)
}