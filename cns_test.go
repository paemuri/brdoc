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
		{"InvalidDigit_ShouldReturnFalse", false, "915501701930306"},
		{"InvalidDigit_ShouldReturnFalse", false, "174224171330004"},
		{"InvalidDigit_ShouldReturnFalse", false, "259755733880001"},
		{"InvalidDigit_ShouldReturnFalse", false, "259755-33880001"},
		{"InvalidDigit_ShouldReturnFalse", false, "2597550R3880001"},
		{"Valid_ShouldReturnTrue", true, "259955733880001"},
		{"Valid_ShouldReturnTrue", true, "174524171330004"},
		{"Valid_ShouldReturnTrue", true, "915501701930006"},
		{"Valid_ShouldReturnTrue", true, "750655718440005"},
		{"Valid_ShouldReturnTrue", true, "864397336660007"},
		{"Valid_ShouldReturnTrue", true, "285054561330005"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCNS(tc.v))
		})
	}
}
