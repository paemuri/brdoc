package brdoc

import (
	"testing"
)

func TestIsCivilCertificate(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsCivilCertificate, []docCase{
			{"empty", ""},
			{"letters", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalidCases(t, IsCivilCertificate, []docCase{
			{"too short", "1045390155201310001202100001232"},
			{"too long", "104539015520131000120210000123211"},
			{"slash before the check digits", "104539.01.55.2013.1.00012.021.0000123/21"},
		})
	})

	t.Run("rejects invalid fields", func(t *testing.T) {
		assertInvalidCases(t, IsCivilCertificate, []docCase{
			{"service 56 instead of 55", "10453901562013100012021000012321"},
			{"type of book 8", "10453901552013800012021000012321"},
			{"type of book 0", "10453901552013000012021000012321"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalidCases(t, IsCivilCertificate, []docCase{
			{"changed second digit", "10453901552013100012021000012322"},
			{"changed first digit", "10453901552013100012021000012311"},
			{"remainder 10 as the digit 0 instead of 1", "10453901552013100012020000000304"},
		})
	})

	// Calculated by hand, as real ones are personal data.
	t.Run("accepts valid documents", func(t *testing.T) {
		assertValidCases(t, IsCivilCertificate, []docCase{
			{"without mask", "10453901552013100012021000012321"},
			{"with dots and hyphen", "104539.01.55.2013.1.00012.021.0000123-21"},
			{"with spaces", "104539 01 55 2013 1 00012 021 0000123 21"},
			{"with remainder 10", "10453901552013100012020000000314"},
			{"of a marriage", "00000001552024200001001000000145"},
		})
	})
}
