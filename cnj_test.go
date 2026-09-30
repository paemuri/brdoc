package brdoc

import (
	"testing"
)

func TestIsCNJ(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAA-AA.AAAA.A.AA.AAAA", false},

		{"InvalidFormat", "1001051-87.2025.4.01.320", false},
		{"InvalidFormat", "1001051-87.2025.4.01.32011", false},
		{"InvalidFormat", "1001051.87-2025.4.01.3201", false},
		{"InvalidFormat", "100-15.2008.2.00", false},

		{"InvalidDigit", "1001051-88.2025.4.01.3201", false},
		{"InvalidDigit", "1001051-87.2025.4.01.3202", false},
		// Fictitious examples from the official rules, with invalid digits.
		{"InvalidDigit", "0000100-15.2008.4.01.0000", false},

		// Calculated by hand, with a valid digit but a segment (J) of 0.
		{"InvalidSegment", "0000100-69.2008.0.00.0000", false},

		// Calculated by hand, with a valid digit but a court (TR) that does not
		// exist in the segment (J).
		{"InvalidCourt", "0000100-41.2008.4.07.0000", false},
		{"InvalidCourt", "0000100-89.2008.9.01.0000", false},
		{"InvalidCourt", "0000100-03.2008.8.00.0000", false},

		// Public numbers of processes.
		{"Valid", "1001051-87.2025.4.01.3201", true},
		{"Valid", "10010518720254013201", true},
		{"Valid", "1002234-26.2026.4.01.3600", true},
		{"Valid", "1042229-16.2025.4.01.0000", true},
		{"Valid", "1000438-90.2026.5.02.0042", true},
		{"Valid", "1000371-65.2022.5.02.0463", true},
		{"Valid", "1000190-75.2026.5.02.0705", true},
		{"Valid", "0606999-77.2026.6.26.0000", true},
		{"Valid", "0607000-62.2026.6.26.0000", true},
		{"Valid", "0607004-02.2026.6.26.0000", true},

		// Calculated by hand.
		{"Valid", "0000100-68.2008.8.26.0100", true},
		{"Valid", "0001234-37.2020.5.02.0001", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsCNJ(tc.doc))
		})
	}
}
