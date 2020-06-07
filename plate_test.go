package brdoc

import (
	"testing"
)

func TestIsPlate(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "#$%¨&*(ABCDEF"},
		{"ValidOldFormat_ShouldReturnTrue", true, "AAA0000"},
		{"ValidOldFormat_ShouldReturnTrue", true, "ABC-1234"},
		{"ValidNewFormat_ShouldReturnTrue", true, "AAA0A00"},
		{"ValidNewFormat_ShouldReturnTrue", true, "ABC1D23"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsPlate(tc.v))
		})
	}
}

func TestIsNationalPlate(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "#$%¨&*(ABCDEF"},
		{"InvalidNewFormat_ShouldReturnFlase", false, "AAA0A00"},
		{"InvalidNewFormat_ShouldReturnFlase", false, "ABC1D23"},
		{"ValidOldFormat_ShouldReturnTrue", true, "AAA0000"},
		{"ValidOldFormat_ShouldReturnTrue", true, "ABC-1234"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsNationalPlate(tc.v))
		})
	}
}

func TestIsMercosulPlate(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "#$%¨&*(ABCDEF"},
		{"InvalidOldFormat_ShouldReturnFlase", false, "AAA0000"},
		{"InvalidOldFormat_ShouldReturnFlase", false, "ABC-1234"},
		{"ValidNewFormat_ShouldReturnTrue", true, "AAA0A00"},
		{"ValidNewFormat_ShouldReturnTrue", true, "ABC1D23"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsMercosulPlate(tc.v))
		})
	}
}
