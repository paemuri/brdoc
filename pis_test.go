package brdoc

import (
	"testing"
)

func TestIsPIS(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsPIS, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"letters", "AAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid masks", func(t *testing.T) {
		assertInvalidCases(t, IsPIS, []docCase{
			{"CPF mask", "123.456.789-19"},
			{"spaces", "123 456 789 19"},
			{"hyphen in the wrong place", "123456-78919"},
		})

		// Groups other than 3, 5 and 2 digits.
		assertInvalid(t, IsPIS,
			"123.456.7891-9",
			"246.9135.782-7",
			"370.3703.673-3",
		)
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsPIS,
			"123.45678.91-0",
			"246.91357.82-8",
		)
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalidCases(t, IsPIS, []docCase{
			{"without mask", "00000000000"},
			{"with mask", "000.00000.00-0"},
		})
	})

	// Calculated by hand, as real ones are personal data.
	t.Run("accepts valid documents", func(t *testing.T) {
		t.Run("with mask", func(t *testing.T) {
			assertValid(t, IsPIS,
				"123.45678.91-9",
				"246.91357.82-7",
				"370.37036.73-3",
			)
		})

		t.Run("without mask", func(t *testing.T) {
			assertValidCases(t, IsPIS, []docCase{
				{"check digit 9", "12345678919"},
				{"check digit 7", "24691357827"},
				{"check digit 0", "98765431280"},
			})
		})
	})
}
