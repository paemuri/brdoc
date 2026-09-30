package brdoc

import (
	"fmt"
	"testing"
)

func TestIsPhone(t *testing.T) {
	for i, tc := range []struct {
		name     string
		phone    string
		valid    bool
		ufs      []UF
		validUFs bool
	}{
		{"InvalidData", "3467875434578764345789654", false, []UF{}, false},
		{"InvalidData", "", false, []UF{}, false},
		{"InvalidData", "AAAAAAAA", false, []UF{}, false},

		{"InvalidDDI", "+0 1199999999", false, []UF{}, false},
		{"InvalidDDI", "+1 1199999999", false, []UF{}, false},
		{"InvalidDDI", "+5 1199999999", false, []UF{}, false},
		{"InvalidDDI", "+9 1199999999", false, []UF{}, false},
		{"InvalidDDI", "+51 1199999999", false, []UF{}, false},
		{"InvalidDDI", "+555 1199999999", false, []UF{}, false},

		{"InvalidDDD", "0039999999", false, []UF{}, false},
		{"InvalidDDD", "0139999999", false, []UF{}, false},
		{"InvalidDDD", "0239999999", false, []UF{}, false},
		{"InvalidDDD", "0339999999", false, []UF{}, false},
		{"InvalidDDD", "0439999999", false, []UF{}, false},
		{"InvalidDDD", "0539999999", false, []UF{}, false},
		{"InvalidDDD", "0639999999", false, []UF{}, false},
		{"InvalidDDD", "0739999999", false, []UF{}, false},
		{"InvalidDDD", "0839999999", false, []UF{}, false},
		{"InvalidDDD", "0939999999", false, []UF{}, false},
		{"InvalidDDD", "1039999999", false, []UF{}, false},
		{"InvalidDDD", "2039999999", false, []UF{}, false},
		{"InvalidDDD", "2339999999", false, []UF{}, false},
		{"InvalidDDD", "2539999999", false, []UF{}, false},
		{"InvalidDDD", "2639999999", false, []UF{}, false},
		{"InvalidDDD", "2939999999", false, []UF{}, false},
		{"InvalidDDD", "3039999999", false, []UF{}, false},
		{"InvalidDDD", "3639999999", false, []UF{}, false},
		{"InvalidDDD", "3939999999", false, []UF{}, false},
		{"InvalidDDD", "4039999999", false, []UF{}, false},
		{"InvalidDDD", "5039999999", false, []UF{}, false},
		{"InvalidDDD", "5239999999", false, []UF{}, false},
		{"InvalidDDD", "5639999999", false, []UF{}, false},
		{"InvalidDDD", "5739999999", false, []UF{}, false},
		{"InvalidDDD", "5839999999", false, []UF{}, false},
		{"InvalidDDD", "5939999999", false, []UF{}, false},
		{"InvalidDDD", "6039999999", false, []UF{}, false},
		{"InvalidDDD", "7039999999", false, []UF{}, false},
		{"InvalidDDD", "7239999999", false, []UF{}, false},
		{"InvalidDDD", "7639999999", false, []UF{}, false},
		{"InvalidDDD", "7839999999", false, []UF{}, false},
		{"InvalidDDD", "8039999999", false, []UF{}, false},
		{"InvalidDDD", "9039999999", false, []UF{}, false},

		{"InvalidNumber", "+55 11 999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 9999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 19999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 199999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 299999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 399999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 499999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 599999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 699999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 9999999999", false, []UF{}, false},
		{"InvalidNumber", "+55 11 99999999999", false, []UF{}, false},
		{"InvalidNumber", "1179999999", false, []UF{}, false},
		{"InvalidNumber", "1189999999", false, []UF{}, false},
		{"InvalidNumber", "1199999999", false, []UF{}, false},
		{"InvalidNumber", "(11) 9999-9999", false, []UF{}, false},

		{"InvalidFormat", "+55 11 999 999 999", false, []UF{}, false},
		{"InvalidFormat", "+ 55 11 999999999", false, []UF{}, false},
		{"InvalidFormat", "+55 11 9999-99999", false, []UF{}, false},
		{"InvalidFormat", "+55 11 9999.99999", false, []UF{}, false},
		{"InvalidFormat", "+55 11 9999 99999", false, []UF{}, false},
		{"InvalidFormat", "+55 11 99999/9999", false, []UF{}, false},
		{"InvalidFormat", "+55 11 99999\\9999", false, []UF{}, false},
		{"InvalidFormat", "+55 (11 999999999", false, []UF{}, false},
		{"InvalidFormat", "+55 11) 999999999", false, []UF{}, false},

		{"InvalidUF", "1139999999", true, []UF{RJ}, false},
		{"InvalidUF", "1239999999", true, []UF{RJ}, false},
		{"InvalidUF", "1339999999", true, []UF{RJ}, false},
		{"InvalidUF", "1439999999", true, []UF{RJ}, false},
		{"InvalidUF", "1539999999", true, []UF{RJ}, false},
		{"InvalidUF", "1639999999", true, []UF{RJ}, false},
		{"InvalidUF", "1739999999", true, []UF{RJ}, false},
		{"InvalidUF", "1839999999", true, []UF{RJ}, false},
		{"InvalidUF", "1939999999", true, []UF{RJ}, false},
		{"InvalidUF", "2139999999", true, []UF{ES}, false},
		{"InvalidUF", "2239999999", true, []UF{ES}, false},
		{"InvalidUF", "2439999999", true, []UF{ES}, false},
		{"InvalidUF", "2739999999", true, []UF{MG}, false},
		{"InvalidUF", "2839999999", true, []UF{MG}, false},
		{"InvalidUF", "3139999999", true, []UF{PR}, false},
		{"InvalidUF", "3239999999", true, []UF{PR}, false},
		{"InvalidUF", "3339999999", true, []UF{PR}, false},
		{"InvalidUF", "3439999999", true, []UF{PR}, false},
		{"InvalidUF", "3539999999", true, []UF{PR}, false},
		{"InvalidUF", "3739999999", true, []UF{PR}, false},
		{"InvalidUF", "3839999999", true, []UF{PR}, false},
		{"InvalidUF", "4139999999", true, []UF{SC}, false},
		{"InvalidUF", "4239999999", true, []UF{RS}, false},
		{"InvalidUF", "4339999999", true, []UF{SC}, false},
		{"InvalidUF", "4439999999", true, []UF{SC}, false},
		{"InvalidUF", "4539999999", true, []UF{SC}, false},
		{"InvalidUF", "4639999999", true, []UF{SC}, false},
		{"InvalidUF", "4739999999", true, []UF{RS}, false},
		{"InvalidUF", "4839999999", true, []UF{RS}, false},
		{"InvalidUF", "4939999999", true, []UF{RS}, false},
		{"InvalidUF", "5139999999", true, []UF{DF}, false},
		{"InvalidUF", "5439999999", true, []UF{DF}, false},
		{"InvalidUF", "5539999999", true, []UF{DF}, false},
		{"InvalidUF", "6139999999", true, []UF{MT}, false},
		{"InvalidUF", "6239999999", true, []UF{TO}, false},
		{"InvalidUF", "6339999999", true, []UF{MT}, false},
		{"InvalidUF", "6439999999", true, []UF{TO}, false},
		{"InvalidUF", "6539999999", true, []UF{MS}, false},
		{"InvalidUF", "6639999999", true, []UF{MS}, false},
		{"InvalidUF", "6739999999", true, []UF{AC}, false},
		{"InvalidUF", "6839999999", true, []UF{RO}, false},
		{"InvalidUF", "6939999999", true, []UF{BA}, false},
		{"InvalidUF", "7139999999", true, []UF{SE}, false},
		{"InvalidUF", "7339999999", true, []UF{SE}, false},
		{"InvalidUF", "7439999999", true, []UF{SE}, false},
		{"InvalidUF", "7539999999", true, []UF{SE}, false},
		{"InvalidUF", "7739999999", true, []UF{SE}, false},
		{"InvalidUF", "7939999999", true, []UF{PE}, false},
		{"InvalidUF", "8139999999", true, []UF{AL}, false},
		{"InvalidUF", "8239999999", true, []UF{PB}, false},
		{"InvalidUF", "8339999999", true, []UF{RN}, false},
		{"InvalidUF", "8439999999", true, []UF{CE}, false},
		{"InvalidUF", "8539999999", true, []UF{PI}, false},
		{"InvalidUF", "8639999999", true, []UF{PA}, false},
		{"InvalidUF", "8739999999", true, []UF{AL}, false},
		{"InvalidUF", "8839999999", true, []UF{PI}, false},
		{"InvalidUF", "8939999999", true, []UF{PA}, false},
		{"InvalidUF", "9139999999", true, []UF{AM}, false},
		{"InvalidUF", "9239999999", true, []UF{PA}, false},
		{"InvalidUF", "9339999999", true, []UF{RR}, false},
		{"InvalidUF", "9439999999", true, []UF{RR}, false},
		{"InvalidUF", "9539999999", true, []UF{AP}, false},
		{"InvalidUF", "9639999999", true, []UF{AM}, false},
		{"InvalidUF", "9739999999", true, []UF{MA}, false},
		{"InvalidUF", "9839999999", true, []UF{SP}, false},
		{"InvalidUF", "9939999999", true, []UF{SP}, false},

		// All possible formats.
		{"Valid", "+55 (11) 99999-9999", true, []UF{SP}, true},
		{"Valid", "+55(11)99999-9999", true, []UF{SP}, true},
		{"Valid", "+55(11)3999-9999", true, []UF{SP}, true},
		{"Valid", "55 (11) 99999-9999", true, []UF{SP}, true},
		{"Valid", "0055 11 999999999", true, []UF{SP}, true},
		{"Valid", "005511999999999", true, []UF{SP}, true},
		{"Valid", "00551139999999", true, []UF{SP}, true},
		{"Valid", "(11) 99999-9999", true, []UF{SP}, true},
		{"Valid", "(11) 99999.9999", true, []UF{SP}, true},
		{"Valid", "(11) 99999 9999", true, []UF{SP}, true},
		{"Valid", "(11) 9-9999-9999", true, []UF{SP}, true},
		{"Valid", "(11) 9-9999.9999", true, []UF{SP}, true},
		{"Valid", "(11) 9-9999 9999", true, []UF{SP}, true},
		{"Valid", "(11) 9.9999-9999", true, []UF{SP}, true},
		{"Valid", "(11) 9.9999.9999", true, []UF{SP}, true},
		{"Valid", "(11) 9.9999 9999", true, []UF{SP}, true},
		{"Valid", "(11) 9 9999-9999", true, []UF{SP}, true},
		{"Valid", "(11) 9 9999.9999", true, []UF{SP}, true},
		{"Valid", "(11) 9 9999 9999", true, []UF{SP}, true},
		{"Valid", "(11) 3999-9999", true, []UF{SP}, true},
		{"Valid", "(11) 3999.9999", true, []UF{SP}, true},
		{"Valid", "(11) 3999 9999", true, []UF{SP}, true},
		{"Valid", "11999999999", true, []UF{SP}, true},
		{"Valid", "11899999999", true, []UF{SP}, true},
		{"Valid", "11799999999", true, []UF{SP}, true},
		{"Valid", "+55 11 8-9999-9999", true, []UF{SP}, true},
		{"Valid", "1139999999", true, []UF{SP}, true},
		{"Valid", "1129999999", true, []UF{SP}, true},
		{"Valid", "1139999999", true, []UF{SP}, true},
		{"Valid", "1149999999", true, []UF{SP}, true},
		{"Valid", "1159999999", true, []UF{SP}, true},
		{"Valid", "1169999999", true, []UF{SP}, true},
		// All possible DDD digits.
		{"Valid", "1139999999", true, []UF{SP}, true},
		{"Valid", "1239999999", true, []UF{SP}, true},
		{"Valid", "1339999999", true, []UF{SP}, true},
		{"Valid", "1439999999", true, []UF{SP}, true},
		{"Valid", "1539999999", true, []UF{SP}, true},
		{"Valid", "1639999999", true, []UF{SP}, true},
		{"Valid", "1739999999", true, []UF{SP}, true},
		{"Valid", "1839999999", true, []UF{SP}, true},
		{"Valid", "1939999999", true, []UF{SP}, true},
		{"Valid", "2139999999", true, []UF{RJ}, true},
		{"Valid", "2239999999", true, []UF{RJ}, true},
		{"Valid", "2439999999", true, []UF{RJ}, true},
		{"Valid", "2739999999", true, []UF{ES}, true},
		{"Valid", "2839999999", true, []UF{ES}, true},
		{"Valid", "3139999999", true, []UF{MG}, true},
		{"Valid", "3239999999", true, []UF{MG}, true},
		{"Valid", "3339999999", true, []UF{MG}, true},
		{"Valid", "3439999999", true, []UF{MG}, true},
		{"Valid", "3539999999", true, []UF{MG}, true},
		{"Valid", "3739999999", true, []UF{MG}, true},
		{"Valid", "3839999999", true, []UF{MG}, true},
		{"Valid", "4139999999", true, []UF{PR}, true},
		{"Valid", "4239999999", true, []UF{PR}, true},
		{"Valid", "4339999999", true, []UF{PR}, true},
		{"Valid", "4439999999", true, []UF{PR}, true},
		{"Valid", "4539999999", true, []UF{PR}, true},
		{"Valid", "4639999999", true, []UF{PR}, true},
		{"Valid", "4739999999", true, []UF{SC}, true},
		{"Valid", "4839999999", true, []UF{SC}, true},
		{"Valid", "4939999999", true, []UF{SC}, true},
		{"Valid", "5139999999", true, []UF{RS}, true},
		{"Valid", "5439999999", true, []UF{RS}, true},
		{"Valid", "5539999999", true, []UF{RS}, true},
		{"Valid", "6139999999", true, []UF{DF}, true},
		{"Valid", "6239999999", true, []UF{GO}, true},
		{"Valid", "6339999999", true, []UF{TO}, true},
		{"Valid", "6439999999", true, []UF{GO}, true},
		{"Valid", "6539999999", true, []UF{MT}, true},
		{"Valid", "6639999999", true, []UF{MT}, true},
		{"Valid", "6739999999", true, []UF{MS}, true},
		{"Valid", "6839999999", true, []UF{AC}, true},
		{"Valid", "6939999999", true, []UF{RO}, true},
		{"Valid", "7139999999", true, []UF{BA}, true},
		{"Valid", "7339999999", true, []UF{BA}, true},
		{"Valid", "7439999999", true, []UF{BA}, true},
		{"Valid", "7539999999", true, []UF{BA}, true},
		{"Valid", "7739999999", true, []UF{BA}, true},
		{"Valid", "7939999999", true, []UF{SE}, true},
		{"Valid", "8139999999", true, []UF{PE}, true},
		{"Valid", "8239999999", true, []UF{AL}, true},
		{"Valid", "8339999999", true, []UF{PB}, true},
		{"Valid", "8439999999", true, []UF{RN}, true},
		{"Valid", "8539999999", true, []UF{CE}, true},
		{"Valid", "8639999999", true, []UF{PI}, true},
		{"Valid", "8739999999", true, []UF{PE}, true},
		{"Valid", "8839999999", true, []UF{CE}, true},
		{"Valid", "8939999999", true, []UF{PI}, true},
		{"Valid", "9139999999", true, []UF{PA}, true},
		{"Valid", "9239999999", true, []UF{AM}, true},
		{"Valid", "9339999999", true, []UF{PA}, true},
		{"Valid", "9439999999", true, []UF{PA}, true},
		{"Valid", "9539999999", true, []UF{RR}, true},
		{"Valid", "9639999999", true, []UF{AP}, true},
		{"Valid", "9739999999", true, []UF{AM}, true},
		{"Valid", "9839999999", true, []UF{MA}, true},
		{"Valid", "9939999999", true, []UF{MA}, true},
		// Validating multiple or no UFs.
		{"Valid", "1139999999", true, []UF{BA, SP, MG}, true},
		{"Valid", "1239999999", true, []UF{SP, SP, SP}, true},
		{"Valid", "1339999999", true, []UF{}, true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			validFrom := IsPhoneFrom(tc.phone, tc.ufs...)
			assertEq(t, tc.valid && tc.validUFs, validFrom)

			valid, ufs := IsPhone(tc.phone)
			assertEq(t, tc.valid, valid)

			if tc.validUFs && len(tc.ufs) == 1 {
				found := false
				for _, uf := range ufs {
					found = found || uf == tc.ufs[0]
				}
				assertEq(t, true, found)
			}
		})
	}
}

func TestIsPhoneSharedDDD(t *testing.T) {
	for i, tc := range []struct {
		name  string
		phone string
		ufs   []UF
		valid bool
	}{
		{"MainUF", "4239999999", []UF{PR}, true},
		{"SharedUF", "4239999999", []UF{SC}, true},
		{"SharedUF", "4739999999", []UF{PR}, true},
		{"SharedUF", "4939999999", []UF{PR}, true},
		{"SharedUF", "6139999999", []UF{GO}, true},
		{"NotSharedUF", "4839999999", []UF{PR}, false},
		{"NotSharedUF", "6239999999", []UF{DF}, false},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsPhoneFrom(tc.phone, tc.ufs...))
		})
	}

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
		assertEq(t, fmt.Sprint(tc.ufs), fmt.Sprint(ufs))
	}

	_, ufs := IsPhone("4239999999")
	ufs[0] = RS
	_, ufs = IsPhone("4239999999")
	assertEq(t, PR, ufs[0])
}
