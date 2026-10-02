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
			"27182818287",
			"54365636560",
			"81548454849",
			"61118911841",
		)
	})

	// Calculated by hand, as real ones are personal data.
	t.Run("accepts valid documents", func(t *testing.T) {
		assertValidCases(t, IsRENAVAM, []docCase{
			{"check digit 6", "27182818286"},
			{"check digit 7", "54365636567"},
			{"check digit 8", "81548454848"},
			{"leading zero", "08731273120"},
			{"check digit 1", "35914091401"},
			{"remainder 10", "61118911840"},
		})
	})
}
