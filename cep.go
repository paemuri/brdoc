package brdoc

import (
	"regexp"
	"strconv"
)

var (
	CEPRegexp = regexp.MustCompile(`^\d{5}-?\d{3}$`)
)

// IsCEP verifies if `doc` is a valid CEP.
// `ufs` represents the possible Federative Units the CEP should matches.
// If none is provided, it validates the document for any state/district.
func IsCEP(doc string, ufs ...FederativeUnit) bool {
	if !CEPRegexp.MatchString(doc) {
		return false
	}

	h, _ := strconv.Atoi(doc[0:3])

	// If no UF is specified, anything between 010 and 999 is valid.
	if len(ufs) == 0 {
		return h >= 10
	}

	for _, uf := range ufs {
		if (uf == SP && h >= 10 && h <= 199) ||
			(uf == RJ && h >= 200 && h <= 289) ||
			(uf == ES && h >= 290 && h <= 299) ||
			(uf == MG && h >= 300 && h <= 399) ||
			(uf == BA && h >= 400 && h <= 489) ||
			(uf == SE && h >= 490 && h <= 499) ||
			(uf == PE && h >= 500 && h <= 569) ||
			(uf == AL && h >= 570 && h <= 579) ||
			(uf == PB && h >= 580 && h <= 589) ||
			(uf == RN && h >= 590 && h <= 599) ||
			(uf == CE && h >= 600 && h <= 639) ||
			(uf == PI && h >= 640 && h <= 649) ||
			(uf == MA && h >= 650 && h <= 659) ||
			(uf == PA && h >= 660 && h <= 688) ||
			(uf == AP && h == 689) ||
			(uf == AM && h >= 690 && h <= 692) ||
			(uf == RR && h == 693) ||
			(uf == AM && h >= 694 && h <= 698) ||
			(uf == AC && h == 699) ||
			(uf == DF && h >= 700 && h <= 727) ||
			(uf == GO && h >= 728 && h <= 729) ||
			(uf == DF && h >= 730 && h <= 736) ||
			(uf == GO && h >= 737 && h <= 767) ||
			(uf == RO && h >= 768 && h <= 769) ||
			(uf == TO && h >= 770 && h <= 779) ||
			(uf == MT && h >= 780 && h <= 788) ||
			(uf == MS && h >= 790 && h <= 799) ||
			(uf == PR && h >= 800 && h <= 879) ||
			(uf == SC && h >= 880 && h <= 899) ||
			(uf == RS && h >= 900 && h <= 999) {

			return true
		}
	}

	return false
}
