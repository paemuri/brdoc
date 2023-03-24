package brdoc

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected, actual bool) {
	if expected != actual {
		t.Fatalf(`
Not equal!
expected: %v
actual:   %v
`, expected, actual)
	}
}

func assertIntEqual(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatalf(`
Not equal!
expected: %v
actual:   %v
`, expected, actual)
	}
}

func assertStringEqual(t *testing.T, expected, actual string) {
	if expected != actual {
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
