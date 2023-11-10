package brdoc

import (
	"testing"
)

func TestIsValidTituloEleitoral(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "123456789011"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "000000000000"},
		{"InvalidData_ShouldReturnFalse", false, "123456789012345"},
		{"InvalidPattern_ShouldReturnFalse", false, "1234.5678.9012"},
		{"InvalidPattern_ShouldReturnFalse", false, "1234 5678 9012"},
		{"InvalidPattern_ShouldReturnFalse", false, "12345678-9012"},
		{"InvalidDigits_ShouldReturnFalse", false, "1234.5678901-6"},
		{"InvalidDigits_ShouldReturnFalse", false, "120.1641.414-9"},
		{"InvalidDigits_ShouldReturnFalse", false, "12345678901"},
		{"InvalidDigits_ShouldReturnFalse", false, "1234567890123"},
		{"InvalidDigits_ShouldReturnFalse", false, "ABCD56789012"},
		{"InvalidDigits_ShouldReturnFalse", false, "217633 460 930"},
		{"Valid_ShouldReturnTrue", true, "217633460930"},
		{"Valid_ShouldReturnTrue", true, "3244567800167"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, isValidTituloEleitoral(tc.v))
		})
	}
}