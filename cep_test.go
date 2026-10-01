package brdoc

import (
	"testing"
)

// isCEP reports if `doc` is valid by both `IsCEP` and `IsCEPFrom`, to be used
// with `assertValid` and `assertInvalid`.
func isCEP(doc string) bool {
	valid, _ := IsCEP(doc)
	return valid || IsCEPFrom(doc)
}

// assertCEPUF checks that `doc` is valid and related to `uf`.
func assertCEPUF(t *testing.T, doc string, uf UF) {
	valid, docUF := IsCEP(doc)
	if !valid || docUF != uf || !IsCEPFrom(doc, uf) {
		t.Errorf("expected %q to be a valid CEP from %s, got %s", doc, uf, docUF)
	}
}

// assertCEPNotUF checks that `doc` is valid, but not related to any of `ufs`.
func assertCEPNotUF(t *testing.T, doc string, ufs ...UF) {
	if valid, _ := IsCEP(doc); !valid {
		t.Errorf("expected %q to be valid", doc)
	}
	if IsCEPFrom(doc, ufs...) {
		t.Errorf("expected %q not to be from %v", doc, ufs)
	}
}

func TestIsCEP(t *testing.T) {
	t.Run("rejects invalid data", func(t *testing.T) {
		assertInvalid(t, isCEP,
			"3467875434578764345789654",
			"",
			"AAAAAAAA",
		)
	})

	t.Run("rejects the prefix 00", func(t *testing.T) {
		assertInvalid(t, isCEP,
			"00000-000",
			"00801-000",
		)
	})

	t.Run("rejects invalid formats", func(t *testing.T) {
		assertInvalid(t, isCEP,
			"10000 000",
			"25000.000",
			"29500/000",
			"1-0000000",
			"10-000000",
			"100-00000",
			"1000-0000",
			"100000-00",
			"1000000-0",
			"1000x000",
			"1000x-000",
			"1000000x",
			"10000-00x",
		)
	})

	t.Run("returns the UF of each range", func(t *testing.T) {
		t.Run("with hyphen", func(t *testing.T) {
			assertCEPUF(t, "10000-000", SP)
			assertCEPUF(t, "25000-000", RJ)
			assertCEPUF(t, "29500-000", ES)
			assertCEPUF(t, "35000-000", MG)
			assertCEPUF(t, "45000-000", BA)
			assertCEPUF(t, "49500-000", SE)
			assertCEPUF(t, "53000-000", PE)
			assertCEPUF(t, "57500-000", AL)
			assertCEPUF(t, "58500-000", PB)
			assertCEPUF(t, "59500-000", RN)
			assertCEPUF(t, "62000-000", CE)
			assertCEPUF(t, "64500-000", PI)
			assertCEPUF(t, "65500-000", MA)
			assertCEPUF(t, "67000-000", PA)
			assertCEPUF(t, "68900-000", AP)
			assertCEPUF(t, "69100-000", AM)
			assertCEPUF(t, "69300-000", RR)
			assertCEPUF(t, "69600-000", AM)
			assertCEPUF(t, "69900-000", AC)
			assertCEPUF(t, "71500-000", DF)
			assertCEPUF(t, "72800-000", GO)
			assertCEPUF(t, "73200-000", DF)
			assertCEPUF(t, "74500-000", GO)
			assertCEPUF(t, "76800-000", RO)
			assertCEPUF(t, "77500-000", TO)
			assertCEPUF(t, "78500-000", MT)
			assertCEPUF(t, "79500-000", MS)
			assertCEPUF(t, "84000-000", PR)
			assertCEPUF(t, "89000-000", SC)
			assertCEPUF(t, "95000-000", RS)
		})

		t.Run("without hyphen", func(t *testing.T) {
			assertCEPUF(t, "10000000", SP)
			assertCEPUF(t, "25000000", RJ)
			assertCEPUF(t, "29500000", ES)
			assertCEPUF(t, "35000000", MG)
			assertCEPUF(t, "45000000", BA)
			assertCEPUF(t, "49500000", SE)
			assertCEPUF(t, "53000000", PE)
			assertCEPUF(t, "57500000", AL)
			assertCEPUF(t, "58500000", PB)
			assertCEPUF(t, "59500000", RN)
			assertCEPUF(t, "62000000", CE)
			assertCEPUF(t, "64500000", PI)
			assertCEPUF(t, "65500000", MA)
			assertCEPUF(t, "67000000", PA)
			assertCEPUF(t, "68900000", AP)
			assertCEPUF(t, "69100000", AM)
			assertCEPUF(t, "69300000", RR)
			assertCEPUF(t, "69600000", AM)
			assertCEPUF(t, "69900000", AC)
			assertCEPUF(t, "71500000", DF)
			assertCEPUF(t, "72800000", GO)
			assertCEPUF(t, "73200000", DF)
			assertCEPUF(t, "74500000", GO)
			assertCEPUF(t, "76800000", RO)
			assertCEPUF(t, "77500000", TO)
			assertCEPUF(t, "78500000", MT)
			assertCEPUF(t, "79500000", MS)
			assertCEPUF(t, "84000000", PR)
			assertCEPUF(t, "89000000", SC)
			assertCEPUF(t, "95000000", RS)
		})
	})

	t.Run("rejects UFs that are not of the range", func(t *testing.T) {
		assertCEPNotUF(t, "10000-000", RJ)
		assertCEPNotUF(t, "25000-000", ES)
		assertCEPNotUF(t, "29500-000", MG)
		assertCEPNotUF(t, "35000-000", BA)
		assertCEPNotUF(t, "45000-000", SE)
		assertCEPNotUF(t, "49500-000", PE)
		assertCEPNotUF(t, "53000-000", AL)
		assertCEPNotUF(t, "57500-000", PB)
		assertCEPNotUF(t, "58500-000", RN)
		assertCEPNotUF(t, "59500-000", CE)
		assertCEPNotUF(t, "62000-000", PI)
		assertCEPNotUF(t, "64500-000", MA)
		assertCEPNotUF(t, "65500-000", PA)
		assertCEPNotUF(t, "67000-000", AP)
		assertCEPNotUF(t, "68900-000", AM)
		assertCEPNotUF(t, "69100-000", RR)
		assertCEPNotUF(t, "69300-000", AM)
		assertCEPNotUF(t, "69600-000", AC)
		assertCEPNotUF(t, "69900-000", DF)
		assertCEPNotUF(t, "71500-000", GO)
		assertCEPNotUF(t, "72800-000", DF)
		assertCEPNotUF(t, "73200-000", GO)
		assertCEPNotUF(t, "74500-000", RO)
		assertCEPNotUF(t, "76800-000", TO)
		assertCEPNotUF(t, "77500-000", MT)
		assertCEPNotUF(t, "78500-000", MS)
		assertCEPNotUF(t, "79500-000", PR)
		assertCEPNotUF(t, "84000-000", SC)
		assertCEPNotUF(t, "89000-000", RS)
		assertCEPNotUF(t, "95000-000", SP)
		assertCEPNotUF(t, "29500-000", MT, MS, MG)
	})

	t.Run("accepts a list of UFs", func(t *testing.T) {
		if !IsCEPFrom("10000-000", BA, SP, MG) {
			t.Errorf("expected %q to be from %v", "10000-000", []UF{BA, SP, MG})
		}
		if !IsCEPFrom("25000-000", RJ, RJ, RJ) {
			t.Errorf("expected %q to be from %v", "25000-000", []UF{RJ, RJ, RJ})
		}
		if !IsCEPFrom("29500-000") {
			t.Errorf("expected %q to be from %v", "29500-000", []UF{})
		}
	})
}
