package brdoc

import (
	"regexp"
	"strconv"
)

var cellPhoneRegexp = regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)(?:((?:9\s?\d|[6789]{1})\d{3})(?:(\-|\s|\.))?(\d{4}))$`)
var residentialPhoneRegexp = regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)(?:((?:\s?\d|[2345]{1})\d{3})(?:(\-|\s|\.))?(\d{4}))$`)

var mapDDD = map[int]FederativeUnit{
	11: SP,
	12: SP,
	13: SP,
	14: SP,
	15: SP,
	16: SP,
	17: SP,
	18: SP,
	19: SP,
	21: RJ,
	22: RJ,
	24: RJ,
	27: ES,
	28: ES,
	31: MG,
	32: MG,
	33: MG,
	34: MG,
	35: MG,
	37: MG,
	38: MG,
	41: PR,
	42: PR,
	43: PR,
	44: PR,
	45: PR,
	46: PR,
	47: SC,
	48: SC,
	49: SC,
	51: RS,
	53: RS,
	54: RS,
	55: RS,
	61: DF,
	62: GO,
	63: TO,
	64: GO,
	65: MT,
	66: MT,
	67: MS,
	68: AC,
	69: RO,
	71: BA,
	73: BA,
	74: BA,
	75: BA,
	77: BA,
	79: SE,
	81: PE,
	82: AL,
	83: PB,
	84: RN,
	85: CE,
	86: PI,
	87: PE,
	88: CE,
	89: PI,
	91: PA,
	92: AM,
	93: PA,
	94: PA,
	95: RR,
	96: AP,
	97: AM,
	98: MA,
	99: MA,
}

// IsPhone verifies if `phone` is a phone number valid and return UF from DDD.
func IsPhone(phone string) (isValid bool, uf FederativeUnit) {
	if !cellPhoneRegexp.MatchString(phone) && !residentialPhoneRegexp.MatchString(phone) {
		isValid = false
		return
	}

	cellPhoneGroups := cellPhoneRegexp.FindStringSubmatch(phone)
	residentialPhoneGroups := residentialPhoneRegexp.FindStringSubmatch(phone)
	var groups []string

	if residentialPhoneGroups != nil {
		groups = residentialPhoneGroups
	} else {
		groups = cellPhoneGroups
	}

	ddd, err := strconv.Atoi(groups[2])
	if err != nil || ddd < 11 || ddd > 99 {
		isValid = false
		return
	}

	uf, isValid = mapDDD[ddd]

	return
}
