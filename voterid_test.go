package brdoc

import (
	"testing"
)

func TestIsVoterID(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsVoterID, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"letters", "AAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalidCases(t, IsVoterID, []docCase{
			{"16 digits", "135 7924 6010 8000"},
			{"hyphens", "1357-9246-0108"},
			{"15 digits", "1357 9246 0108 000"},
			{"two spaces", "1357  9246 0108"},
			{"spaces in the wrong places", "13579 246 0108"},
			{"leading space", " 1357 9246 0108"},
		})
	})

	// Calculated by hand, with valid check digits.
	t.Run("rejects invalid UFs", func(t *testing.T) {
		assertInvalidCases(t, IsVoterID, []docCase{
			{"UF 00", "135792460000"},
			{"UF 29", "135792462909"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsVoterID,
			"135792460109",
			"271584920215",
		)
	})

	// Calculated by hand, as real ones are personal data.
	t.Run("accepts valid documents", func(t *testing.T) {
		t.Run("of every UF", func(t *testing.T) {
			assertValidCases(t, IsVoterID, []docCase{
				{"UF 01", "135792460108"},
				{"UF 02", "271584920205"},
				{"UF 03", "407377380370"},
				{"UF 04", "543169840477"},
				{"UF 05", "678962300582"},
				{"UF 06", "814754760671"},
				{"UF 07", "950547220701"},
				{"UF 08", "086339680809"},
				{"UF 09", "222132140906"},
				{"UF 10", "357924601066"},
				{"UF 11", "493717061147"},
				{"UF 12", "629509521228"},
				{"UF 13", "765301981392"},
				{"UF 14", "901094441473"},
				{"UF 15", "036886901546"},
				{"UF 16", "172679361627"},
				{"UF 17", "308471821724"},
				{"UF 18", "444264281805"},
				{"UF 19", "580056741902"},
				{"UF 20", "715849202046"},
				{"UF 21", "851641662100"},
				{"UF 22", "987434122283"},
				{"UF 23", "123226582399"},
				{"UF 24", "259019042461"},
				{"UF 25", "394811502500"},
				{"UF 26", "530603962690"},
				{"UF 27", "666396422763"},
				{"UF 28", "802188882852"},
			})
		})

		t.Run("with remainder 10 in the check digits", func(t *testing.T) {
			assertValidCases(t, IsVoterID, []docCase{
				{"second check digit", "407377380370"},
				{"both check digits", "394811502500"},
			})
		})

		t.Run("with spaces", func(t *testing.T) {
			assertValidCases(t, IsVoterID, []docCase{
				{"between all groups", "1357 9246 0108"},
				{"between all groups, with a leading zero", "0863 3968 0809"},
				{"between some groups", "27158492 0205"},
			})
		})
	})
}
