package brdoc

import (
	"testing"
)

func TestIsCEP(t *testing.T) {
	for i, tc := range []struct {
		name     string
		doc      string
		valid    bool
		ufs      []UF
		validUFs bool
	}{
		{"InvalidData", "3467875434578764345789654", false, []UF{}, false},
		{"InvalidData", "", false, []UF{}, false},
		{"InvalidData", "AAAAAAAA", false, []UF{}, false},

		// First two digits can't be 00.
		{"InvalidInitialDigits", "00000-000", false, []UF{}, false},
		{"InvalidInitialDigits", "00801-000", false, []UF{}, false},

		{"InvalidFormat", "10000 000", false, []UF{}, false},
		{"InvalidFormat", "25000.000", false, []UF{}, false},
		{"InvalidFormat", "29500/000", false, []UF{}, false},
		{"InvalidFormat", "1-0000000", false, []UF{}, false},
		{"InvalidFormat", "10-000000", false, []UF{}, false},
		{"InvalidFormat", "100-00000", false, []UF{}, false},
		{"InvalidFormat", "1000-0000", false, []UF{}, false},
		{"InvalidFormat", "100000-00", false, []UF{}, false},
		{"InvalidFormat", "1000000-0", false, []UF{}, false},
		{"InvalidFormat", "1000x000", false, []UF{}, false},
		{"InvalidFormat", "1000x-000", false, []UF{}, false},
		{"InvalidFormat", "1000000x", false, []UF{}, false},
		{"InvalidFormat", "10000-00x", false, []UF{}, false},

		{"InvalidUF", "10000-000", true, []UF{RJ}, false},
		{"InvalidUF", "25000-000", true, []UF{ES}, false},
		{"InvalidUF", "29500-000", true, []UF{MG}, false},
		{"InvalidUF", "35000-000", true, []UF{BA}, false},
		{"InvalidUF", "45000-000", true, []UF{SE}, false},
		{"InvalidUF", "49500-000", true, []UF{PE}, false},
		{"InvalidUF", "53000-000", true, []UF{AL}, false},
		{"InvalidUF", "57500-000", true, []UF{PB}, false},
		{"InvalidUF", "58500-000", true, []UF{RN}, false},
		{"InvalidUF", "59500-000", true, []UF{CE}, false},
		{"InvalidUF", "62000-000", true, []UF{PI}, false},
		{"InvalidUF", "64500-000", true, []UF{MA}, false},
		{"InvalidUF", "65500-000", true, []UF{PA}, false},
		{"InvalidUF", "67000-000", true, []UF{AP}, false},
		{"InvalidUF", "68900-000", true, []UF{AM}, false},
		{"InvalidUF", "69100-000", true, []UF{RR}, false},
		{"InvalidUF", "69300-000", true, []UF{AM}, false},
		{"InvalidUF", "69600-000", true, []UF{AC}, false},
		{"InvalidUF", "69900-000", true, []UF{DF}, false},
		{"InvalidUF", "71500-000", true, []UF{GO}, false},
		{"InvalidUF", "72800-000", true, []UF{DF}, false},
		{"InvalidUF", "73200-000", true, []UF{GO}, false},
		{"InvalidUF", "74500-000", true, []UF{RO}, false},
		{"InvalidUF", "76800-000", true, []UF{TO}, false},
		{"InvalidUF", "77500-000", true, []UF{MT}, false},
		{"InvalidUF", "78500-000", true, []UF{MS}, false},
		{"InvalidUF", "79500-000", true, []UF{PR}, false},
		{"InvalidUF", "84000-000", true, []UF{SC}, false},
		{"InvalidUF", "89000-000", true, []UF{RS}, false},
		{"InvalidUF", "95000-000", true, []UF{SP}, false},
		{"InvalidUF", "29500-000", true, []UF{MT, MS, MG}, false},

		{"Valid", "10000-000", true, []UF{SP}, true},
		{"Valid", "25000-000", true, []UF{RJ}, true},
		{"Valid", "29500-000", true, []UF{ES}, true},
		{"Valid", "35000-000", true, []UF{MG}, true},
		{"Valid", "45000-000", true, []UF{BA}, true},
		{"Valid", "49500-000", true, []UF{SE}, true},
		{"Valid", "53000-000", true, []UF{PE}, true},
		{"Valid", "57500-000", true, []UF{AL}, true},
		{"Valid", "58500-000", true, []UF{PB}, true},
		{"Valid", "59500-000", true, []UF{RN}, true},
		{"Valid", "62000-000", true, []UF{CE}, true},
		{"Valid", "64500-000", true, []UF{PI}, true},
		{"Valid", "65500-000", true, []UF{MA}, true},
		{"Valid", "67000-000", true, []UF{PA}, true},
		{"Valid", "68900-000", true, []UF{AP}, true},
		{"Valid", "69100-000", true, []UF{AM}, true},
		{"Valid", "69300-000", true, []UF{RR}, true},
		{"Valid", "69600-000", true, []UF{AM}, true},
		{"Valid", "69900-000", true, []UF{AC}, true},
		{"Valid", "71500-000", true, []UF{DF}, true},
		{"Valid", "72800-000", true, []UF{GO}, true},
		{"Valid", "73200-000", true, []UF{DF}, true},
		{"Valid", "74500-000", true, []UF{GO}, true},
		{"Valid", "76800-000", true, []UF{RO}, true},
		{"Valid", "77500-000", true, []UF{TO}, true},
		{"Valid", "78500-000", true, []UF{MT}, true},
		{"Valid", "79500-000", true, []UF{MS}, true},
		{"Valid", "84000-000", true, []UF{PR}, true},
		{"Valid", "89000-000", true, []UF{SC}, true},
		{"Valid", "95000-000", true, []UF{RS}, true},
		// No dash, also valid.
		{"Valid", "10000000", true, []UF{SP}, true},
		{"Valid", "25000000", true, []UF{RJ}, true},
		{"Valid", "29500000", true, []UF{ES}, true},
		{"Valid", "35000000", true, []UF{MG}, true},
		{"Valid", "45000000", true, []UF{BA}, true},
		{"Valid", "49500000", true, []UF{SE}, true},
		{"Valid", "53000000", true, []UF{PE}, true},
		{"Valid", "57500000", true, []UF{AL}, true},
		{"Valid", "58500000", true, []UF{PB}, true},
		{"Valid", "59500000", true, []UF{RN}, true},
		{"Valid", "62000000", true, []UF{CE}, true},
		{"Valid", "64500000", true, []UF{PI}, true},
		{"Valid", "65500000", true, []UF{MA}, true},
		{"Valid", "67000000", true, []UF{PA}, true},
		{"Valid", "68900000", true, []UF{AP}, true},
		{"Valid", "69100000", true, []UF{AM}, true},
		{"Valid", "69300000", true, []UF{RR}, true},
		{"Valid", "69600000", true, []UF{AM}, true},
		{"Valid", "69900000", true, []UF{AC}, true},
		{"Valid", "71500000", true, []UF{DF}, true},
		{"Valid", "72800000", true, []UF{GO}, true},
		{"Valid", "73200000", true, []UF{DF}, true},
		{"Valid", "74500000", true, []UF{GO}, true},
		{"Valid", "76800000", true, []UF{RO}, true},
		{"Valid", "77500000", true, []UF{TO}, true},
		{"Valid", "78500000", true, []UF{MT}, true},
		{"Valid", "79500000", true, []UF{MS}, true},
		{"Valid", "84000000", true, []UF{PR}, true},
		{"Valid", "89000000", true, []UF{SC}, true},
		{"Valid", "95000000", true, []UF{RS}, true},
		// Validating multiple or no UFs.
		{"Valid", "10000-000", true, []UF{BA, SP, MG}, true},
		{"Valid", "25000-000", true, []UF{RJ, RJ, RJ}, true},
		{"Valid", "29500-000", true, []UF{}, true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			validFrom := IsCEPFrom(tc.doc, tc.ufs...)
			assertEq(t, tc.valid && tc.validUFs, validFrom)

			valid, uf := IsCEP(tc.doc)
			assertEq(t, tc.valid, valid)

			if tc.validUFs && len(tc.ufs) == 1 {
				assertEq(t, tc.ufs[0], uf)
			}
		})
	}
}
