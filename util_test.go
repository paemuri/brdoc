package brdoc

import (
	"testing"
)

// assertValid checks that `fn` returns true for every one of `docs`.
func assertValid(t *testing.T, fn func(string) bool, docs ...string) {
	for _, doc := range docs {
		if !fn(doc) {
			t.Errorf("expected %q to be valid", doc)
		}
	}
}

// assertInvalid checks that `fn` returns false for every one of `docs`.
func assertInvalid(t *testing.T, fn func(string) bool, docs ...string) {
	for _, doc := range docs {
		if fn(doc) {
			t.Errorf("expected %q to be invalid", doc)
		}
	}
}

// docCase is a named document, for tests with a subtest for each document.
type docCase struct {
	name string
	doc  string
}

// assertValidCases checks, in a subtest for each case, that `fn` returns true.
func assertValidCases(t *testing.T, fn func(string) bool, cases []docCase) {
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if !fn(c.doc) {
				t.Errorf("expected %q to be valid", c.doc)
			}
		})
	}
}

// assertInvalidCases checks, in a subtest for each case, that `fn` returns
// false.
func assertInvalidCases(t *testing.T, fn func(string) bool, cases []docCase) {
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if fn(c.doc) {
				t.Errorf("expected %q to be invalid", c.doc)
			}
		})
	}
}
