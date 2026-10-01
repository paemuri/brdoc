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
		// With 16 digits.
		assertInvalid(t, IsVoterID,
			"915 5017 0193 0306",
			"174 2241 7133 0004",
			"259 7557 3388 0001",
		)

		assertInvalidCases(t, IsVoterID, []docCase{
			{"hyphens", "808-2536-1743-0486"},
			{"15 digits", "9999 0236 0200 834"},
			{"two spaces", "3810  2666 2437"},
			{"spaces in the wrong places", "38102 666 2437"},
			{"leading space", " 3810 2666 2437"},
			{"hyphens instead of spaces", "3810-2666-2437"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsVoterID,
			"111122223333",
			"098756718298",
		)
	})

	t.Run("accepts valid documents", func(t *testing.T) {
		t.Run("without spaces", func(t *testing.T) {
			assertValid(t, IsVoterID,
				"381026662437",
				"048751641724",
				"285823622500",
				"115180722291",
				"455422250574",
				"122364350701",
				"103464042020",
				"737736771015",
				"222803251481",
				"877174621180",
				"837133461872",
				"686457161910",
				"300020430205",
				"704478161317",
				"875456481201",
				"513443030671",
				"816231560876",
				"211404870302",
				"557212661694",
				"353615220469",
				"350436562380",
				"684185562631",
				"035200670175",
				"468332130981",
				"034315432186",
				"426044362739",
			)
		})

		t.Run("with spaces", func(t *testing.T) {
			assertValidCases(t, IsVoterID, []docCase{
				{"between all groups", "3810 2666 2437"},
				{"between all groups, with a leading zero", "0487 5164 1724"},
				{"between some groups", "28582362 2500"},
			})
		})
	})
}
