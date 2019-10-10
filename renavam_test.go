package brdoc

import (
	"testing"
)

func TestIsRenavam(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Valid
		{"81285820857", true},
		{"17058862280", true},

		// Invalid
		{"17058862281", false},
		{"38919643060", false},
	} {
		t.Logf("#%d Renavam validation of %s should return %v: ", i, v.v, v.r)
		if IsRenavam(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
