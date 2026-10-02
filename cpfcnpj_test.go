package brdoc

import (
	"testing"
)

func TestIsCPF(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsCPF, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"letters", "AAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid masks", func(t *testing.T) {
		assertInvalidCases(t, IsCPF, []docCase{
			{"spaces", "123 456 789 09"},
			{"hyphens and dot swapped", "987-654-321.00"},
		})
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		assertInvalid(t, IsCPF,
			"123.456.789-90",
			"987.654.321-01",
		)
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalid(t, IsCPF,
			"000.000.000-00",
			"111.111.111-11",
			"222.222.222-22",
			"333.333.333-33",
			"444.444.444-44",
			"555.555.555-55",
			"666.666.666-66",
			"777.777.777-77",
			"888.888.888-88",
			"999.999.999-99",
		)
	})

	// Calculated by hand, as real ones are personal data.
	t.Run("accepts valid documents", func(t *testing.T) {
		t.Run("with mask", func(t *testing.T) {
			assertValid(t, IsCPF,
				"123.456.789-09",
				"987.654.321-00",
			)
		})

		t.Run("without mask", func(t *testing.T) {
			assertValidCases(t, IsCPF, []docCase{
				{"check digits 09", "12345678909"},
				{"check digits 00", "98765432100"},
				{"leading zeros", "00000019100"},
			})
		})
	})
}

func TestIsCNPJ(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalidCases(t, IsCNPJ, []docCase{
			{"too long", "3467875434578764345789654"},
			{"empty", ""},
			{"letters in the check digits", "AAAAAAAAAAAAAA"},
		})
	})

	t.Run("rejects invalid masks", func(t *testing.T) {
		assertInvalid(t, IsCNPJ,
			"26-637-142.0001/58",
			"74-221-325.0001/30",
		)
	})

	t.Run("rejects invalid check digits", func(t *testing.T) {
		t.Run("numeric", func(t *testing.T) {
			assertInvalid(t, IsCNPJ,
				"26.637.142/0001-85",
				"74.221.325/0001-03",
			)
		})

		t.Run("alphanumeric", func(t *testing.T) {
			assertInvalid(t, IsCNPJ,
				"WE.0PP.M5F/0001-92",
			)
		})
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalid(t, IsCNPJ,
			"00.000.000/0000-00",
			"11.111.111/1111-11",
			"22.222.222/2222-22",
			"33.333.333/3333-33",
			"44.444.444/4444-44",
			"55.555.555/5555-55",
			"66.666.666/6666-66",
			"77.777.777/7777-77",
			"88.888.888/8888-88",
			"99.999.999/9999-99",
		)
	})

	t.Run("rejects the order number 0000", func(t *testing.T) {
		assertInvalid(t, IsCNPJ,
			"26.637.142/0000-77",
		)
	})

	t.Run("accepts numeric documents", func(t *testing.T) {
		t.Run("with mask", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"26.637.142/0001-58",
				"74.221.325/0001-30",
			)
		})

		t.Run("without mask", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"26637142000158",
				"74221325000130",
			)
		})

		t.Run("with the base number 00.000.000", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"00.000.000/0001-91",
			)
		})

		t.Run("with base numbers listed as invalid, but issued", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"66.666.666/0001-91",
			)
		})

		t.Run("with order number over 0300 and base number starting with 000", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"00.360.305/0301-00",
			)
		})
	})

	t.Run("accepts alphanumeric documents", func(t *testing.T) {
		t.Run("with mask", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"19.JA2.KO8/Z001-51",
				"SC.RCN.1NI/0001-30",
				"4Y.2OP.G99/0001-41",
				"WE.0PP.M4F/0001-91",
				"12.ABC.345/01DE-35",
				"AA.AAA.AAA/0001-91",
			)
		})

		t.Run("without mask", func(t *testing.T) {
			assertValid(t, IsCNPJ,
				"19JA2KO8Z00151",
				"SCRCN1NI000130",
				"4Y2OPG99000141",
				"WE0PPM4F000191",
				"12ABC34501DE35",
			)
		})
	})
}
