package brdoc

import "testing"

func TestIsVoterID(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"Invalid_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"Invalid_ShouldReturnFalse", false, ""},
		{"Invalid_ShouldReturnFalse", false, "AAAAAAAAAAA"},
		{"Invalid_ShouldReturnFalse", false, "915 5017 0193 0306"},
		{"Invalid_ShouldReturnFalse", false, "174 2241 7133 0004"},
		{"Invalid_ShouldReturnFalse", false, "259 7557 3388 0001"},
		{"Invalid_ShouldReturnFalse", false, "808-2536-1743-0486"},
		{"Invalid_ShouldReturnFalse", false, "9999 0236 0200 834"},
		{"Invalid_ShouldReturnFalse", false, "111122223333"},
		{"Invalid_ShouldReturnFalse", false, "098756718298"},
		{"Valid_ShouldReturnTrue", true, "381026662437"},
		{"Valid_ShouldReturnTrue", true, "048751641724"},
		{"Valid_ShouldReturnTrue", true, "285823622500"},
		{"Valid_ShouldReturnTrue", true, "115180722291"},
		{"Valid_ShouldReturnTrue", true, "455422250574"},
		{"Valid_ShouldReturnTrue", true, "122364350701"},
		{"Valid_ShouldReturnTrue", true, "103464042020"},
		{"Valid_ShouldReturnTrue", true, "737736771015"},
		{"Valid_ShouldReturnTrue", true, "222803251481"},
		{"Valid_ShouldReturnTrue", true, "877174621180"},
		{"Valid_ShouldReturnTrue", true, "837133461872"},
		{"Valid_ShouldReturnTrue", true, "686457161910"},
		{"Valid_ShouldReturnTrue", true, "300020430205"},
		{"Valid_ShouldReturnTrue", true, "704478161317"},
		{"Valid_ShouldReturnTrue", true, "875456481201"},
		{"Valid_ShouldReturnTrue", true, "513443030671"},
		{"Valid_ShouldReturnTrue", true, "816231560876"},
		{"Valid_ShouldReturnTrue", true, "211404870302"},
		{"Valid_ShouldReturnTrue", true, "557212661694"},
		{"Valid_ShouldReturnTrue", true, "353615220469"},
		{"Valid_ShouldReturnTrue", true, "350436562380"},
		{"Valid_ShouldReturnTrue", true, "684185562631"},
		{"Valid_ShouldReturnTrue", true, "035200670175"},
		{"Valid_ShouldReturnTrue", true, "468332130981"},
		{"Valid_ShouldReturnTrue", true, "034315432186"},
		{"Valid_ShouldReturnTrue", true, "426044362739"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsVoterID(tc.v))
		})
	}
}
