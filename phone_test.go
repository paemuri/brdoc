package brdoc

import (
	"fmt"
	"testing"
)

// isPhone reports if `phone` is valid by both `IsPhone` and `IsPhoneFrom`, to
// be used with `assertValid` and `assertInvalid`.
func isPhone(phone string) bool {
	valid, _ := IsPhone(phone)
	return valid || IsPhoneFrom(phone)
}

// assertPhoneUF checks that `phone` is valid and related to `uf`.
func assertPhoneUF(t *testing.T, phone string, uf UF) {
	valid, ufs := IsPhone(phone)
	if !valid {
		t.Errorf("expected %q to be valid", phone)
		return
	}

	found := false
	for _, u := range ufs {
		found = found || u == uf
	}
	if !found || !IsPhoneFrom(phone, uf) {
		t.Errorf("expected %q to be from %s, got %v", phone, uf, ufs)
	}
}

// assertPhoneNotUF checks that `phone` is valid, but not related to `uf`.
func assertPhoneNotUF(t *testing.T, phone string, uf UF) {
	if valid, _ := IsPhone(phone); !valid {
		t.Errorf("expected %q to be valid", phone)
	}
	if IsPhoneFrom(phone, uf) {
		t.Errorf("expected %q not to be from %s", phone, uf)
	}
}

func TestIsPhone(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalid(t, isPhone,
			"3467875434578764345789654",
			"",
			"AAAAAAAA",
		)
	})

	t.Run("rejects invalid country codes", func(t *testing.T) {
		assertInvalid(t, isPhone,
			"+0 1199999999",
			"+1 1199999999",
			"+5 1199999999",
			"+9 1199999999",
			"+51 1199999999",
			"+555 1199999999",
		)
	})

	t.Run("rejects invalid area codes", func(t *testing.T) {
		assertInvalid(t, isPhone,
			"0039999999",
			"0139999999",
			"0239999999",
			"0339999999",
			"0439999999",
			"0539999999",
			"0639999999",
			"0739999999",
			"0839999999",
			"0939999999",
			"1039999999",
			"2039999999",
			"2339999999",
			"2539999999",
			"2639999999",
			"2939999999",
			"3039999999",
			"3639999999",
			"3939999999",
			"4039999999",
			"5039999999",
			"5239999999",
			"5639999999",
			"5739999999",
			"5839999999",
			"5939999999",
			"6039999999",
			"7039999999",
			"7239999999",
			"7639999999",
			"7839999999",
			"8039999999",
			"9039999999",
		)
	})

	t.Run("rejects invalid numbers", func(t *testing.T) {
		assertInvalid(t, isPhone,
			"+55 11 999999",
			"+55 11 9999999",
			"+55 11 19999999",
			"+55 11 199999999",
			"+55 11 299999999",
			"+55 11 399999999",
			"+55 11 499999999",
			"+55 11 599999999",
			"+55 11 699999999",
			"+55 11 9999999999",
			"+55 11 99999999999",
			"1179999999",
			"1189999999",
			"1199999999",
			"(11) 9999-9999",
		)
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalid(t, isPhone,
			"+55 11 999 999 999",
			"+ 55 11 999999999",
			"+55 11 9999-99999",
			"+55 11 9999.99999",
			"+55 11 9999 99999",
			"+55 11 99999/9999",
			"+55 11 99999\\9999",
			"+55 (11 999999999",
			"+55 11) 999999999",
		)
	})

	t.Run("accepts all formats", func(t *testing.T) {
		assertPhoneUF(t, "+55 (11) 99999-9999", SP)
		assertPhoneUF(t, "+55(11)99999-9999", SP)
		assertPhoneUF(t, "+55(11)3999-9999", SP)
		assertPhoneUF(t, "55 (11) 99999-9999", SP)
		assertPhoneUF(t, "0055 11 999999999", SP)
		assertPhoneUF(t, "005511999999999", SP)
		assertPhoneUF(t, "00551139999999", SP)
		assertPhoneUF(t, "(11) 99999-9999", SP)
		assertPhoneUF(t, "(11) 99999.9999", SP)
		assertPhoneUF(t, "(11) 99999 9999", SP)
		assertPhoneUF(t, "(11) 9-9999-9999", SP)
		assertPhoneUF(t, "(11) 9-9999.9999", SP)
		assertPhoneUF(t, "(11) 9-9999 9999", SP)
		assertPhoneUF(t, "(11) 9.9999-9999", SP)
		assertPhoneUF(t, "(11) 9.9999.9999", SP)
		assertPhoneUF(t, "(11) 9.9999 9999", SP)
		assertPhoneUF(t, "(11) 9 9999-9999", SP)
		assertPhoneUF(t, "(11) 9 9999.9999", SP)
		assertPhoneUF(t, "(11) 9 9999 9999", SP)
		assertPhoneUF(t, "(11) 3999-9999", SP)
		assertPhoneUF(t, "(11) 3999.9999", SP)
		assertPhoneUF(t, "(11) 3999 9999", SP)
		assertPhoneUF(t, "11999999999", SP)
		assertPhoneUF(t, "11899999999", SP)
		assertPhoneUF(t, "11799999999", SP)
		assertPhoneUF(t, "+55 11 8-9999-9999", SP)
		assertPhoneUF(t, "1139999999", SP)
		assertPhoneUF(t, "1129999999", SP)
		assertPhoneUF(t, "1139999999", SP)
		assertPhoneUF(t, "1149999999", SP)
		assertPhoneUF(t, "1159999999", SP)
		assertPhoneUF(t, "1169999999", SP)
	})

	t.Run("returns the UFs of each area code", func(t *testing.T) {
		assertPhoneUF(t, "1139999999", SP)
		assertPhoneUF(t, "1239999999", SP)
		assertPhoneUF(t, "1339999999", SP)
		assertPhoneUF(t, "1439999999", SP)
		assertPhoneUF(t, "1539999999", SP)
		assertPhoneUF(t, "1639999999", SP)
		assertPhoneUF(t, "1739999999", SP)
		assertPhoneUF(t, "1839999999", SP)
		assertPhoneUF(t, "1939999999", SP)
		assertPhoneUF(t, "2139999999", RJ)
		assertPhoneUF(t, "2239999999", RJ)
		assertPhoneUF(t, "2439999999", RJ)
		assertPhoneUF(t, "2739999999", ES)
		assertPhoneUF(t, "2839999999", ES)
		assertPhoneUF(t, "3139999999", MG)
		assertPhoneUF(t, "3239999999", MG)
		assertPhoneUF(t, "3339999999", MG)
		assertPhoneUF(t, "3439999999", MG)
		assertPhoneUF(t, "3539999999", MG)
		assertPhoneUF(t, "3739999999", MG)
		assertPhoneUF(t, "3839999999", MG)
		assertPhoneUF(t, "4139999999", PR)
		assertPhoneUF(t, "4239999999", PR)
		assertPhoneUF(t, "4339999999", PR)
		assertPhoneUF(t, "4439999999", PR)
		assertPhoneUF(t, "4539999999", PR)
		assertPhoneUF(t, "4639999999", PR)
		assertPhoneUF(t, "4739999999", SC)
		assertPhoneUF(t, "4839999999", SC)
		assertPhoneUF(t, "4939999999", SC)
		assertPhoneUF(t, "5139999999", RS)
		assertPhoneUF(t, "5439999999", RS)
		assertPhoneUF(t, "5539999999", RS)
		assertPhoneUF(t, "6139999999", DF)
		assertPhoneUF(t, "6239999999", GO)
		assertPhoneUF(t, "6339999999", TO)
		assertPhoneUF(t, "6439999999", GO)
		assertPhoneUF(t, "6539999999", MT)
		assertPhoneUF(t, "6639999999", MT)
		assertPhoneUF(t, "6739999999", MS)
		assertPhoneUF(t, "6839999999", AC)
		assertPhoneUF(t, "6939999999", RO)
		assertPhoneUF(t, "7139999999", BA)
		assertPhoneUF(t, "7339999999", BA)
		assertPhoneUF(t, "7439999999", BA)
		assertPhoneUF(t, "7539999999", BA)
		assertPhoneUF(t, "7739999999", BA)
		assertPhoneUF(t, "7939999999", SE)
		assertPhoneUF(t, "8139999999", PE)
		assertPhoneUF(t, "8239999999", AL)
		assertPhoneUF(t, "8339999999", PB)
		assertPhoneUF(t, "8439999999", RN)
		assertPhoneUF(t, "8539999999", CE)
		assertPhoneUF(t, "8639999999", PI)
		assertPhoneUF(t, "8739999999", PE)
		assertPhoneUF(t, "8839999999", CE)
		assertPhoneUF(t, "8939999999", PI)
		assertPhoneUF(t, "9139999999", PA)
		assertPhoneUF(t, "9239999999", AM)
		assertPhoneUF(t, "9339999999", PA)
		assertPhoneUF(t, "9439999999", PA)
		assertPhoneUF(t, "9539999999", RR)
		assertPhoneUF(t, "9639999999", AP)
		assertPhoneUF(t, "9739999999", AM)
		assertPhoneUF(t, "9839999999", MA)
		assertPhoneUF(t, "9939999999", MA)
	})

	t.Run("rejects UFs that are not of the area code", func(t *testing.T) {
		assertPhoneNotUF(t, "1139999999", RJ)
		assertPhoneNotUF(t, "1239999999", RJ)
		assertPhoneNotUF(t, "1339999999", RJ)
		assertPhoneNotUF(t, "1439999999", RJ)
		assertPhoneNotUF(t, "1539999999", RJ)
		assertPhoneNotUF(t, "1639999999", RJ)
		assertPhoneNotUF(t, "1739999999", RJ)
		assertPhoneNotUF(t, "1839999999", RJ)
		assertPhoneNotUF(t, "1939999999", RJ)
		assertPhoneNotUF(t, "2139999999", ES)
		assertPhoneNotUF(t, "2239999999", ES)
		assertPhoneNotUF(t, "2439999999", ES)
		assertPhoneNotUF(t, "2739999999", MG)
		assertPhoneNotUF(t, "2839999999", MG)
		assertPhoneNotUF(t, "3139999999", PR)
		assertPhoneNotUF(t, "3239999999", PR)
		assertPhoneNotUF(t, "3339999999", PR)
		assertPhoneNotUF(t, "3439999999", PR)
		assertPhoneNotUF(t, "3539999999", PR)
		assertPhoneNotUF(t, "3739999999", PR)
		assertPhoneNotUF(t, "3839999999", PR)
		assertPhoneNotUF(t, "4139999999", SC)
		assertPhoneNotUF(t, "4239999999", RS)
		assertPhoneNotUF(t, "4339999999", SC)
		assertPhoneNotUF(t, "4439999999", SC)
		assertPhoneNotUF(t, "4539999999", SC)
		assertPhoneNotUF(t, "4639999999", SC)
		assertPhoneNotUF(t, "4739999999", RS)
		assertPhoneNotUF(t, "4839999999", RS)
		assertPhoneNotUF(t, "4939999999", RS)
		assertPhoneNotUF(t, "5139999999", DF)
		assertPhoneNotUF(t, "5439999999", DF)
		assertPhoneNotUF(t, "5539999999", DF)
		assertPhoneNotUF(t, "6139999999", MT)
		assertPhoneNotUF(t, "6239999999", TO)
		assertPhoneNotUF(t, "6339999999", MT)
		assertPhoneNotUF(t, "6439999999", TO)
		assertPhoneNotUF(t, "6539999999", MS)
		assertPhoneNotUF(t, "6639999999", MS)
		assertPhoneNotUF(t, "6739999999", AC)
		assertPhoneNotUF(t, "6839999999", RO)
		assertPhoneNotUF(t, "6939999999", BA)
		assertPhoneNotUF(t, "7139999999", SE)
		assertPhoneNotUF(t, "7339999999", SE)
		assertPhoneNotUF(t, "7439999999", SE)
		assertPhoneNotUF(t, "7539999999", SE)
		assertPhoneNotUF(t, "7739999999", SE)
		assertPhoneNotUF(t, "7939999999", PE)
		assertPhoneNotUF(t, "8139999999", AL)
		assertPhoneNotUF(t, "8239999999", PB)
		assertPhoneNotUF(t, "8339999999", RN)
		assertPhoneNotUF(t, "8439999999", CE)
		assertPhoneNotUF(t, "8539999999", PI)
		assertPhoneNotUF(t, "8639999999", PA)
		assertPhoneNotUF(t, "8739999999", AL)
		assertPhoneNotUF(t, "8839999999", PI)
		assertPhoneNotUF(t, "8939999999", PA)
		assertPhoneNotUF(t, "9139999999", AM)
		assertPhoneNotUF(t, "9239999999", PA)
		assertPhoneNotUF(t, "9339999999", RR)
		assertPhoneNotUF(t, "9439999999", RR)
		assertPhoneNotUF(t, "9539999999", AP)
		assertPhoneNotUF(t, "9639999999", AM)
		assertPhoneNotUF(t, "9739999999", MA)
		assertPhoneNotUF(t, "9839999999", SP)
		assertPhoneNotUF(t, "9939999999", SP)
	})

	t.Run("accepts a list of UFs", func(t *testing.T) {
		if !IsPhoneFrom("1139999999", BA, SP, MG) {
			t.Errorf("expected %q to be from %v", "1139999999", []UF{BA, SP, MG})
		}
		if !IsPhoneFrom("1239999999", SP, SP, SP) {
			t.Errorf("expected %q to be from %v", "1239999999", []UF{SP, SP, SP})
		}
		if !IsPhoneFrom("1339999999") {
			t.Errorf("expected %q to be from %v", "1339999999", []UF{})
		}
	})
}

func TestIsPhoneSharedDDD(t *testing.T) {
	t.Run("accepts all the UFs of the area code", func(t *testing.T) {
		assertPhoneUF(t, "4239999999", PR)
		assertPhoneUF(t, "4239999999", SC)
		assertPhoneUF(t, "4739999999", PR)
		assertPhoneUF(t, "4939999999", PR)
		assertPhoneUF(t, "6139999999", GO)
	})

	t.Run("rejects UFs that do not share the area code", func(t *testing.T) {
		assertPhoneNotUF(t, "4839999999", PR)
		assertPhoneNotUF(t, "6239999999", DF)
	})

	t.Run("returns all the UFs of the area code, in order", func(t *testing.T) {
		for _, tc := range []struct {
			phone string
			ufs   []UF
		}{
			{"4239999999", []UF{PR, SC}},
			{"4739999999", []UF{SC, PR}},
			{"4939999999", []UF{SC, PR}},
			{"6139999999", []UF{DF, GO}},
			{"4839999999", []UF{SC}},
		} {
			_, ufs := IsPhone(tc.phone)
			if fmt.Sprint(ufs) != fmt.Sprint(tc.ufs) {
				t.Errorf("expected %q to be from %v, got %v", tc.phone, tc.ufs, ufs)
			}
		}
	})

	t.Run("does not change the internal table", func(t *testing.T) {
		_, ufs := IsPhone("4239999999")
		ufs[0] = RS

		_, ufs = IsPhone("4239999999")
		if ufs[0] != PR {
			t.Errorf("expected the first UF to be %s, got %s", PR, ufs[0])
		}
	})
}
