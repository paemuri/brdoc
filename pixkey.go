package brdoc

import (
	"regexp"
)

var (
	pixKeyCPFRegexp   = regexp.MustCompile(`^\d{11}$`)
	pixKeyCNPJRegexp  = regexp.MustCompile(`^[0-9A-Z]{14}$`)
	pixKeyPhoneRegexp = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	pixKeyEVPRegexp   = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	pixKeyEmailRegexp = regexp.MustCompile("^[a-z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$")
)

// IsPixKey verifies if the given string is a valid Pix key: a CPF or CNPJ
// without mask, a phone number in the E.164 format (e.g. +5561912345678), an
// email in lowercase with up to 77 characters, or a random key (EVP), a UUID in
// lowercase.
//
// The format of each type of key is defined by the [API do DICT], and the
// formats without mask by section 1 of the [Manual Operacional do DICT].
//
// [API do DICT]: https://www.bcb.gov.br/content/estabilidadefinanceira/pix/API-DICT.html
// [Manual Operacional do DICT]: https://www.bcb.gov.br/content/estabilidadefinanceira/pix/Regulamento_Pix/X_ManualOperacionaldoDICT.pdf
func IsPixKey(key string) bool {
	switch {
	case pixKeyCPFRegexp.MatchString(key):
		return IsCPF(key)
	case pixKeyCNPJRegexp.MatchString(key):
		return IsCNPJ(key)
	}

	return pixKeyPhoneRegexp.MatchString(key) ||
		pixKeyEVPRegexp.MatchString(key) ||
		(len(key) <= 77 && pixKeyEmailRegexp.MatchString(key))
}
