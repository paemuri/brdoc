package brdoc

import (
	"testing"
)

func TestIsBankSlip(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsBankSlip, []docCase{
			{"empty", ""},
			{"letters", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalidCases(t, IsBankSlip, []docCase{
			{"barcode too short", "1049432420000032112005507722213334777777777"},
			{"barcode too long", "104943242000003211200550772221333477777777711"},
			{"barcode with spaces", "10494 32420000032112 0055077222133347777777771"},
			{"line too short", "10490.05505 77222.133348 77777.777713 4 3242000003211"},
			{"line with hyphens", "10490-05505 77222-133348 77777-777713 4 32420000032112"},
			{"line with 2 spaces", "10490.05505  77222.133348  77777.777713  4  32420000032112"},
			{"revenue line too short", "82620000215-0 04820097412-8 32201540982-2 9010860594-8"},
			{"revenue line with dots", "82620000215.0 04820097412.8 32201540982.2 90108605940.8"},
		})
	})

	t.Run("accepts the slips of banks", func(t *testing.T) {
		// Example of the Caixa specification for its slips.
		t.Run("barcode", func(t *testing.T) {
			assertValid(t, IsBankSlip,
				"10494324200000321120055077222133347777777771",
			)
		})

		t.Run("line", func(t *testing.T) {
			assertValidCases(t, IsBankSlip, []docCase{
				{"with mask", "10490.05505 77222.133348 77777.777713 4 32420000032112"},
				{"without mask", "10490055057722213334877777777713432420000032112"},
				{"without dots", "1049005505 77222133348 77777777713 4 32420000032112"},
			})
		})

		// Calculated by hand, with valid check digits.
		t.Run("with the code 988", func(t *testing.T) {
			assertValid(t, IsBankSlip,
				"98802000000123456781234567890123456789012345",
				"98801.23451 67890.123457 67890.123457 2 00000012345678",
			)
		})
	})

	t.Run("rejects invalid slips of banks", func(t *testing.T) {
		assertInvalidCases(t, IsBankSlip, []docCase{
			{"general check digit", "10495324200000321120055077222133347777777771"},
			{"general check digit 0", "10490324200000321120055077222133347777777771"},
			{"check digit of the field 1", "10490.05506 77222.133348 77777.777713 4 32420000032112"},
			{"check digit of the field 2", "10490.05505 77222.133349 77777.777713 4 32420000032112"},
			{"check digit of the field 3", "10490.05505 77222.133348 77777.777714 4 32420000032112"},
			{"general check digit in the line", "10490.05505 77222.133348 77777.777713 5 32420000032112"},
		})

		// Calculated by hand, with valid check digits.
		assertInvalidCases(t, IsBankSlip, []docCase{
			{"currency 5", "10457324200000321120055077222133347777777771"},
			{"code 988 with currency 9", "98899000000123456781234567890123456789012345"},
			{"code 988 with due date factor", "98807324200123456781234567890123456789012345"},
			{"line starting with 8", "80490.05500 77222.133348 77777.777713 9 32420000032112"},
		})
	})

	// Calculated by hand from the example of the FEBRABAN layout, with valid
	// check digits.
	t.Run("accepts revenue slips", func(t *testing.T) {
		t.Run("with modulo 10", func(t *testing.T) {
			assertValidCases(t, IsBankSlip, []docCase{
				{"barcode", "82620000215048200974123220154098290108605940"},
				{"line with mask", "82620000215-0 04820097412-8 32201540982-2 90108605940-8"},
				{"line without mask", "826200002150048200974128322015409822901086059408"},
				{"line without hyphens", "826200002150 048200974128 322015409822 901086059408"},
				{"identifier 7", "82700000215048200974123220154098290108605940"},
				{"segment 9", "89650000215048200974123220154098290108605940"},
			})
		})

		t.Run("with modulo 11", func(t *testing.T) {
			assertValidCases(t, IsBankSlip, []docCase{
				{"barcode", "82890000215048200974123220154098290108605940"},
				{"line", "82890000215-9 04820097412-7 32201540982-1 90108605940-3"},
				{"identifier 9", "82970000215048200974123220154098290108605940"},
			})
		})
	})

	t.Run("rejects invalid revenue slips", func(t *testing.T) {
		assertInvalidCases(t, IsBankSlip, []docCase{
			{"check digit", "82630000215048200974123220154098290108605940"},
			{"check digit of the block 1", "82620000215-1 04820097412-8 32201540982-2 90108605940-8"},
			{"check digit of the block 4", "82620000215-0 04820097412-8 32201540982-2 90108605940-9"},
		})

		// Calculated by hand, with valid check digits.
		assertInvalidCases(t, IsBankSlip, []docCase{
			{"segment 8", "88660000215048200974123220154098290108605940"},
			{"segment 0", "80640000215048200974123220154098290108605940"},
			{"identifier 2, of the example", "82210000215048200974123220154098290108605940"},
			{"line not starting with 8", "72640000215-0 04820097412-8 32201540982-2 90108605940-8"},
		})
	})
}
