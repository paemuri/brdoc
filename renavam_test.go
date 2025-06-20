package brdoc

import (
	"testing"
)

func TestIsRENAVAM(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAA", false},

		{"InvalidDigit", "38872054170", false},
		{"InvalidDigit", "40999838209", false},
		{"InvalidDigit", "31789431480", false},
		{"InvalidDigit", "38919643060", false},

		{"Valid", "13824652268", true},
		{"Valid", "08543317523", true},
		{"Valid", "09769017014", true},
		{"Valid", "01993520012", true},
		{"Valid", "04598137389", true},
		{"Valid", "05204907510", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsRENAVAM(tc.doc))
		})
	}
}
