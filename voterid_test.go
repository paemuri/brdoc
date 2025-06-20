package brdoc

import (
	"testing"
)

func TestIsVoterID(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAA", false},

		{"InvalidDigit", "111122223333", false},
		{"InvalidDigit", "098756718298", false},

		{"InvalidFormat", "915 5017 0193 0306", false},
		{"InvalidFormat", "174 2241 7133 0004", false},
		{"InvalidFormat", "259 7557 3388 0001", false},
		{"InvalidFormat", "808-2536-1743-0486", false},
		{"InvalidFormat", "9999 0236 0200 834", false},

		{"Valid", "381026662437", true},
		{"Valid", "048751641724", true},
		{"Valid", "285823622500", true},
		{"Valid", "115180722291", true},
		{"Valid", "455422250574", true},
		{"Valid", "122364350701", true},
		{"Valid", "103464042020", true},
		{"Valid", "737736771015", true},
		{"Valid", "222803251481", true},
		{"Valid", "877174621180", true},
		{"Valid", "837133461872", true},
		{"Valid", "686457161910", true},
		{"Valid", "300020430205", true},
		{"Valid", "704478161317", true},
		{"Valid", "875456481201", true},
		{"Valid", "513443030671", true},
		{"Valid", "816231560876", true},
		{"Valid", "211404870302", true},
		{"Valid", "557212661694", true},
		{"Valid", "353615220469", true},
		{"Valid", "350436562380", true},
		{"Valid", "684185562631", true},
		{"Valid", "035200670175", true},
		{"Valid", "468332130981", true},
		{"Valid", "034315432186", true},
		{"Valid", "426044362739", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsVoterID(tc.doc))
		})
	}
}
