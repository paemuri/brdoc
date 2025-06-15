package brdoc

import (
	"testing"
)

func TestIsCEP(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
		u        []UF
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654", []UF{}},
		{"InvalidData_ShouldReturnFalse", false, "", []UF{}},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAA", []UF{}},
		{"InvalidData_ShouldReturnFalse", false, "00000-000", []UF{}},
		{"InvalidFormat_ShouldReturnFalse", false, "10000 000", []UF{}},
		{"InvalidFormat_ShouldReturnFalse", false, "25000.000", []UF{}},
		{"InvalidUF_ShouldReturnFalse", false, "10000-000", []UF{RJ}},
		{"InvalidUF_ShouldReturnFalse", false, "25000-000", []UF{ES}},
		{"InvalidUF_ShouldReturnFalse", false, "29500-000", []UF{MG}},
		{"InvalidUF_ShouldReturnFalse", false, "35000-000", []UF{BA}},
		{"InvalidUF_ShouldReturnFalse", false, "45000-000", []UF{SE}},
		{"InvalidUF_ShouldReturnFalse", false, "49500-000", []UF{PE}},
		{"InvalidUF_ShouldReturnFalse", false, "53000-000", []UF{AL}},
		{"InvalidUF_ShouldReturnFalse", false, "57500-000", []UF{PB}},
		{"InvalidUF_ShouldReturnFalse", false, "58500-000", []UF{RN}},
		{"InvalidUF_ShouldReturnFalse", false, "59500-000", []UF{CE}},
		{"InvalidUF_ShouldReturnFalse", false, "62000-000", []UF{PI}},
		{"InvalidUF_ShouldReturnFalse", false, "64500-000", []UF{MA}},
		{"InvalidUF_ShouldReturnFalse", false, "65500-000", []UF{PA}},
		{"InvalidUF_ShouldReturnFalse", false, "67000-000", []UF{AP}},
		{"InvalidUF_ShouldReturnFalse", false, "68900-000", []UF{AM}},
		{"InvalidUF_ShouldReturnFalse", false, "69100-000", []UF{RR}},
		{"InvalidUF_ShouldReturnFalse", false, "69300-000", []UF{AM}},
		{"InvalidUF_ShouldReturnFalse", false, "69600-000", []UF{AC}},
		{"InvalidUF_ShouldReturnFalse", false, "69900-000", []UF{DF}},
		{"InvalidUF_ShouldReturnFalse", false, "71500-000", []UF{GO}},
		{"InvalidUF_ShouldReturnFalse", false, "72800-000", []UF{DF}},
		{"InvalidUF_ShouldReturnFalse", false, "73200-000", []UF{GO}},
		{"InvalidUF_ShouldReturnFalse", false, "74500-000", []UF{RO}},
		{"InvalidUF_ShouldReturnFalse", false, "76800-000", []UF{TO}},
		{"InvalidUF_ShouldReturnFalse", false, "77500-000", []UF{MT}},
		{"InvalidUF_ShouldReturnFalse", false, "78500-000", []UF{MS}},
		{"InvalidUF_ShouldReturnFalse", false, "79500-000", []UF{PR}},
		{"InvalidUF_ShouldReturnFalse", false, "84000-000", []UF{SC}},
		{"InvalidUF_ShouldReturnFalse", false, "89000-000", []UF{RS}},
		{"InvalidUF_ShouldReturnFalse", false, "95000-000", []UF{SP}},
		{"InvalidUF_ShouldReturnFalse", false, "29500-000", []UF{MT, MS, MG}},
		{"InvalidUF_ShouldReturnFalse", false, "00801-000", []UF{SP}},
		{"InvalidUF_ShouldReturnFalse", false, "00801-000", []UF{}},
		{"Valid_ShouldReturnTrue", true, "10000-000", []UF{SP}},
		{"Valid_ShouldReturnTrue", true, "25000-000", []UF{RJ}},
		{"Valid_ShouldReturnTrue", true, "29500-000", []UF{ES}},
		{"Valid_ShouldReturnTrue", true, "35000-000", []UF{MG}},
		{"Valid_ShouldReturnTrue", true, "45000-000", []UF{BA}},
		{"Valid_ShouldReturnTrue", true, "49500-000", []UF{SE}},
		{"Valid_ShouldReturnTrue", true, "53000-000", []UF{PE}},
		{"Valid_ShouldReturnTrue", true, "57500-000", []UF{AL}},
		{"Valid_ShouldReturnTrue", true, "58500-000", []UF{PB}},
		{"Valid_ShouldReturnTrue", true, "59500-000", []UF{RN}},
		{"Valid_ShouldReturnTrue", true, "62000-000", []UF{CE}},
		{"Valid_ShouldReturnTrue", true, "64500-000", []UF{PI}},
		{"Valid_ShouldReturnTrue", true, "65500-000", []UF{MA}},
		{"Valid_ShouldReturnTrue", true, "67000-000", []UF{PA}},
		{"Valid_ShouldReturnTrue", true, "68900-000", []UF{AP}},
		{"Valid_ShouldReturnTrue", true, "69100-000", []UF{AM}},
		{"Valid_ShouldReturnTrue", true, "69300-000", []UF{RR}},
		{"Valid_ShouldReturnTrue", true, "69600-000", []UF{AM}},
		{"Valid_ShouldReturnTrue", true, "69900-000", []UF{AC}},
		{"Valid_ShouldReturnTrue", true, "71500-000", []UF{DF}},
		{"Valid_ShouldReturnTrue", true, "72800-000", []UF{GO}},
		{"Valid_ShouldReturnTrue", true, "73200-000", []UF{DF}},
		{"Valid_ShouldReturnTrue", true, "74500-000", []UF{GO}},
		{"Valid_ShouldReturnTrue", true, "76800-000", []UF{RO}},
		{"Valid_ShouldReturnTrue", true, "77500-000", []UF{TO}},
		{"Valid_ShouldReturnTrue", true, "78500-000", []UF{MT}},
		{"Valid_ShouldReturnTrue", true, "79500-000", []UF{MS}},
		{"Valid_ShouldReturnTrue", true, "84000-000", []UF{PR}},
		{"Valid_ShouldReturnTrue", true, "89000-000", []UF{SC}},
		{"Valid_ShouldReturnTrue", true, "95000-000", []UF{RS}},
		{"Valid_ShouldReturnTrue", true, "10000-000", []UF{SP, SP, SP}},
		{"Valid_ShouldReturnTrue", true, "25000-000", []UF{}},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCEPFrom(tc.v, tc.u...))
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
		{"Invalid_ShouldReturnFalse", false, "AAAAAAAA"},
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
			assertEqual(t, tc.expected, cepRegexp.MatchString(tc.v))
		})
	}
}
