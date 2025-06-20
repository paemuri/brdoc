package brdoc

import (
	"testing"
)

func TestIsCNS(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAA", false},

		{"InvalidDigit", "915 5017 0193 0306", false},
		{"InvalidDigit", "174 2241 7133 0004", false},
		{"InvalidDigit", "259 7557 3388 0001", false},

		{"InvalidFormat", "808-2536-1743-0486", false},
		{"InvalidFormat", "9999 0236 0200 834", false},

		{"Valid", "174 5984 3528 0018", true},
		{"Valid", "259 9557 3388 0001", true},
		{"Valid", "174 5241 7133 0004", true},
		{"Valid", "915 5017 0193 0006", true},
		{"Valid", "750 6557 1844 0005", true},
		{"Valid", "864 3973 3666 0007", true},
		{"Valid", "285 0545 6133 0005", true},
		// No spaces.
		{"Valid", "174598435280018", true},
		{"Valid", "259955733880001", true},
		{"Valid", "174524171330004", true},
		{"Valid", "915501701930006", true},
		{"Valid", "750655718440005", true},
		{"Valid", "864397336660007", true},
		{"Valid", "285054561330005", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsCNS(tc.doc))
		})
	}
}
