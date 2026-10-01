package brdoc

import (
	"testing"
)

func TestIsNFE(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsNFE, []docCase{
			{"empty", ""},
			{"letters", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalidCases(t, IsNFE, []docCase{
			{"too short", "5206043300991100250655012000000780026730161"},
			{"too long", "520604330099110025065501200000078002673016155"},
			{"hyphens", "52060-4330-0991-1002-5065-5012-0000-0078-0026-7301-615"},
			{"spaces in the wrong places", "5206 043300991100250655012000000780026730161 5"},
			{"letter out of the CNPJ", "35260712ABC34501DE35A50010000001231123456787"},
			{"lowercase letters", "35260712abc34501de35550010000001231123456787"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalidCases(t, IsNFE, []docCase{
			{"numeric CNPJ", "52060433009911002506550120000007800267301614"},
			{"alphanumeric CNPJ", "35260712ABC34501DE35550010000001231123456788"},
		})
	})

	// Calculated by hand, with valid check digits.
	t.Run("rejects invalid fields", func(t *testing.T) {
		assertInvalidCases(t, IsNFE, []docCase{
			{"UF 99", "99260733009911002506550010000001231123456787"},
			{"month 13", "35261333009911002506550010000001231123456788"},
			{"model 59", "35260733009911002506590010000001231123456784"},
		})
	})

	// Calculated by hand, with valid check digits.
	t.Run("rejects invalid emitters", func(t *testing.T) {
		assertInvalidCases(t, IsNFE, []docCase{
			{"invalid CNPJ", "35260733009911002507550010000001231123456781"},
			{"CNPJ in an NF-e series reserved for CPF", "35260733009911002506559200000001231123456785"},
			{"CPF in an NFC-e", "35260700012345678909650010000001231123456782"},
			{"CPF in an NFC-e, in a series reserved for CPF", "35260700012345678909659200000001231123456788"},
		})
	})

	t.Run("accepts the example of the official rules", func(t *testing.T) {
		assertValidCases(t, IsNFE, []docCase{
			{"without spaces", "52060433009911002506550120000007800267301615"},
			{"with spaces", "5206 0433 0099 1100 2506 5501 2000 0007 8002 6730 1615"},
		})
	})

	t.Run("accepts public access keys, from the Portal da Transparência", func(t *testing.T) {
		assertValid(t, IsNFE,
			"35260805289870000146550030001287001004948540",
			"33260801213619000147550010000695481329703460",
			"43260894868742000187550010004474901004797820",
			"31260810198974000347550060000428511794412064",
			"53260815441682000145550020000005171205003959",
		)
	})

	t.Run("accepts access keys calculated by hand", func(t *testing.T) {
		assertValidCases(t, IsNFE, []docCase{
			{"alphanumeric CNPJ", "35260712ABC34501DE35550010000001231123456787"},
			{"CT-e", "35260733009911002506570010000001231123456787"},
			{"CNPJ in an NFC-e", "35260733009911002506650010000001231123456782"},
			{"CPF in an NF-e series reserved for CPF", "35260700012345678909559200000001231123456785"},
			{"CPF in another NF-e series", "35260700012345678909550010000001231123456780"},
			{"CPF in a CT-e", "35260700012345678909570010000001231123456787"},
		})
	})
}
