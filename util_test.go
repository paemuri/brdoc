package brdoc

import (
	"fmt"
	"testing"
)

func assertEq(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf(`
Not equal!
expected: %v
actual:   %v
`, expected, actual)
	}
}

func assertNotEq(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		t.Fatalf(`
Not equal!
expected: %v
actual:   %v
`, expected, actual)
	}
}

func testName(i int, name string) string {
	return fmt.Sprintf("%d_%s", i, name)
}
