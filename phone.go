package brdoc

import (
	"regexp"
	"strconv"
)

var (
	phoneRegexp = regexp.MustCompile(`^(?:(?:(?:\+|00)?55\s?)?(\([1-9][0-9]\)|[1-9][0-9])\s?)((9[-.\s]?\d|[2-9]{1})\d{3}[-.\s]?\d{4})$`)
)

// IsPhone verifies if `phone` is a phone number valid and return UF from DDD.
func IsPhone(phone string) (isValid bool, uf FederativeUnit) {
	matches := phoneRegexp.FindStringSubmatch(phone)
	if matches == nil {
		isValid = false
		return
	}

	match := matches[1]
	cleanNonDigits(&match)

	ddd, err := strconv.Atoi(match)
	if err != nil || ddd < 11 || ddd > 99 {
		isValid = false
		return
	}

	if ddd >= 11 && ddd <= 19 {
		return true, SP
	}
	if ddd == 21 || ddd == 22 || ddd == 24 {
		return true, RJ
	}
	if ddd == 27 || ddd == 28 {
		return true, ES
	}
	if (ddd >= 31 && ddd <= 35) || ddd == 37 || ddd == 38 {
		return true, MG
	}
	if ddd >= 41 && ddd <= 46 {
		return true, PR
	}
	if ddd >= 47 && ddd <= 49 {
		return true, SC
	}
	if ddd == 51 || (ddd >= 53 && ddd <= 55) {
		return true, RS
	}
	if ddd == 61 {
		return true, DF
	}
	if ddd == 62 || ddd == 64 {
		return true, GO
	}
	if ddd == 63 {
		return true, TO
	}
	if ddd == 65 || ddd == 66 {
		return true, MT
	}
	if ddd == 67 {
		return true, MS
	}
	if ddd == 68 {
		return true, AC
	}
	if ddd == 69 {
		return true, RO
	}
	if ddd == 71 || (ddd >= 73 && ddd <= 75) || ddd == 77 {
		return true, BA
	}
	if ddd == 79 {
		return true, SE
	}
	if ddd == 81 || ddd == 87 {
		return true, PE
	}
	if ddd == 82 {
		return true, AL
	}
	if ddd == 83 {
		return true, PB
	}
	if ddd == 84 {
		return true, RN
	}
	if ddd == 85 || ddd == 88 {
		return true, CE
	}
	if ddd == 86 {
		return true, PI
	}
	if ddd == 89 {
		return true, PI
	}
	if ddd == 91 || ddd == 93 || ddd == 94 {
		return true, PA
	}
	if ddd == 92 || ddd == 97 {
		return true, AM
	}
	if ddd == 95 {
		return true, RR
	}
	if ddd == 96 {
		return true, AP
	}
	if ddd == 98 || ddd == 99 {
		return true, MA
	}

	isValid = false
	return
}
