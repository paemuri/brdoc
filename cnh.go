package brdoc

import (
	"strconv"
)

// IsCNH verifies if the given string is a valid CNH
// document.
func IsCNH(doc string) bool {
	runes := []rune(doc)
	if len(runes) != 11 {
		return false
	}
	vBase := 0
	vSoma := 0
	vFator := 9
	for i := 0; i < 9; i++ {
		n, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}
		vSoma = vSoma + (n * vFator)
		vFator--
	}
	vResto := vSoma % 11
	if vResto == 10 {
		vBase = -2
	}
	if vResto > 9 {
		vResto = 0
	}
	sResultado := strconv.Itoa(vResto)
	vSoma = 0
	vFator = 1
	for i := 0; i < 9; i++ {
		n, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}
		vSoma = vSoma + (n * vFator)
		vFator++
	}
	if (vSoma%11)+vBase < 0 {
		vResto = 11 + (vSoma % 11) + vBase
	}
	if (vSoma%11)+vBase >= 0 {
		vResto = (vSoma % 11) + vBase
	}
	if vResto > 9 {
		vResto = 0
	}
	sResultado = sResultado + strconv.Itoa(vResto)
	if string(runes[9:]) != sResultado {
		return false
	}
	return true
}
