package brdoc

import "testing"

func TestIsCEP(t *testing.T) {
	for i, v := range []struct {
		v string
		u []FederativeUnit
		r bool
	}{
		// Invalid format.
		{"3467875434578764345789654", []FederativeUnit{}, false},
		{"", []FederativeUnit{}, false},
		{"#$%¨&*(ABCDEF", []FederativeUnit{}, false},

		// Invalid Federative Unit.
		{"10000-000", []FederativeUnit{RJ}, false},
		{"25000-000", []FederativeUnit{ES}, false},
		{"29500-000", []FederativeUnit{MG}, false},
		{"35000-000", []FederativeUnit{BA}, false},
		{"45000-000", []FederativeUnit{SE}, false},
		{"49500-000", []FederativeUnit{PE}, false},
		{"53000-000", []FederativeUnit{AL}, false},
		{"57500-000", []FederativeUnit{PB}, false},
		{"58500-000", []FederativeUnit{RN}, false},
		{"59500-000", []FederativeUnit{CE}, false},
		{"62000-000", []FederativeUnit{PI}, false},
		{"64500-000", []FederativeUnit{MA}, false},
		{"65500-000", []FederativeUnit{PA}, false},
		{"67000-000", []FederativeUnit{AP}, false},
		{"68900-000", []FederativeUnit{AM}, false},
		{"69100-000", []FederativeUnit{RR}, false},
		{"69300-000", []FederativeUnit{AM}, false},
		{"69600-000", []FederativeUnit{AC}, false},
		{"69900-000", []FederativeUnit{DF}, false},
		{"71500-000", []FederativeUnit{GO}, false},
		{"72800-000", []FederativeUnit{DF}, false},
		{"73200-000", []FederativeUnit{GO}, false},
		{"74500-000", []FederativeUnit{RO}, false},
		{"76800-000", []FederativeUnit{TO}, false},
		{"77500-000", []FederativeUnit{MT}, false},
		{"78500-000", []FederativeUnit{MS}, false},
		{"79500-000", []FederativeUnit{PR}, false},
		{"84000-000", []FederativeUnit{SC}, false},
		{"89000-000", []FederativeUnit{RS}, false},
		{"95000-000", []FederativeUnit{SP}, false},

		// Valid.
		{"10000-000", []FederativeUnit{SP}, true},
		{"25000-000", []FederativeUnit{RJ}, true},
		{"29500-000", []FederativeUnit{ES}, true},
		{"35000-000", []FederativeUnit{MG}, true},
		{"45000-000", []FederativeUnit{BA}, true},
		{"49500-000", []FederativeUnit{SE}, true},
		{"53000-000", []FederativeUnit{PE}, true},
		{"57500-000", []FederativeUnit{AL}, true},
		{"58500-000", []FederativeUnit{PB}, true},
		{"59500-000", []FederativeUnit{RN}, true},
		{"62000-000", []FederativeUnit{CE}, true},
		{"64500-000", []FederativeUnit{PI}, true},
		{"65500-000", []FederativeUnit{MA}, true},
		{"67000-000", []FederativeUnit{PA}, true},
		{"68900-000", []FederativeUnit{AP}, true},
		{"69100-000", []FederativeUnit{AM}, true},
		{"69300-000", []FederativeUnit{RR}, true},
		{"69600-000", []FederativeUnit{AM}, true},
		{"69900-000", []FederativeUnit{AC}, true},
		{"71500-000", []FederativeUnit{DF}, true},
		{"72800-000", []FederativeUnit{GO}, true},
		{"73200-000", []FederativeUnit{DF}, true},
		{"74500-000", []FederativeUnit{GO}, true},
		{"76800-000", []FederativeUnit{RO}, true},
		{"77500-000", []FederativeUnit{TO}, true},
		{"78500-000", []FederativeUnit{MT}, true},
		{"79500-000", []FederativeUnit{MS}, true},
		{"84000-000", []FederativeUnit{PR}, true},
		{"89000-000", []FederativeUnit{SC}, true},
		{"95000-000", []FederativeUnit{RS}, true},

		// With different sizes of Federative Unit slices.
		{"10000-000", []FederativeUnit{SP, SP, SP}, true},
		{"25000-000", []FederativeUnit{}, true},
		{"00000-000", []FederativeUnit{}, true},
		{"00000-000", []FederativeUnit{SP}, false},
		{"29500-000", []FederativeUnit{MT, MS, MG}, false},
	} {
		t.Logf("#%d CEP validation of %s should return %v: ", i, v.v, v.r)
		if IsCEP(v.v, v.u...) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestValidateCEPFormat(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Invalid.
		{"3467875434578764345789654", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},
		{"0-0000000", false},
		{"00-000000", false},
		{"000-00000", false},
		{"0000-0000", false},
		{"000000-00", false},
		{"0000000-0", false},
		{"0000x000", false},
		{"0000x-000", false},
		{"0000000x", false},
		{"00000-00x", false},

		// Valid.
		{"00000000", true},
		{"00000-000", true},
		{"12345678", true},
		{"87654-321", true},
		{"99999999", true},
		{"99999-999", true},
	} {
		t.Logf("#%d CEP format validation of %s should return %v: ", i, v.v, v.r)
		if ValidateCEPFormat(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
