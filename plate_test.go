package brdoc

import (
	"testing"
)

func TestIsPlate(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Random data.
		{"3467875434578764345789654", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},

		// Valid old format.
		{"AAA0000", true},
		{"ABC-1234", true},

		// Valid new format.
		{"AAA0A00", true},
		{"ABC1D23", true},
	} {
		t.Logf("#%d plate validation of %s should return %v: ", i, v.v, v.r)
		if IsPlate(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestIsNationalPlate(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Random data.
		{"3467875434578764345789654", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},

		// Invalid new format.
		{"AAA0A00", false},
		{"ABC1D23", false},

		// Valid.
		{"AAA0000", true},
		{"ABC-1234", true},
	} {
		t.Logf("#%d national plate validation of %s should return %v: ", i, v.v, v.r)
		if IsNationalPlate(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestIsMercosulPlate(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Random data.
		{"3467875434578764345789654", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},

		// Invalid old format.
		{"AAA0000", false},
		{"ABC-1234", false},

		// Valid.
		{"AAA0A00", true},
		{"ABC1D23", true},
	} {
		t.Logf("#%d Mercosul plate validation of %s should return %v: ", i, v.v, v.r)
		if IsMercosulPlate(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
