package brdoc

import . "testing"

func Assert(t *T, value string, result, expected bool) {
	if result == expected {
		t.Logf("The result of %s should be \"%v\": ja!", value, expected)
	} else {
		t.Errorf("The result of %s should be \"%v\": nein!", value, expected)
	}
}
