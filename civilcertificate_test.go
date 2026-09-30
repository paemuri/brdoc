package brdoc

import (
	"testing"
)

func TestIsCivilCertificate(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", false},

		{"InvalidFormat", "1045390155201310001202100001232", false},
		{"InvalidFormat", "104539015520131000120210000123211", false},
		{"InvalidFormat", "104539.01.55.2013.1.00012.021.0000123/21", false},

		// The code of the civil registry (55) and the type of book (1 to 7).
		{"InvalidService", "10453901562013100012021000012321", false},
		{"InvalidBook", "10453901552013800012021000012321", false},
		{"InvalidBook", "10453901552013000012021000012321", false},

		{"InvalidDigit", "10453901552013100012021000012322", false},
		{"InvalidDigit", "10453901552013100012021000012311", false},
		// With the remainder 10 as the digit 0 instead of 1.
		{"InvalidDigit", "10453901552013100012020000000304", false},

		// Calculated by hand, as real ones are personal data.
		{"Valid", "10453901552013100012021000012321", true},
		{"Valid", "104539.01.55.2013.1.00012.021.0000123-21", true},
		{"Valid", "104539 01 55 2013 1 00012 021 0000123 21", true},
		{"Valid", "10453901552013100012020000000314", true},
		{"Valid", "00000001552024200001001000000145", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsCivilCertificate(tc.doc))
		})
	}
}
