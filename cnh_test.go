package brdoc

import (
	"testing"
)

func TestIsCNH(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAAAAA"},
		{"InvalidFormat_ShouldReturnFalse", false, "8195247601-1"},
		{"InvalidFormat_ShouldReturnFalse", false, "337989413-53"},
		{"InvalidFormat_ShouldReturnFalse", false, "872 227 006 00"},
		{"InvalidFormat_ShouldReturnFalse", false, "459.911.677-05"},
		{"InvalidDigit_ShouldReturnFalse", false, "02102234243"},
		{"InvalidDigit_ShouldReturnFalse", false, "02102234142"},
		{"InvalidDigit_ShouldReturnFalse", false, "13798941353"},
		{"InvalidDigit_ShouldReturnFalse", false, "00676003001"},
		{"Valid_ShouldReturnTrue", true, "81952476011"},
		{"Valid_ShouldReturnTrue", true, "33798941353"},
		{"Valid_ShouldReturnTrue", true, "87222700600"},
		{"Valid_ShouldReturnTrue", true, "45991167705"},
		{"Valid_ShouldReturnTrue", true, "19595699996"},
		{"Valid_ShouldReturnTrue", true, "00067600300"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCNH(tc.v))
		})
	}
}
