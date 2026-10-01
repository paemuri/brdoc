package brdoc

import (
	"regexp"
)

var (
	bankSlipBarcodeRegexp = regexp.MustCompile(`^\d{44}$`)
	bankSlipLineRegexp    = regexp.MustCompile(
		`^\d{5}\.?\d{5} ?\d{5}\.?\d{6} ?\d{5}\.?\d{6} ?\d ?\d{14}$`,
	)
	// Not official logic: the mask of the line of a revenue slip.
	bankSlipRevenueLineRegexp = regexp.MustCompile(
		`^\d{11}-?\d( ?\d{11}-?\d){3}$`,
	)
)

// IsBankSlip verifies if the given string is a valid bank slip, either its
// barcode or its digitable line. It works for both the slips of banks and the
// slips of utility bills and taxes (revenue slips).
//
// The format and the check digits of the slips of banks are defined by
// [Carta-Circular BACEN 2.926/2000], and the code 988 of the institutions
// identified only by the ISPB by Annex V of the [Convenção FEBRABAN]. The
// format and the check digits of the revenue slips are defined by the
// [Layout FEBRABAN].
//
// [Carta-Circular BACEN 2.926/2000]: https://www.bcb.gov.br/estabilidadefinanceira/exibenormativo?tipo=Carta%20Circular&numero=2926
// [Convenção FEBRABAN]: https://cmsarquivos.febraban.org.br/Arquivos/documentos/PDF/Conven%C3%A7%C3%A3o%20da%20Cobran%C3%A7a%20-%2005_02_2021_f.pdf
// [Layout FEBRABAN]: https://cmsarquivos.febraban.org.br/Arquivos/documentos/PDF/Layout%20-%20C%C3%B3digo%20de%20Barras%20-%20Vers%C3%A3o%208%20-%2011_05_2026.pdf
func IsBankSlip(doc string) bool {
	switch {
	case bankSlipBarcodeRegexp.MatchString(doc):
		if doc[0] == '8' {
			return validBankSlipRevenue(doc)
		}
		return validBankSlip(doc)
	case bankSlipLineRegexp.MatchString(doc):
		cleanNonDigits(&doc)
		if doc[0] == '8' {
			return false
		}
		return validBankSlipLine(doc)
	case bankSlipRevenueLineRegexp.MatchString(doc):
		cleanNonDigits(&doc)
		if doc[0] != '8' {
			return false
		}
		return validBankSlipRevenueLine(doc)
	}

	return false
}

// validBankSlip verifies the barcode of a slip of a bank. The currency is 9
// (real), or 0 for the code 988, whose due date factor is always 0000.
func validBankSlip(doc string) bool {
	if doc[:3] == "988" {
		if doc[3] != '0' || doc[5:9] != "0000" {
			return false
		}
	} else if doc[3] != '9' {
		return false
	}

	return toInt(rune(doc[4])) == calcBankSlipDigit(doc[:4]+doc[5:])
}

// validBankSlipLine verifies the digitable line of a slip of a bank, whose 3
// first fields have a check digit each, and then the barcode it represents.
func validBankSlipLine(doc string) bool {
	// Not official logic: the calculation of the modulo 10, which is detailed
	// only for revenue slips.
	for _, field := range [][2]int{{0, 9}, {10, 20}, {21, 31}} {
		digit := calcMod10Digit(doc[field[0]:field[1]])
		if toInt(rune(doc[field[1]])) != digit {
			return false
		}
	}

	return validBankSlip(doc[:4] + doc[32:] + doc[4:9] + doc[10:20] +
		doc[21:31])
}

// validBankSlipRevenue verifies the barcode of a revenue slip. The check digit
// is calculated with modulo 10 or 11, depending on the third digit.
func validBankSlipRevenue(doc string) bool {
	// The segments are from 1 to 7, and 9.
	if doc[1] == '0' || doc[1] == '8' {
		return false
	}

	calcDigit := bankSlipRevenueCalcDigit(doc[2])
	if calcDigit == nil {
		return false
	}

	return toInt(rune(doc[3])) == calcDigit(doc[:3]+doc[4:])
}

// validBankSlipRevenueLine verifies the digitable line of a revenue slip, whose
// 4 blocks of 11 digits have a check digit each, and then the barcode it
// represents.
func validBankSlipRevenueLine(doc string) bool {
	calcDigit := bankSlipRevenueCalcDigit(doc[2])
	if calcDigit == nil {
		return false
	}

	for i := 0; i < 48; i += 12 {
		if toInt(rune(doc[i+11])) != calcDigit(doc[i:i+11]) {
			return false
		}
	}

	return validBankSlipRevenue(doc[:11] + doc[12:23] + doc[24:35] +
		doc[36:47])
}

// bankSlipRevenueCalcDigit returns the check digit calculation of a revenue
// slip, by the identifier of its value (6 or 7 for modulo 10, and 8 or 9 for
// modulo 11), or nil if the identifier is not valid.
func bankSlipRevenueCalcDigit(id byte) func(string) int {
	switch id {
	case '6', '7':
		return calcMod10Digit
	case '8', '9':
		return calcMod11Digit
	}

	return nil
}

// calcBankSlipDigit returns 11 minus the remainder of the weighted sum of `doc`
// divided by 11, or 1 if it is 10 or 11. The weights go from 2 to 9, from right
// to left.
func calcBankSlipDigit(doc string) int {
	digit := 11 - calcMod11Sum(doc)%11
	if digit >= 10 {
		return 1
	}

	return digit
}
