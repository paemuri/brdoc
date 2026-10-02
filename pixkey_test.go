package brdoc

import (
	"testing"
)

func TestIsPixKey(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsPixKey, []docCase{
			{"empty", ""},
			{"letters", "AAAAAAAAAAAA"},
		})
	})

	t.Run("accepts every type of key", func(t *testing.T) {
		assertValidCases(t, IsPixKey, []docCase{
			{"CPF", "12345678909"},
			{"numeric CNPJ", "00038166000105"},
			{"alphanumeric CNPJ", "12ABC34501DE35"},
			{"phone", "+5561912345678"},
			{"email", "fulano_da_silva.recebedor@example.com"},
			{"email with 77 characters", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com"},
			{"random key", "123e4567-e12b-12d1-a456-426655440000"},
		})
	})

	t.Run("rejects invalid CPFs and CNPJs", func(t *testing.T) {
		assertInvalidCases(t, IsPixKey, []docCase{
			{"CPF with mask", "123.456.789-09"},
			{"CPF with invalid check digits", "12345678900"},
			{"CPF with all digits equal", "11111111111"},
			{"CNPJ with mask", "00.038.166/0001-05"},
			{"CNPJ with invalid check digits", "00038166000106"},
			{"CNPJ in lowercase", "12abc34501de35"},
		})
	})

	t.Run("rejects invalid phones", func(t *testing.T) {
		assertInvalidCases(t, IsPixKey, []docCase{
			{"without plus sign", "5561912345678"},
			{"with mask", "+55 (61) 91234-5678"},
			{"starting with 0", "+0561912345678"},
			{"with 16 digits", "+5561912345678901"},
		})
	})

	t.Run("rejects invalid emails", func(t *testing.T) {
		assertInvalidCases(t, IsPixKey, []docCase{
			{"in uppercase", "Fulano@example.com"},
			{"with 78 characters", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com"},
			{"without at sign", "fulano.example.com"},
			{"domain starting with hyphen", "fulano@-example.com"},
		})
	})

	t.Run("rejects invalid random keys", func(t *testing.T) {
		assertInvalidCases(t, IsPixKey, []docCase{
			{"in uppercase", "123E4567-E12B-12D1-A456-426655440000"},
			{"without hyphens", "123e4567e12b12d1a456426655440000"},
			{"with braces", "{123e4567-e12b-12d1-a456-426655440000}"},
		})
	})
}
