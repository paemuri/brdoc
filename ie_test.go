package brdoc

import (
	"testing"
)

func TestIsIE(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		uf    UF
		valid bool
	}{
		{"NotImplementedUF", "110.042.490.114", RJ, false},

		{"SP_InvalidData", "", SP, false},
		{"SP_InvalidData", "AAAAAAAAAAAA", SP, false},

		{"SP_InvalidFormat", "11004249011", SP, false},
		{"SP_InvalidFormat", "1100424901145", SP, false},
		{"SP_InvalidFormat", "110-042-490-114", SP, false},
		{"SP_InvalidFormat", "p-01100424.3/002", SP, false},
		{"SP_InvalidFormat", "P-0110042.43/002", SP, false},
		{"SP_InvalidFormat", "P-01100424.3/0021", SP, false},

		{"SP_InvalidDigit", "110.042.491.114", SP, false},
		{"SP_InvalidDigit", "110.042.490.115", SP, false},
		{"SP_InvalidDigit", "P-01100424.4/002", SP, false},

		// Examples from the official rules.
		{"SP_Valid", "110.042.490.114", SP, true},
		{"SP_Valid", "110042490114", SP, true},
		{"SP_Valid", "P-01100424.3/002", SP, true},
		{"SP_Valid", "P011004243002", SP, true},
		{"SP_Valid", "P-01100424.3/999", SP, true},

		// Public IEs of companies.
		{"SP_Valid", "310.035.324.119", SP, true},
		{"SP_Valid", "108.354.656.114", SP, true},
		{"SP_Valid", "142.270.790.110", SP, true},
		{"SP_Valid", "142.484.958.110", SP, true},
		{"SP_Valid", "102.654.009.110", SP, true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsIE(tc.doc, tc.uf))
		})
	}
}
