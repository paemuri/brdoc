package brdoc

import (
	"regexp"
	"strconv"
)

var (
	phoneRegexp = regexp.MustCompile(
		`^(?:(?:(?:\+|00)?55\s?)?(\([1-9][0-9]\)|[1-9][0-9])\s?)(([7-9][-.\s]?\d|[2-6])\d{3}[-.\s]?\d{4})$`,
	)
)

// phoneDDDs maps each area code to its UFs.
var phoneDDDs = map[int][]UF{
	11: {SP},
	12: {SP},
	13: {SP},
	14: {SP},
	15: {SP},
	16: {SP},
	17: {SP},
	18: {SP},
	19: {SP},
	21: {RJ},
	22: {RJ},
	24: {RJ},
	27: {ES},
	28: {ES},
	31: {MG},
	32: {MG},
	33: {MG},
	34: {MG},
	35: {MG},
	37: {MG},
	38: {MG},
	41: {PR},
	42: {PR, SC},
	43: {PR},
	44: {PR},
	45: {PR},
	46: {PR},
	47: {SC, PR},
	48: {SC},
	49: {SC, PR},
	51: {RS},
	53: {RS},
	54: {RS},
	55: {RS},
	61: {DF, GO},
	62: {GO},
	63: {TO},
	64: {GO},
	65: {MT},
	66: {MT},
	67: {MS},
	68: {AC},
	69: {RO},
	71: {BA},
	73: {BA},
	74: {BA},
	75: {BA},
	77: {BA},
	79: {SE},
	81: {PE},
	82: {AL},
	83: {PB},
	84: {RN},
	85: {CE},
	86: {PI},
	87: {PE},
	88: {CE},
	89: {PI},
	91: {PA},
	92: {AM},
	93: {PA},
	94: {PA},
	95: {RR},
	96: {AP},
	97: {AM},
	98: {MA},
	99: {MA},
}

// IsPhoneFrom verifies if `phone` is a valid Brazilian phone number. Also, it
// validates if any of its related UFs is part of the given options. If none is
// provided, it validates the phone for any state/district. This function is a
// wrapper around [IsPhone].
func IsPhoneFrom(phone string, ufs ...UF) bool {
	valid, phoneUFs := IsPhone(phone)
	if !valid {
		return false
	}

	for _, uf := range phoneUFs {
		if isFrom(uf, ufs) {
			return true
		}
	}

	return false
}

// IsPhone verifies if `phone` is a valid Brazilian phone number and returns all
// the UFs related to its area code.
//
// The numbering rules are defined by Articles 11, 12 and 15 of
// [Res. Anatel 749/2022], and the area codes of each UF are listed by [Anatel].
//
// [Res. Anatel 749/2022]: https://informacoes.anatel.gov.br/legislacao/resolucoes/2022/1641-resolucao-749
// [Anatel]: https://www.anatel.gov.br/dadosabertos/PDA/Codigo_Nacional/PGCN.csv
func IsPhone(phone string) (valid bool, ufs []UF) {
	matches := phoneRegexp.FindStringSubmatch(phone)
	if matches == nil {
		return false, nil
	}

	match := matches[1]
	cleanNonDigits(&match)

	ddd, err := strconv.Atoi(match)
	if err != nil {
		return false, nil
	}

	ufs, valid = phoneDDDs[ddd]
	if !valid {
		return false, nil
	}

	return true, append([]UF(nil), ufs...)
}
