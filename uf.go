package brdoc

// FederativeUnit represents a first-level administrative division in Brazil: either a state or a federal district.
type FederativeUnit uint8

const (
	AC FederativeUnit = iota
	AL
	AP
	AM
	BA
	CE
	DF
	ES
	GO
	MA
	MT
	MS
	MG
	PA
	PB
	PR
	PE
	PI
	RJ
	RN
	RS
	RO
	RR
	SC
	SP
	SE
	TO
)
