package brdoc

import (
	"testing"
)

func TestIsPlate(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAA", false},

		{"ValidOldFormat", "AAA0000", true},
		{"ValidOldFormat", "ABC-1234", true},

		{"ValidNewFormat", "AAA0A00", true},
		{"ValidNewFormat", "ABC1D23", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsPlate(tc.doc))
		})
	}
}

func TestIsNationalPlate(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAA", false},

		{"InvalidNewFormat", "AAA0A00", false},
		{"InvalidNewFormat", "ABC1D23", false},

		{"ValidOldFormat", "AAA0000", true},
		{"ValidOldFormat", "ABC-1234", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsNationalPlate(tc.doc))
		})
	}
}

func TestIsMercosulPlate(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAA", false},

		{"InvalidOldFormat", "AAA0000", false},
		{"InvalidOldFormat", "ABC-1234", false},

		{"ValidNewFormat", "AAA0A00", true},
		{"ValidNewFormat", "ABC1D23", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsMercosulPlate(tc.doc))
		})
	}
}
