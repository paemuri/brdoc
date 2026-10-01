package brdoc

import (
	"testing"
)

func TestIsCNJ(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsCNJ, []docCase{
			{"empty", ""},
			{"letters", "AAAAAAA-AA.AAAA.A.AA.AAAA"},
		})
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalidCases(t, IsCNJ, []docCase{
			{"too short", "1001051-87.2025.4.01.320"},
			{"too long", "1001051-87.2025.4.01.32011"},
			{"separators swapped", "1001051.87-2025.4.01.3201"},
			{"leading zeros omitted", "100-15.2008.2.00"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalidCases(t, IsCNJ, []docCase{
			{"changed check digits", "1001051-88.2025.4.01.3201"},
			{"changed origin", "1001051-87.2025.4.01.3202"},
			{"fictitious example of the official rules", "0000100-15.2008.4.01.0000"},
		})
	})

	// Calculated by hand, with valid check digits.
	t.Run("rejects invalid segments and courts", func(t *testing.T) {
		assertInvalidCases(t, IsCNJ, []docCase{
			{"segment (J) 0", "0000100-69.2008.0.00.0000"},
			{"court (TR) 07 in the federal justice", "0000100-41.2008.4.07.0000"},
			{"court (TR) 01 in the state military justice", "0000100-89.2008.9.01.0000"},
			{"court (TR) 00 in the state justice", "0000100-03.2008.8.00.0000"},
		})
	})

	t.Run("accepts public numbers of processes", func(t *testing.T) {
		t.Run("of the federal justice", func(t *testing.T) {
			assertValidCases(t, IsCNJ, []docCase{
				{"with mask", "1001051-87.2025.4.01.3201"},
				{"without mask", "10010518720254013201"},
				{"with another origin", "1002234-26.2026.4.01.3600"},
				{"originated in the court", "1042229-16.2025.4.01.0000"},
			})
		})

		t.Run("of the labor justice", func(t *testing.T) {
			assertValid(t, IsCNJ,
				"1000438-90.2026.5.02.0042",
				"1000371-65.2022.5.02.0463",
				"1000190-75.2026.5.02.0705",
			)
		})

		t.Run("of the electoral justice", func(t *testing.T) {
			assertValid(t, IsCNJ,
				"0606999-77.2026.6.26.0000",
				"0607000-62.2026.6.26.0000",
				"0607004-02.2026.6.26.0000",
			)
		})
	})

	t.Run("accepts numbers calculated by hand", func(t *testing.T) {
		assertValidCases(t, IsCNJ, []docCase{
			{"state justice", "0000100-68.2008.8.26.0100"},
			{"labor justice", "0001234-37.2020.5.02.0001"},
		})
	})
}
