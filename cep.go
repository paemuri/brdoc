package brdoc

// IsCEP verifies if `doc` is a valid CEP.
// `ufs` represents the possible Federative Units the CEP should matches.
// If none is provided, it validates the document for any state/district.
func IsCEP(doc string, ufs ...FederativeUnits) bool {
	return false
}
