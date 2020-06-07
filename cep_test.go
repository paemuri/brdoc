package brdoc

import (
	"testing"
)

func TestIsCEP(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
		u        []FederativeUnit
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654", []FederativeUnit{}},
		{"InvalidData_ShouldReturnFalse", false, "", []FederativeUnit{}},
		{"InvalidData_ShouldReturnFalse", false, "#$%¨&*(ABCDEF", []FederativeUnit{}},
		{"InvalidData_ShouldReturnFalse", false, "00000-000", []FederativeUnit{}},
		{"InvalidFormat_ShouldReturnFalse", false, "10000 000", []FederativeUnit{}},
		{"InvalidFormat_ShouldReturnFalse", false, "25000.000", []FederativeUnit{}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "10000-000", []FederativeUnit{RJ}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "25000-000", []FederativeUnit{ES}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "29500-000", []FederativeUnit{MG}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "35000-000", []FederativeUnit{BA}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "45000-000", []FederativeUnit{SE}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "49500-000", []FederativeUnit{PE}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "53000-000", []FederativeUnit{AL}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "57500-000", []FederativeUnit{PB}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "58500-000", []FederativeUnit{RN}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "59500-000", []FederativeUnit{CE}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "62000-000", []FederativeUnit{PI}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "64500-000", []FederativeUnit{MA}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "65500-000", []FederativeUnit{PA}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "67000-000", []FederativeUnit{AP}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "68900-000", []FederativeUnit{AM}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69100-000", []FederativeUnit{RR}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69300-000", []FederativeUnit{AM}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69600-000", []FederativeUnit{AC}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69900-000", []FederativeUnit{DF}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "71500-000", []FederativeUnit{GO}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "72800-000", []FederativeUnit{DF}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "73200-000", []FederativeUnit{GO}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "74500-000", []FederativeUnit{RO}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "76800-000", []FederativeUnit{TO}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "77500-000", []FederativeUnit{MT}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "78500-000", []FederativeUnit{MS}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "79500-000", []FederativeUnit{PR}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "84000-000", []FederativeUnit{SC}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "89000-000", []FederativeUnit{RS}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "95000-000", []FederativeUnit{SP}},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "29500-000", []FederativeUnit{MT, MS, MG}},
		{"Valid_ShouldReturnTrue", true, "10000-000", []FederativeUnit{SP}},
		{"Valid_ShouldReturnTrue", true, "25000-000", []FederativeUnit{RJ}},
		{"Valid_ShouldReturnTrue", true, "29500-000", []FederativeUnit{ES}},
		{"Valid_ShouldReturnTrue", true, "35000-000", []FederativeUnit{MG}},
		{"Valid_ShouldReturnTrue", true, "45000-000", []FederativeUnit{BA}},
		{"Valid_ShouldReturnTrue", true, "49500-000", []FederativeUnit{SE}},
		{"Valid_ShouldReturnTrue", true, "53000-000", []FederativeUnit{PE}},
		{"Valid_ShouldReturnTrue", true, "57500-000", []FederativeUnit{AL}},
		{"Valid_ShouldReturnTrue", true, "58500-000", []FederativeUnit{PB}},
		{"Valid_ShouldReturnTrue", true, "59500-000", []FederativeUnit{RN}},
		{"Valid_ShouldReturnTrue", true, "62000-000", []FederativeUnit{CE}},
		{"Valid_ShouldReturnTrue", true, "64500-000", []FederativeUnit{PI}},
		{"Valid_ShouldReturnTrue", true, "65500-000", []FederativeUnit{MA}},
		{"Valid_ShouldReturnTrue", true, "67000-000", []FederativeUnit{PA}},
		{"Valid_ShouldReturnTrue", true, "68900-000", []FederativeUnit{AP}},
		{"Valid_ShouldReturnTrue", true, "69100-000", []FederativeUnit{AM}},
		{"Valid_ShouldReturnTrue", true, "69300-000", []FederativeUnit{RR}},
		{"Valid_ShouldReturnTrue", true, "69600-000", []FederativeUnit{AM}},
		{"Valid_ShouldReturnTrue", true, "69900-000", []FederativeUnit{AC}},
		{"Valid_ShouldReturnTrue", true, "71500-000", []FederativeUnit{DF}},
		{"Valid_ShouldReturnTrue", true, "72800-000", []FederativeUnit{GO}},
		{"Valid_ShouldReturnTrue", true, "73200-000", []FederativeUnit{DF}},
		{"Valid_ShouldReturnTrue", true, "74500-000", []FederativeUnit{GO}},
		{"Valid_ShouldReturnTrue", true, "76800-000", []FederativeUnit{RO}},
		{"Valid_ShouldReturnTrue", true, "77500-000", []FederativeUnit{TO}},
		{"Valid_ShouldReturnTrue", true, "78500-000", []FederativeUnit{MT}},
		{"Valid_ShouldReturnTrue", true, "79500-000", []FederativeUnit{MS}},
		{"Valid_ShouldReturnTrue", true, "84000-000", []FederativeUnit{PR}},
		{"Valid_ShouldReturnTrue", true, "89000-000", []FederativeUnit{SC}},
		{"Valid_ShouldReturnTrue", true, "95000-000", []FederativeUnit{RS}},
		{"Valid_ShouldReturnTrue", true, "10000-000", []FederativeUnit{SP, SP, SP}},
		{"Valid_ShouldReturnTrue", true, "25000-000", []FederativeUnit{}},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCEP(tc.v, tc.u...))
		})
	}
}

func TestCEPRegexp(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"Invalid_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"Invalid_ShouldReturnFalse", false, ""},
		{"Invalid_ShouldReturnFalse", false, "#$%¨&*(ABCDEF"},
		{"Invalid_ShouldReturnFalse", false, "0-0000000"},
		{"Invalid_ShouldReturnFalse", false, "00-000000"},
		{"Invalid_ShouldReturnFalse", false, "000-00000"},
		{"Invalid_ShouldReturnFalse", false, "0000-0000"},
		{"Invalid_ShouldReturnFalse", false, "000000-00"},
		{"Invalid_ShouldReturnFalse", false, "0000000-0"},
		{"Invalid_ShouldReturnFalse", false, "0000x000"},
		{"Invalid_ShouldReturnFalse", false, "0000x-000"},
		{"Invalid_ShouldReturnFalse", false, "0000000x"},
		{"Invalid_ShouldReturnFalse", false, "00000-00x"},
		{"Valid_ShouldReturnTrue", true, "00000000"},
		{"Valid_ShouldReturnTrue", true, "00000-000"},
		{"Valid_ShouldReturnTrue", true, "12345678"},
		{"Valid_ShouldReturnTrue", true, "87654-321"},
		{"Valid_ShouldReturnTrue", true, "99999999"},
		{"Valid_ShouldReturnTrue", true, "99999-999"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, CEPRegexp.MatchString(tc.v))
		})
	}
}
