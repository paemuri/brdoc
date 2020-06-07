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
		{"Invalid_ShouldReturnFalse", false, ""},
		{"Invalid_ShouldReturnFalse", false, "02102234243"},
		{"Invalid_ShouldReturnFalse", false, "02102234142"},
		{"Invalid_ShouldReturnFalse", false, "13798941353"},
		{"Invalid_ShouldReturnFalse", false, "819524760111"},
		{"Valid_ShouldReturnTrue", true, "81952476011"},
		{"Valid_ShouldReturnTrue", true, "33798941353"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCNH(tc.v))
		})
	}
}
