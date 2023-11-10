package brdoc

import (
	"strconv"
)

func isValidTituloEleitoral(numeroTitulo string) bool {
	if _, err := strconv.Atoi(numeroTitulo); err != nil {
		return false
	}

	sequentialNumber, uf, digVerifiers := getTituloEleitoralParts(numeroTitulo)

	if !verifyLength(numeroTitulo, uf) {
		return false
	}

	dv1 := calculateDV1(sequentialNumber)
	verifiedDV1 := verifyDV1(dv1, uf, digVerifiers)
	verifiedDV2 := verifyDV2(uf, dv1, digVerifiers)
	verifiedUF := verifyUF(uf)

	return verifiedDV1 && verifiedDV2 && verifiedUF
}

func getTituloEleitoralParts(inputString string) (string, string, string) {
	sequentialNumber := inputString[:8]
	uf := inputString[len(inputString)-4 : len(inputString)-2]
	digVerifiers := inputString[len(inputString)-2:]

	return sequentialNumber, uf, digVerifiers
}

func verifyLength(numeroTitulo string, uf string) bool {
	if len(numeroTitulo) == 12 {
		return true
	}

	return len(numeroTitulo) == 13 && (uf == "01" || uf == "02")
}

func calculateDV1(sequentialNumber string) int {
	x := []int{2, 3, 4, 5, 6, 7, 8, 9}

	v1 := 0
	for i := 0; i < 8; i++ {
		num, _ := strconv.Atoi(string(sequentialNumber[i]))
		v1 += num * x[i]
	}

	dv1 := v1 % 11

	return dv1
}

func verifyDV1(dv1 int, uf string, digVerifiers string) bool {
	if dv1 == 0 && (uf == "01" || uf == "02") {
		dv1 = 1
	}

	if dv1 == 10 {
		dv1 = 0
	}

	return int(digVerifiers[0]-'0') == dv1
}

func verifyDV2(uf string, dv1 int, digVerifiers string) bool {
	x := []int{7, 8, 9}
	v2 := (int(uf[0]-'0') * x[0]) + (int(uf[1]-'0') * x[1]) + (dv1 * x[2])

	dv2 := v2 % 11

	if (uf == "01" || uf == "02") && dv2 == 0 {
		dv2 = 1
	}

	return int(digVerifiers[1]-'0') == dv2
}

func verifyUF(uf string) bool {
	validUFs := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28"}
	for _, validUF := range validUFs {
		if uf == validUF {
			return true
		}
	}
	return false
}
