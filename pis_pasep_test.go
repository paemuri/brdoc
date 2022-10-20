package brdoc

import (
	"testing"
)

func TestIsPIS(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "11111111111"},
		{"InvalidData_ShouldReturnFalse", false, "103951990150"},
		{"InvalidData_ShouldReturnFalse", false, "103.951.9901-0"},
		{"Valid_ShouldReturnTrue", true, "103.951.9901-5"},
		{"Valid_ShouldReturnTrue", true, "103.95199.01-5"},
		{"Valid_ShouldReturnTrue", true, "120.6372.482-4"},
		{"Valid_ShouldReturnTrue", true, "120.1641.414-8"},
		{"Valid_ShouldReturnTrue", true, "10395199015"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsPIS(tc.v))
		})
	}
}
