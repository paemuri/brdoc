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
		{"Invalid_ShouldReturnFalse", false, "38872054170"},
		{"Invalid_ShouldReturnFalse", false, "40999838209"},
		{"Invalid_ShouldReturnFalse", false, "31789431480"},
		{"Invalid_ShouldReturnFalse", false, "38919643060"},
		{"Invalid_ShouldReturnFalse", false, "123"},
		{"Invalid_ShouldReturnFalse", false, "abcdefghijk"},
		{"Valid_ShouldReturnTrue", true, "81285820857"},
		{"Valid_ShouldReturnTrue", true, "03135375023"},
		{"Valid_ShouldReturnTrue", true, "10016211275"},
		{"Valid_ShouldReturnTrue", true, "38872054176"},
		{"Valid_ShouldReturnTrue", true, "40999838201"},
		{"Valid_ShouldReturnTrue", true, "31789431483"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsRENAVAM(tc.v))
		})
	}
}
