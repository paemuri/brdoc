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
			{"CPF mask", "103.951.990-15"},
			{"spaces", "103 951 990 15"},
			{"hyphen in the wrong place", "103951-99015"},
		})

		// Groups other than 3, 5 and 2 digits.
		assertInvalid(t, IsPIS,
			"103.951.9901-5",
			"120.6372.482-4",
			"120.1641.414-8",
		)
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsPIS,
			"103.95199.01-6",
			"120.16414.14-9",
		)
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalidCases(t, IsPIS, []docCase{
			{"without mask", "00000000000"},
			{"with mask", "000.00000.00-0"},
		})
	})

	t.Run("accepts valid documents", func(t *testing.T) {
		t.Run("with mask", func(t *testing.T) {
			assertValid(t, IsPIS,
				"103.95199.01-5",
				"120.63724.82-4",
				"120.16414.14-8",
			)
		})

		t.Run("without mask", func(t *testing.T) {
			assertValid(t, IsPIS,
				"10395199015",
				"12016414148",
			)
		})
	})
}
