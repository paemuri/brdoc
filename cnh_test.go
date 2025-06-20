package brdoc

import (
	"testing"
)

func TestIsCNH(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAA", false},

		{"InvalidDigit", "02102234243", false},
		{"InvalidDigit", "02102234142", false},
		{"InvalidDigit", "13798941353", false},
		{"InvalidDigit", "00676003001", false},

		{"InvalidFormat", "8195247601-1", false},
		{"InvalidFormat", "337989413-53", false},
		{"InvalidFormat", "872 227 006 00", false},
		{"InvalidFormat", "459.911.677-05", false},

		{"Valid", "81952476011", true},
		{"Valid", "33798941353", true},
		{"Valid", "87222700600", true},
		{"Valid", "45991167705", true},
		{"Valid", "19595699996", true},
		{"Valid", "00067600300", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsCNH(tc.doc))
		})
	}
}
