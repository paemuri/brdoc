package brdoc

// IsIE verifies if `doc` is a valid IE (Inscrição Estadual) of the given `uf`.
func IsIE(doc string, uf UF) bool {
	// The rules of each UF are published by [1].
	// [1]: http://www.sintegra.gov.br/insc_est.html.

	switch uf {
	case PR:
		return isIEPR(doc)
	case SP:
		return isIESP(doc)
	default:
		return false
	}
}
