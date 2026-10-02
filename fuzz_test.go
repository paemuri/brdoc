//go:build go1.18
// +build go1.18

package brdoc

import (
	"testing"
)

// FuzzAll calls every function with random data, to find panics. Native fuzzing
// requires Go 1.18, so this file is ignored by older versions.
func FuzzAll(f *testing.F) {
	for _, seed := range []string{
		"",
		"123.456.789-09",
		"12.ABC.345/01DE-35",
		"(11) 99999-9999",
		"01001-000",
		"ABC1D23",
		"0001234-56.2024.8.26.0100",
		"10490.05505 77222.133348 77777.777713 4 32420000032112",
		"00020126580014br.gov.bcb.pix0136123e4567-e12b-12d1-a456-4266554400005204000053039865802BR5913Fulano de Tal6008BRASILIA62070503***63041D3D",
	} {
		f.Add(seed)
	}

	ufs := []UF{
		AC, AL, AP, AM, BA, CE, DF, ES, GO, MA, MT, MS, MG, PA, PB, PR, PE,
		PI, RJ, RN, RS, RO, RR, SC, SP, SE, TO, "", "XX",
	}

	f.Fuzz(func(t *testing.T, doc string) {
		for _, fn := range []func(string) bool{
			IsCPF, IsCNPJ, IsPlate, IsVoterID, IsCNH, IsPIS, IsRENAVAM, IsCNS,
			IsCNJ, IsCivilCertificate, IsNFE, IsBankSlip, IsCNO, IsPixKey,
			IsPixBRCode,
		} {
			fn(doc)
		}

		IsPhone(doc)
		IsPhoneFrom(doc, SP, RJ)
		IsCEP(doc)
		IsCEPFrom(doc, SP, RJ)
		for _, uf := range ufs {
			IsIE(doc, uf)
		}
	})
}
