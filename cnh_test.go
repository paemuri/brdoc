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
			{"hyphen before the last digit", "8195247601-1"},
			{"hyphen before the check digits", "337989413-53"},
			{"spaces", "872 227 006 00"},
			{"dots and hyphen", "459.911.677-05"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsCNH,
			"02102234243",
			"02102234142",
			"13798941353",
			"00676003001",
		)
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalid(t, IsCNH,
			"00000000000",
			"11111111111",
			"99999999999",
		)
	})

	t.Run("accepts valid documents", func(t *testing.T) {
		assertValid(t, IsCNH,
			"81952476011",
			"33798941353",
			"87222700600",
			"45991167705",
			"19595699996",
			"00067600300",
		)
	})
}
