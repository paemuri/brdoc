package brdoc

import (
	"testing"
)

func TestIsCNO(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsCNO, []docCase{
			{"empty", ""},
			{"letters", "AAAAAAAAAAAA"},
			{"all digits equal", "111111111111"},
		})
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalidCases(t, IsCNO, []docCase{
			{"too short", "90001234567"},
			{"too long", "9000123456781"},
			{"hyphen in place of slash", "90.001.23456-78"},
			{"dots in the wrong places", "900.012.3456/78"},
		})
	})

	// Calculated by hand, with valid check digits.
	t.Run("accepts valid documents", func(t *testing.T) {
		assertValidCases(t, IsCNO, []docCase{
			{"without mask", "900012345678"},
			{"with mask", "90.001.23456/78"},
			{"check digit 0", "123456789010"},
			{"leading zeros", "000000000016"},
		})
	})

	// Works of companies, from the open data of the Receita Federal.
	t.Run("accepts real documents", func(t *testing.T) {
		assertValidCases(t, IsCNO, []docCase{
			{"registered as CNO", "900072928963"},
			{"registered as CEI", "351400466371"},
			{"registered as CEI, with check digit 0", "010010144970"},
			{"registered as CEI, with activity 7", "010010092278"},
			{"registered as CEI, with mask", "13.076.07698/76"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsCNO,
			"900012345670",
			"90.001.23456/79",
			"123456789011",
		)
	})
}
