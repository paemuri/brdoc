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
		{"03135375023", true},
		{"10016211275", true},
		{"38872054176", true},
		{"40999838201", true},
		{"31789431483", true},

		// Invalid
		{"38872054170", false},
		{"40999838209", false},
		{"31789431480", false},
		{"38919643060", false},
		{"123", false},
		{"abcdefghijk", false},
	} {
		t.Logf("#%d Renavam validation of %s should return %v: ", i, v.v, v.r)
		if IsRenavam(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
