package brdoc

import (
	"regexp"
	"strconv"
)

var (
	cepRegexp = regexp.MustCompile(`^\d{5}-?\d{3}$`)
)

// IsCEP verifies if `doc` is a valid CEP.
// Also, it validates if its related UF is part of the giving options. If none
// is provided, it validates the document for any state/district.
// This function is a wrapper on `IsCEP`.
func IsCEPFrom(doc string, ufs ...UF) bool {
	valid, uf := IsCEP(doc)
	if !valid {
		return false
	}

	return isFrom(uf, ufs)
}

// IsCEP verifies if `doc` is a valid CEP and returns its related UF.
func IsCEP(doc string) (valid bool, uf UF) {
	if !cepRegexp.MatchString(doc) {
		valid = false
		return
	}

	h, _ := strconv.Atoi(doc[0:3])

	if h >= 10 && h <= 199 {
		return true, SP
	}
	if h >= 200 && h <= 289 {
		return true, RJ
	}
	if h >= 290 && h <= 299 {
		return true, ES
	}
	if h >= 300 && h <= 399 {
		return true, MG
	}
	if h >= 400 && h <= 489 {
		return true, BA
	}
	if h >= 490 && h <= 499 {
		return true, SE
	}
	if h >= 500 && h <= 569 {
		return true, PE
	}
	if h >= 570 && h <= 579 {
		return true, AL
	}
	if h >= 580 && h <= 589 {
		return true, PB
	}
	if h >= 590 && h <= 599 {
		return true, RN
	}
	if h >= 600 && h <= 639 {
		return true, CE
	}
	if h >= 640 && h <= 649 {
		return true, PI
	}
	if h >= 650 && h <= 659 {
		return true, MA
	}
	if h >= 660 && h <= 688 {
		return true, PA
	}
	if h == 689 {
		return true, AP
	}
	if h >= 690 && h <= 692 {
		return true, AM
	}
	if h == 693 {
		return true, RR
	}
	if h >= 694 && h <= 698 {
		return true, AM
	}
	if h == 699 {
		return true, AC
	}
	if h >= 700 && h <= 727 {
		return true, DF
	}
	if h >= 728 && h <= 729 {
		return true, GO
	}
	if h >= 730 && h <= 736 {
		return true, DF
	}
	if h >= 737 && h <= 767 {
		return true, GO
	}
	if h >= 768 && h <= 769 {
		return true, RO
	}
	if h >= 770 && h <= 779 {
		return true, TO
	}
	if h >= 780 && h <= 788 {
		return true, MT
	}
	if h >= 790 && h <= 799 {
		return true, MS
	}
	if h >= 800 && h <= 879 {
		return true, PR
	}
	if h >= 880 && h <= 899 {
		return true, SC
	}
	if h >= 900 && h <= 999 {
		return true, RS
	}

	valid = false
	return
}
