package brdoc

import (
	"testing"
)

func TestIsNFE(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", false},

		{"InvalidFormat", "5206043300991100250655012000000780026730161", false},
		{"InvalidFormat", "520604330099110025065501200000078002673016155", false},
		{"InvalidFormat", "52060-4330-0991-1002-5065-5012-0000-0078-0026-7301-615", false},
		{"InvalidFormat", "5206 043300991100250655012000000780026730161 5", false},
		// Letters are only allowed in the first 12 positions of the CNPJ.
		{"InvalidFormat", "35260712ABC34501DE35A50010000001231123456787", false},
		{"InvalidFormat", "35260712abc34501de35550010000001231123456787", false},

		{"InvalidDigit", "52060433009911002506550120000007800267301614", false},
		{"InvalidDigit", "35260712ABC34501DE35550010000001231123456788", false},

		// Calculated by hand, with valid digits but invalid fields.
		{"InvalidUF", "99260733009911002506550010000001231123456787", false},
		{"InvalidMonth", "35261333009911002506550010000001231123456788", false},
		{"InvalidModel", "35260733009911002506590010000001231123456784", false},

		// Calculated by hand, with valid digits but invalid emitters.
		{"InvalidEmitter", "35260733009911002507550010000001231123456781", false},
		{"InvalidEmitter", "35260700012345678909550010000001231123456780", false},
		{"InvalidEmitter", "35260733009911002506559200000001231123456785", false},
		{"InvalidEmitter", "35260700012345678909650010000001231123456782", false},
		{"InvalidEmitter", "35260700012345678909659200000001231123456788", false},

		// Example from the official rules.
		{"Valid", "52060433009911002506550120000007800267301615", true},
		{"Valid", "5206 0433 0099 1100 2506 5501 2000 0007 8002 6730 1615", true},

		// Calculated by hand, with an alphanumeric CNPJ and with a CT-e.
		{"Valid", "35260712ABC34501DE35550010000001231123456787", true},
		{"Valid", "35260733009911002506570010000001231123456787", true},
		// Calculated by hand, with a CNPJ in an NFC-e, and a CPF in an NF-e and
		// in a CT-e.
		{"Valid", "35260733009911002506650010000001231123456782", true},
		{"Valid", "35260700012345678909559200000001231123456785", true},
		{"Valid", "35260700012345678909570010000001231123456787", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsNFE(tc.doc))
		})
	}
}
