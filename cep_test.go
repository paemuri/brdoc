package brdoc

import (
	"testing"
)

func TestIsCEP(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
		u        FederativeUnit
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654", 0},
		{"InvalidData_ShouldReturnFalse", false, "", 0},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAA", 0},
		{"InvalidData_ShouldReturnFalse", false, "00000-000", 0},
		{"InvalidFormat_ShouldReturnFalse", false, "10000 000", 0},
		{"InvalidFormat_ShouldReturnFalse", false, "25000.000", 0},
		{"InvalidFederativeUnit_ShouldReturnTrue", true, "10000-000", 0},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "25000-000", ES},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "29500-000", MG},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "35000-000", BA},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "45000-000", SE},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "49500-000", PE},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "53000-000", AL},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "57500-000", PB},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "58500-000", RN},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "59500-000", CE},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "62000-000", PI},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "64500-000", MA},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "65500-000", PA},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "67000-000", AP},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "68900-000", AM},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69100-000", RR},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69300-000", AM},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69600-000", AC},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "69900-000", DF},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "71500-000", GO},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "72800-000", DF},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "73200-000", GO},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "74500-000", RO},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "76800-000", TO},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "77500-000", MT},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "78500-000", MS},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "79500-000", PR},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "84000-000", SC},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "89000-000", RS},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "95000-000", SP},
		{"InvalidFederativeUnit_ShouldReturnFalse", false, "29500-000", MG},
		{"Valid_ShouldReturnTrue", true, "10000-000", SP},
		{"Valid_ShouldReturnTrue", true, "25000-000", RJ},
		{"Valid_ShouldReturnTrue", true, "29500-000", ES},
		{"Valid_ShouldReturnTrue", true, "35000-000", MG},
		{"Valid_ShouldReturnTrue", true, "45000-000", BA},
		{"Valid_ShouldReturnTrue", true, "49500-000", SE},
		{"Valid_ShouldReturnTrue", true, "53000-000", PE},
		{"Valid_ShouldReturnTrue", true, "57500-000", AL},
		{"Valid_ShouldReturnTrue", true, "58500-000", PB},
		{"Valid_ShouldReturnTrue", true, "59500-000", RN},
		{"Valid_ShouldReturnTrue", true, "62000-000", CE},
		{"Valid_ShouldReturnTrue", true, "64500-000", PI},
		{"Valid_ShouldReturnTrue", true, "65500-000", MA},
		{"Valid_ShouldReturnTrue", true, "67000-000", PA},
		{"Valid_ShouldReturnTrue", true, "68900-000", AP},
		{"Valid_ShouldReturnTrue", true, "69100-000", AM},
		{"Valid_ShouldReturnTrue", true, "69300-000", RR},
		{"Valid_ShouldReturnTrue", true, "69600-000", AM},
		{"Valid_ShouldReturnTrue", true, "69900-000", AC},
		{"Valid_ShouldReturnTrue", true, "71500-000", DF},
		{"Valid_ShouldReturnTrue", true, "72800-000", GO},
		{"Valid_ShouldReturnTrue", true, "73200-000", DF},
		{"Valid_ShouldReturnTrue", true, "74500-000", GO},
		{"Valid_ShouldReturnTrue", true, "76800-000", RO},
		{"Valid_ShouldReturnTrue", true, "77500-000", TO},
		{"Valid_ShouldReturnTrue", true, "78500-000", MT},
		{"Valid_ShouldReturnTrue", true, "79500-000", MS},
		{"Valid_ShouldReturnTrue", true, "84000-000", PR},
		{"Valid_ShouldReturnTrue", true, "89000-000", SC},
		{"Valid_ShouldReturnTrue", true, "95000-000", RS},
		{"Valid_ShouldReturnTrue", true, "10000-000", SP},
		{"Valid_ShouldReturnTrue", true, "25000-000", 0},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCEP(tc.v, tc.u))
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
			assertEqual(t, tc.expected, CEPRegexp.MatchString(tc.v))
		})
	}
}
