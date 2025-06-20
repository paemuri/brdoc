package brdoc

import (
	"testing"
)

func TestIsPIS(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAA", false},

		{"InvalidDigit", "103.95199.01-6", false},
		{"InvalidDigit", "120.1641.414-9", false},

		{"InvalidFormat", "103.951.990-15", false},
		{"InvalidFormat", "103 951 990 15", false},
		{"InvalidFormat", "103951-99015", false},

		{"Valid", "103.951.9901-5", true},
		{"Valid", "103.95199.01-5", true},
		{"Valid", "10395199015", true},
		{"Valid", "120.6372.482-4", true},
		{"Valid", "120.1641.414-8", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsPIS(tc.doc))
		})
	}
}
