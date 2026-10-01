package brdoc

import (
	"testing"
)

// ieOf returns `IsIE` for the given `uf`, to be used with `assertValid` and
// `assertInvalid`.
func ieOf(uf UF) func(string) bool {
	return func(doc string) bool {
		return IsIE(doc, uf)
	}
}

func TestIsIE(t *testing.T) {
	t.Run("rejects an invalid UF", func(t *testing.T) {
		assertInvalid(t, ieOf(UF("XX")),
			"290887780",
		)
	})

	t.Run("rejects ISENTO, used by those exempt from IE", func(t *testing.T) {
		assertInvalid(t, ieOf(SP),
			"ISENTO",
		)
	})

	t.Run("rejects documents with all digits equal", func(t *testing.T) {
		assertInvalid(t, ieOf(ES), "000000000")
		assertInvalid(t, ieOf(SC), "000000000")
		assertInvalid(t, ieOf(TO), "000000000")
		assertInvalid(t, ieOf(RJ), "00000000")
		assertInvalid(t, ieOf(DF), "00000000000-00")
		assertInvalid(t, ieOf(SP), "000000000000")
	})

	t.Run("AC", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalidCases(t, ieOf(AC), []docCase{
				{"empty", ""},
				{"letters", "AAAAAAAAAAAAA"},
			})
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(AC), []docCase{
				{"too short", "010048230011"},
				{"too long", "01004823001123"},
				{"prefix 02 instead of 01", "02.004.823/001-12"},
				{"separators swapped", "01-004-823.001/12"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(AC),
				"01.004.823/001-02",
				"01.004.823/001-13",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(AC),
				"01.004.823/001-12",
				"0100482300112",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(AC),
				"01.008.267/001-08",
			)
		})
	})

	t.Run("AL", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalidCases(t, ieOf(AL), []docCase{
				{"empty", ""},
				{"letters", "AAAAAAAAA"},
			})
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(AL), []docCase{
				{"too short", "24000004"},
				{"too long", "2400000480"},
				{"prefix 25 instead of 24", "250000048"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(AL),
				"240000049",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(AL),
				"240000048",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(AL),
				"245008403",
			)
		})

		t.Run("accepts types of company not listed by older rules", func(t *testing.T) {
			assertValid(t, ieOf(AL),
				"241000009",
			)
		})
	})

	t.Run("AM", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(AM),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(AM), []docCase{
				{"without the check digit", "04.900.976"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(AM),
				"04.900.976-2",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(AM),
				"04.900.976-1",
				"04.150.272-8",
			)
		})
	})

	t.Run("AP", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(AP),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(AP), []docCase{
				{"too short", "03012345"},
				{"prefix 04 instead of 03", "040123459"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(AP),
				"030123458",
				"030170070",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(AP),
				"030123459",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(AP),
				"030071009",
			)
		})

		t.Run("accepts IEs calculated by hand, one for each range and for results 10 and 11", func(t *testing.T) {
			assertValid(t, ieOf(AP),
				"030170011",
				"030170020",
				"030170071",
				"030190231",
				"030190240",
				"030190290",
			)
		})
	})

	t.Run("BA", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(BA),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(BA), []docCase{
				{"too short", "12345-63"},
				{"too long", "12345678-90"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(BA),
				"123456-64",
				"123456-53",
				"612345-58",
				"1000003-07",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(BA),
				"123456-63",
				"612345-57",
				"1000003-06",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(BA),
				"0004786-96",
				"0010273-89",
			)
		})

		t.Run("accepts IEs calculated by hand, with 9 digits and modulo 11", func(t *testing.T) {
			assertValid(t, ieOf(BA),
				"1600003-00",
				"1700005-80",
			)
		})
	})

	t.Run("CE", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(CE),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(CE), []docCase{
				{"too short", "0600001-5"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(CE),
				"06000001-4",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(CE),
				"06000001-5",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(CE),
				"06863259-2",
				"06305950-9",
			)
		})
	})

	t.Run("DF", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(DF),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(DF), []docCase{
				{"hyphen at the end", "0730000100109-"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(DF),
				"07300001001-08",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(DF),
				"07300001001-09",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(DF),
				"07656443030-76",
				"07679960002-46",
			)
		})
	})

	t.Run("ES", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(ES),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(ES), []docCase{
				{"too short", "99999999"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(ES),
				"999999991",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(ES),
				"999999990",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(ES),
				"000002976",
				"000017744",
			)
		})
	})

	t.Run("GO", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(GO),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(GO), []docCase{
				{"prefix 12, not 10, 11 or from 20 to 29", "12.987.654-7"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(GO),
				"10.987.654-8",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(GO),
				"10.987.654-7",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(GO),
				"10.277.380-7",
				"10.353.194-7",
			)
		})
	})

	t.Run("MA", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(MA),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(MA), []docCase{
				{"prefix 13 instead of 12", "130000385"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(MA),
				"120000386",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(MA),
				"120000385",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(MA),
				"126001049",
			)
		})
	})

	t.Run("MG", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(MG),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(MG), []docCase{
				{"too short", "062.307.904/008"},
				{"hyphen instead of slash", "062.307.904-0081"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(MG),
				"062.307.904/0071",
				"062.307.904/0082",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(MG),
				"062.307.904/0081",
				"0623079040081",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(MG),
				"503.058.237/0015",
				"062.667.789/0073",
				"062.213.378/0083",
			)
		})
	})

	t.Run("MS", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(MS),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(MS), []docCase{
				{"prefix 29, not 28 or 50", "292909575"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(MS),
				"282909576",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(MS),
				"282909575",
				"283242809",
			)
		})
	})

	t.Run("MT", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(MT),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(MT), []docCase{
				{"too short", "013000001-9"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(MT),
				"0013000001-8",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(MT),
				"0013000001-9",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(MT),
				"0013144158-2",
			)
		})
	})

	t.Run("PA", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(PA),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(PA), []docCase{
				{"prefix 16, not 15 or from 75 to 79", "16999999-5"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(PA),
				"15999999-4",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(PA),
				"15999999-5",
				"75000002-3",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(PA),
				"15177432-3",
				"15771637-6",
				"15216699-8",
			)
		})
	})

	t.Run("PB", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(PB),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(PB), []docCase{
				{"too short", "0600001-5"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(PB),
				"06000001-6",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(PB),
				"06000001-5",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(PB),
				"16999351-5",
				"16900390-6",
			)
		})
	})

	t.Run("PE", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(PE),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(PE), []docCase{
				{"hyphen in the wrong place", "032141-840"},
				{"old format, too short", "18.1.001.000004-9"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(PE),
				"0321418-41",
				"0321418-30",
				"18.1.001.0000004-8",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(PE),
				"0321418-40",
				"032141840",
				"18.1.001.0000004-9",
				"18100100000049",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(PE),
				"0250099-07",
				"0163235-30",
				"0916078-76",
				"0992431-05",
				"0267359-20",
			)
		})
	})

	t.Run("PI", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(PI),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(PI), []docCase{
				{"too short", "01234567"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(PI),
				"012345678",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(PI),
				"012345679",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(PI),
				"194155722",
				"195195337",
				"194468976",
			)
		})
	})

	t.Run("PR", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalidCases(t, ieOf(PR), []docCase{
				{"empty", ""},
				{"letters", "AAAAAAAAAA"},
			})
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(PR), []docCase{
				{"too short", "123456785"},
				{"too long", "12345678501"},
				{"dots in the wrong places", "123.456.785-0"},
				{"dot in the wrong place", "12.345678-50"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(PR),
				"123.45678-40",
				"123.45678-51",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(PR),
				"123.45678-50",
				"12345678-50",
				"1234567850",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(PR),
				"099.02241-88",
				"101.79579-92",
			)
		})
	})

	t.Run("RJ", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(RJ),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(RJ), []docCase{
				{"hyphen in the wrong place", "91.018.0-39"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(RJ),
				"91.018.03-8",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(RJ),
				"91.018.03-9",
				"92.001.02-4",
			)
		})
	})

	t.Run("RN", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(RN),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(RN), []docCase{
				{"prefix 21 instead of 20", "21.040.040-1"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(RN),
				"20.040.040-2",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(RN),
				"20.040.040-1",
				"20.0.040.040-0",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(RN),
				"20.300.936-3",
				"20.301.158-9",
			)
		})
	})

	t.Run("RO", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(RO),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(RO), []docCase{
				{"format before 2000", "101.62521-3"},
				{"too short", "000000062521-3"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(RO),
				"0000000062521-4",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(RO),
				"0000000062521-3",
				"00000000625213",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(RO),
				"0000000025563-7",
			)
		})
	})

	t.Run("RR", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(RR),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(RR), []docCase{
				{"prefix 25 instead of 24", "25006628-1"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(RR),
				"24006628-2",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(RR),
				"24006628-1",
				"24001755-6",
				"24003429-0",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(RR),
				"24002153-4",
			)
		})
	})

	t.Run("RS", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(RS),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(RS), []docCase{
				{"too short", "224/365879"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(RS),
				"224/3658793",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(RS),
				"224/3658792",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(RS),
				"900/0000802",
			)
		})
	})

	t.Run("SC", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(SC),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(SC), []docCase{
				{"too short", "251.040.85"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(SC),
				"251.040.853",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(SC),
				"251.040.852",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(SC),
				"252.085.442",
			)
		})
	})

	t.Run("SE", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(SE),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(SE), []docCase{
				{"too short", "2712345-3"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(SE),
				"27123456-4",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(SE),
				"27123456-3",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(SE),
				"27077995-7",
				"27100154-2",
			)
		})
	})

	t.Run("SP", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalidCases(t, ieOf(SP), []docCase{
				{"empty", ""},
				{"letters", "AAAAAAAAAAAA"},
			})
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(SP), []docCase{
				{"too short", "11004249011"},
				{"too long", "1100424901145"},
				{"hyphens instead of dots", "110-042-490-114"},
				{"rural producer, lowercase p", "p-01100424.3/002"},
				{"rural producer, dot in the wrong place", "P-0110042.43/002"},
				{"rural producer, too long", "P-01100424.3/0021"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(SP),
				"110.042.491.114",
				"110.042.490.115",
				"P-01100424.4/002",
			)
		})

		t.Run("accepts the examples of the official rules", func(t *testing.T) {
			assertValid(t, ieOf(SP),
				"110.042.490.114",
				"110042490114",
				"P-01100424.3/002",
				"P011004243002",
				"P-01100424.3/999",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(SP),
				"310.035.324.119",
				"108.354.656.114",
				"142.270.790.110",
				"142.484.958.110",
				"102.654.009.110",
				"805.000.292.111",
			)
		})
	})

	t.Run("TO", func(t *testing.T) {
		t.Run("rejects invalid data", func(t *testing.T) {
			assertInvalid(t, ieOf(TO),
				"",
			)
		})

		t.Run("rejects invalid formats", func(t *testing.T) {
			assertInvalidCases(t, ieOf(TO), []docCase{
				{"too short", "29088778"},
				{"old format, with 11 digits", "29010227836"},
			})
		})

		t.Run("rejects invalid check digits", func(t *testing.T) {
			assertInvalid(t, ieOf(TO),
				"290887781",
			)
		})

		t.Run("accepts public IEs of companies", func(t *testing.T) {
			assertValid(t, ieOf(TO),
				"290887780",
				"299990176",
				"290319986",
			)
		})
	})
}
