package brdoc

import (
	"regexp"
)

var (
	nfeRegexp = regexp.MustCompile(
		`^(?:[0-9A-Z]{4} ?){10}[0-9A-Z]{4}$`,
	)
	nfeCleanRegexp = regexp.MustCompile(`^[0-9]{6}[0-9A-Z]{12}[0-9]{26}$`)
	nfeUFs         = map[string]bool{
		"11": true, "12": true, "13": true, "14": true, "15": true, "16": true,
		"17": true, "21": true, "22": true, "23": true, "24": true, "25": true,
		"26": true, "27": true, "28": true, "29": true, "31": true, "32": true,
		"33": true, "35": true, "41": true, "42": true, "43": true, "50": true,
		"51": true, "52": true, "53": true,
	}
	nfeModels = map[string]bool{
		"55": true, // NF-e (MOC 7.0).
		"57": true, // CT-e (Ajuste SINIEF 9/07).
		"58": true, // MDF-e (Ajuste SINIEF 21/10).
		"62": true, // NFCom (Ajuste SINIEF 7/22).
		"63": true, // BP-e and BP-e TM (Ajuste SINIEF 1/17).
		"64": true, // GTV-e (Ajuste SINIEF 3/20).
		"65": true, // NFC-e (MOC 7.0).
		"66": true, // NF3e (Ajuste SINIEF 1/19).
		"67": true, // CT-e OS (Ajuste SINIEF 9/07).
	}
)

// IsNFE verifies if the given string is a valid access key of an NF-e. It also
// works for NFC-e, CT-e, MDF-e and the other electronic fiscal documents that
// use the same access key.
//
// The format and the check digit algorithm are defined by item 2.2.6 of the
// [MOC 7.0], and the alphanumeric CNPJ by the [NT 2025.001]. The emitter (CNPJ
// or CPF) is also validated, as defined by table 2-4 and item 2.2.7 of the
// [MOC 7.0].
//
// [MOC 7.0]: https://www.nfe.fazenda.gov.br/portal/exibirArquivo.aspx?conteudo=LrBx7WT9PuA=
// [NT 2025.001]: https://www.nfe.fazenda.gov.br/portal/exibirArquivo.aspx?conteudo=5ZkvIZt10mQ=
func IsNFE(doc string) bool {
	if !nfeRegexp.MatchString(doc) {
		return false
	}

	cleanTaxID(&doc)

	if !nfeCleanRegexp.MatchString(doc) {
		return false
	}

	month := doc[4:6]
	if !nfeUFs[doc[:2]] || month < "01" || month > "12" ||
		!nfeModels[doc[20:22]] {
		return false
	}

	if !validNFEEmitter(doc[6:20], doc[20:22], doc[22:25]) {
		return false
	}

	return toInt(rune(doc[43])) == calcMod11Digit(doc[:43])
}

// validNFEEmitter verifies if the emitter is a valid CNPJ or CPF, the latter
// preceded by zeros. As defined by the MOC 7.0, in an NF-e (model 55), the
// series from 910 to 969 are reserved for CPF (table 2-4), and in an NFC-e
// (model 65), the emitter is always a CNPJ (item 2.2.7). CPF is also accepted
// in the other series of an NF-e, as real NF-e of rural producers use them.
func validNFEEmitter(emitter, model, series string) bool {
	isCPF := emitter[:3] == "000" && IsCPF(emitter[3:])
	switch model {
	case "55":
		// Same as comparing ints, as both have 3 digits.
		if "910" <= series && series <= "969" {
			return isCPF
		}
		return isCPF || IsCNPJ(emitter)
	case "65":
		return IsCNPJ(emitter)
	default:
		return isCPF || IsCNPJ(emitter)
	}
}
