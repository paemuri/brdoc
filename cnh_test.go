package brdoc

import (
	"testing"
)

func TestIsCNH(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		{"", false},
		{"02102234243", false},
		{"02102234142", false},
		{"33798941353", true},
		{"13798941353", false},
		{"81952476011", true},
		{"819524760111", false},
	} {
		t.Logf("#%d CNH validation of %s should return %v: ", i, v.v, v.r)
		if IsCNH(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
