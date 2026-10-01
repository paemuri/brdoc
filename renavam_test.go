package brdoc

import (
	"testing"
)

func TestIsRENAVAM(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsRENAVAM, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"letters", "AAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsRENAVAM,
			"38872054170",
			"40999838209",
			"31789431480",
			"38919643060",
		)
	})

	t.Run("accepts valid documents", func(t *testing.T) {
		assertValid(t, IsRENAVAM,
			"13824652268",
			"08543317523",
			"09769017014",
			"01993520012",
			"04598137389",
			"05204907510",
		)
	})
}
