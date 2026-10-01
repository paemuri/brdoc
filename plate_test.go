package brdoc

import (
	"testing"
)

func TestIsPlate(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsPlate, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"only letters", "AAAAAAA"},
		})
	})

	t.Run("accepts the old national format", func(t *testing.T) {
		assertValidCases(t, IsPlate, []docCase{
			{"without hyphen", "AAA0000"},
			{"with hyphen", "ABC-1234"},
		})
	})

	t.Run("accepts the Mercosul format", func(t *testing.T) {
		assertValid(t, IsPlate,
			"AAA0A00",
			"ABC1D23",
		)
	})
}

func TestIsNationalPlate(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsNationalPlate, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"only letters", "AAAAAAA"},
		})
	})

	t.Run("rejects the Mercosul format", func(t *testing.T) {
		assertInvalid(t, IsNationalPlate,
			"AAA0A00",
			"ABC1D23",
		)
	})

	t.Run("accepts the old national format", func(t *testing.T) {
		assertValidCases(t, IsNationalPlate, []docCase{
			{"without hyphen", "AAA0000"},
			{"with hyphen", "ABC-1234"},
		})
	})
}

func TestIsMercosulPlate(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsMercosulPlate, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"only letters", "AAAAAAA"},
		})
	})

	t.Run("rejects the old national format", func(t *testing.T) {
		assertInvalidCases(t, IsMercosulPlate, []docCase{
			{"without hyphen", "AAA0000"},
			{"with hyphen", "ABC-1234"},
		})
	})

	t.Run("accepts the Mercosul format", func(t *testing.T) {
		assertValid(t, IsMercosulPlate,
			"AAA0A00",
			"ABC1D23",
		)
	})
}
