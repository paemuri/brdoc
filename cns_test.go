package brdoc

import (
	"testing"
)

func TestIsCNS(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAAAAA"},
		{"InvalidDigit_ShouldReturnFalse", false, "915 5017 0193 0306"},
		{"InvalidDigit_ShouldReturnFalse", false, "174 2241 7133 0004"},
		{"InvalidDigit_ShouldReturnFalse", false, "259 7557 3388 0001"},
		{"InvalidFormat_ShouldReturnFalse", false, "808-2536-1743-0486"},
		{"InvalidFormat_ShouldReturnFalse", false, "9999 0236 0200 834"},
		{"Valid_ShouldReturnTrue", true, "174 5984 3528 0018"},
		{"Valid_ShouldReturnTrue", true, "259 9557 3388 0001"},
		{"Valid_ShouldReturnTrue", true, "174 5241 7133 0004"},
		{"Valid_ShouldReturnTrue", true, "915 5017 0193 0006"},
		{"Valid_ShouldReturnTrue", true, "750 6557 1844 0005"},
		{"Valid_ShouldReturnTrue", true, "864 3973 3666 0007"},
		{"Valid_ShouldReturnTrue", true, "285 0545 6133 0005"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCNS(tc.v))
		})
	}
}
