package brdoc

import (
	"testing"
)

func TestIsRENAVAM(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAAAAA"},
		{"InvalidDigit_ShouldReturnFalse", false, "38872054170"},
		{"InvalidDigit_ShouldReturnFalse", false, "40999838209"},
		{"InvalidDigit_ShouldReturnFalse", false, "31789431480"},
		{"InvalidDigit_ShouldReturnFalse", false, "38919643060"},
		{"Valid_ShouldReturnTrue", true, "13824652268"},
		{"Valid_ShouldReturnTrue", true, "08543317523"},
		{"Valid_ShouldReturnTrue", true, "09769017014"},
		{"Valid_ShouldReturnTrue", true, "01993520012"},
		{"Valid_ShouldReturnTrue", true, "04598137389"},
		{"Valid_ShouldReturnTrue", true, "05204907510"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsRENAVAM(tc.v))
		})
	}
}
