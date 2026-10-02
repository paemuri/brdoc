package brdoc

import (
	"testing"
)

func TestIsCNH(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsCNH, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"letters", "AAAAAAAAAAA"},
		})
	})

	t.Run("rejects masks", func(t *testing.T) {
		assertInvalidCases(t, IsCNH, []docCase{
			{"hyphen before the last digit", "1234567890-0"},
			{"hyphen before the check digits", "246913578-28"},
			{"spaces", "493 827 156 37"},
			{"dots and hyphen", "617.283.945-64"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsCNH,
			"12345678901",
			"24691357829",
			"37037036700",
			"00012345601",
		)
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalid(t, IsCNH,
			"00000000000",
			"11111111111",
			"99999999999",
		)
	})

	// Calculated by hand, as real ones are personal data.
	t.Run("accepts valid documents", func(t *testing.T) {
		assertValidCases(t, IsCNH, []docCase{
			{"check digits 00", "12345678900"},
			{"check digits 28", "24691357828"},
			{"check digits 37", "49382715637"},
			{"check digits 64", "61728394564"},
			{"remainder 10 in the first check digit, second 7", "37037036707"},
			{"remainder 10 in the first check digit, second 2", "11111110102"},
			{"leading zeros", "00012345610"},
		})
	})
}
