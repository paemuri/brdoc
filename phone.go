package brdoc

import "regexp"

var cellPhoneRegexp = regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)(?:((?:9\s?\d|[6789]{1})\d{3})(?:(\-|\s|\.))?(\d{4}))$`)
var residencialPhoneRegexp = regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)(?:((?:\s?\d|[2345]{1})\d{3})(?:(\-|\s|\.))?(\d{4}))$`)

var mapDDD = map[string]FederativeUnit{
	"61": DF,
	"62": GO,
	"64": GO,
	"65": MT,
	"66": MT,
	"67": MS,
	"82": AL,
	"71": BA,
	"73": BA,
	"74": BA,
	"75": BA,
	"77": BA,
	"85": CE,
	"88": CE,
	"98": MA,
	"99": MA,
	"83": PB,
	"81": PE,
	"87": PE,
	"86": PI,
	"89": PI,
	"84": RN,
	"79": SE,
	"68": AC,
	"96": AP,
	"92": AM,
	"97": AM,
	"91": PA,
	"93": PA,
	"94": PA,
	"69": RO,
	"95": RR,
	"63": TO,
	"27": ES,
	"28": ES,
	"31": MG,
	"32": MG,
	"33": MG,
	"34": MG,
	"35": MG,
	"37": MG,
	"38": MG,
	"21": RJ,
	"22": RJ,
	"24": RJ,
	"11": SP,
	"12": SP,
	"13": SP,
	"14": SP,
	"15": SP,
	"16": SP,
	"17": SP,
	"18": SP,
	"19": SP,
	"41": PR,
	"42": PR,
	"43": PR,
	"44": PR,
	"45": PR,
	"46": PR,
	"51": RS,
	"53": RS,
	"54": RS,
	"55": RS,
	"47": SC,
	"48": SC,
	"49": SC,
}

// IsPhone verifies if `phone` is a phone number valid and return UF from DDD.
func IsPhone(phone string) (isValid bool, uf FederativeUnit) {
	if !cellPhoneRegexp.MatchString(phone) && !residencialPhoneRegexp.MatchString(phone) {
		isValid = false
		return
	}

	cellPhoneGroups := cellPhoneRegexp.FindStringSubmatch(phone)
	residencialPhoneGroups := residencialPhoneRegexp.FindStringSubmatch(phone)
	var groups []string

	if cellPhoneGroups == nil {
		groups = residencialPhoneGroups
	} else {
		groups = cellPhoneGroups
	}

	groupDDD := 2
	uf, isValid = mapDDD[groups[groupDDD]]

	return
}
