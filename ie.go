package brdoc

import (
	"regexp"
	"strconv"
)

// ieRule is one of the formats of IE of a UF.
type ieRule struct {
	// pattern validates the format of the IE, including its mask.
	pattern *regexp.Regexp
	// weights has the weights of each check digit. Each check digit is
	// calculated with the digits before it, and so its position is the number
	// of its weights.
	weights [][]int
	// calcDigit calculates a check digit from the weighted sum.
	calcDigit func(sum int) int
	// validDigits verifies the check digits of the cleaned IE, for the rules
	// that do not fit in weights and calcDigit.
	validDigits func(doc string) bool
}

var (
	ieWeights = []int{9, 8, 7, 6, 5, 4, 3, 2}

	// ieRules has the rules of each UF, with the source of each one.
	ieRules = map[UF][]ieRule{
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_AC.html
		AC: {{
			pattern: regexp.MustCompile(`^01\.?\d{3}\.?\d{3}/?\d{3}-?\d{2}$`),
			weights: [][]int{
				{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
				{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
			},
			calcDigit: calcIEMod11Digit,
		}},
		// The digit is described as the remainder of the sum times 10 divided
		// by 11, which is the same as `calcIEMod11Digit`. Unlike the rules
		// published by the SINTEGRA, the third digit, formerly the type of
		// company, is not restricted.
		// [1]: https://www.sefaz.al.gov.br/calculo
		AL: {{
			pattern:   regexp.MustCompile(`^24\d{7}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_AM.html
		AM: {{
			pattern:   regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}-?\d$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_AP.html
		AP: {{
			pattern:     regexp.MustCompile(`^03\d{7}$`),
			validDigits: validIEAPDigits,
		}},
		// There are IEs with 8 and 9 digits.
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_BA.html
		BA: {{
			pattern:     regexp.MustCompile(`^\d{6,7}-?\d{2}$`),
			validDigits: validIEBADigits,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_CE.html
		CE: {{
			pattern:   regexp.MustCompile(`^\d{8}-?\d$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_DF.html
		DF: {{
			pattern: regexp.MustCompile(`^\d{11}-?\d{2}$`),
			weights: [][]int{
				{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
				{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
			},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_ES.html
		ES: {{
			pattern:   regexp.MustCompile(`^\d{9}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_GO.html
		GO: {{
			pattern:   regexp.MustCompile(`^(1[01]|2\d)\.?\d{3}\.?\d{3}-?\d$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_MA.html
		MA: {{
			pattern:   regexp.MustCompile(`^12\d{7}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_MG.html
		MG: {{
			pattern:     regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}/?\d{4}$`),
			validDigits: validIEMGDigits,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_MS.html
		MS: {{
			pattern:   regexp.MustCompile(`^(28|50)\d{7}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_MT.html
		MT: {{
			pattern:   regexp.MustCompile(`^\d{10}-?\d$`),
			weights:   [][]int{{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_PA.html
		PA: {{
			pattern:   regexp.MustCompile(`^(15|7[5-9])\d{6}-?\d$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_PB.html
		PB: {{
			pattern:   regexp.MustCompile(`^\d{8}-?\d$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// There are current IEs (e-Fisco) and old IEs (CACEPE).
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_PE.html
		PE: {
			{
				pattern: regexp.MustCompile(`^\d{7}-?\d{2}$`),
				weights: [][]int{
					{8, 7, 6, 5, 4, 3, 2},
					{9, 8, 7, 6, 5, 4, 3, 2},
				},
				calcDigit: calcIEMod11Digit,
			},
			{
				pattern:   regexp.MustCompile(`^\d{2}\.?\d\.?\d{3}\.?\d{7}-?\d$`),
				weights:   [][]int{{5, 4, 3, 2, 1, 9, 8, 7, 6, 5, 4, 3, 2}},
				calcDigit: calcIEMod11Minus10Digit,
			},
		},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_PI.html
		PI: {{
			pattern:   regexp.MustCompile(`^\d{9}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// The format is also described by [2].
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_PR.html
		// [2]: https://www.fazenda.pr.gov.br/servicos/Empresa/Cadastro-de-Contribuintes-do-ICMS/Saber-como-se-calcula-o-digito-verificador-da-inscricao-estadual-kZrX1Bol
		PR: {{
			pattern: regexp.MustCompile(`^\d{3}\.?\d{5}-?\d{2}$`),
			weights: [][]int{
				{3, 2, 7, 6, 5, 4, 3, 2},
				{4, 3, 2, 7, 6, 5, 4, 3, 2},
			},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_RJ.html
		RJ: {{
			pattern:   regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{2}-?\d$`),
			weights:   [][]int{{2, 7, 6, 5, 4, 3, 2}},
			calcDigit: calcIEMod11Digit,
		}},
		// There are IEs with 9 and 10 digits. The digit is described as the
		// remainder of the sum times 10 divided by 11, which is the same as
		// `calcIEMod11Digit`.
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_RN.html
		RN: {
			{
				pattern:   regexp.MustCompile(`^20\.?\d{3}\.?\d{3}-?\d$`),
				weights:   [][]int{ieWeights},
				calcDigit: calcIEMod11Digit,
			},
			{
				pattern:   regexp.MustCompile(`^20\.?\d\.?\d{3}\.?\d{3}-?\d$`),
				weights:   [][]int{{10, 9, 8, 7, 6, 5, 4, 3, 2}},
				calcDigit: calcIEMod11Digit,
			},
		},
		// Only the rules adopted from 2000 on, as older IEs, with 9 digits
		// including the code of the municipality, were converted to it.
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_RO.html
		RO: {{
			pattern:   regexp.MustCompile(`^\d{13}-?\d$`),
			weights:   [][]int{{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}},
			calcDigit: calcIEMod11Minus10Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_RR.html
		RR: {{
			pattern:   regexp.MustCompile(`^24\d{6}-?\d$`),
			weights:   [][]int{{1, 2, 3, 4, 5, 6, 7, 8}},
			calcDigit: calcIEMod9Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_RS.html
		RS: {{
			pattern:   regexp.MustCompile(`^\d{3}/?\d{7}$`),
			weights:   [][]int{{2, 9, 8, 7, 6, 5, 4, 3, 2}},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_SC.html
		SC: {{
			pattern:   regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_SE.html
		SE: {{
			pattern:   regexp.MustCompile(`^\d{8}-?\d$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
		// There are IEs of industrials and merchants, and of rural producers.
		// The last 3 digits of the rural producers are not verified.
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_SP.html
		SP: {
			{
				pattern: regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}\.?\d{3}$`),
				weights: [][]int{
					{1, 3, 4, 5, 6, 7, 8, 10},
					{3, 2, 10, 9, 8, 7, 6, 5, 4, 3, 2},
				},
				calcDigit: calcIERemainderDigit,
			},
			{
				pattern:   regexp.MustCompile(`^P-?\d{8}\.?\d/?\d{3}$`),
				weights:   [][]int{{1, 3, 4, 5, 6, 7, 8, 10}},
				calcDigit: calcIERemainderDigit,
			},
		},
		// The format is defined by Article 3 of [2], which replaced the format
		// of 11 digits described by [1], not accepted anymore. The check digit
		// follows [1], which already ignored the 2 digits removed by [2].
		// [1]: http://www.sintegra.gov.br/Cad_Estados/cad_TO.html
		// [2]: https://dtri.sefaz.to.gov.br/legislacao/ntributaria/portarias/sefaz/Portaria676-02.htm
		TO: {{
			pattern:   regexp.MustCompile(`^\d{9}$`),
			weights:   [][]int{ieWeights},
			calcDigit: calcIEMod11Digit,
		}},
	}
)

// IsIE verifies if `doc` is a valid IE (Inscrição Estadual) of the given `uf`.
// "ISENTO", used in place of the IE by those exempt from it, is not a valid IE.
//
// The rules of each UF are published by the [SINTEGRA], unless stated otherwise
// in the code.
//
// [SINTEGRA]: http://www.sintegra.gov.br/insc_est.html
func IsIE(doc string, uf UF) bool {
	for _, rule := range ieRules[uf] {
		if rule.pattern.MatchString(doc) {
			return rule.valid(doc)
		}
	}

	return false
}

// valid verifies the check digits of `doc`, already matched by the pattern.
func (rule ieRule) valid(doc string) bool {
	cleanNonDigits(&doc)

	// Not official logic: reject documents with all digits equal.
	if allEq(doc) {
		return false
	}

	if rule.validDigits != nil {
		return rule.validDigits(doc)
	}

	for _, weights := range rule.weights {
		sum := 0
		for i, weight := range weights {
			sum += toInt(rune(doc[i])) * weight
		}

		if toInt(rune(doc[len(weights)])) != rule.calcDigit(sum) {
			return false
		}
	}

	return true
}

// validIEAPDigits verifies the check digit of an IE of AP, whose calculation
// depends on the range of its number.
func validIEAPDigits(doc string) bool {
	base, _ := strconv.Atoi(doc[:8])

	p, d := 0, 0
	switch {
	case base <= 3017000:
		p, d = 5, 0
	case base <= 3019022:
		p, d = 9, 1
	}

	sum := p
	for i, weight := range ieWeights {
		sum += toInt(rune(doc[i])) * weight
	}

	digit := 11 - sum%11
	switch digit {
	case 10:
		digit = 0
	case 11:
		digit = d
	}

	return toInt(rune(doc[8])) == digit
}

// validIEBADigits verifies the check digits of an IE of BA. The digits are
// calculated with modulo 10 or 11 depending on the first digit, for IEs with 8
// digits, or on the second one, for IEs with 9 digits. The second check digit
// is calculated before the first one.
func validIEBADigits(doc string) bool {
	base := doc[:len(doc)-2]

	calcDigit := func(sum int) int {
		digit := 10 - sum%10
		if digit == 10 {
			return 0
		}

		return digit
	}
	switch base[len(base)-6] {
	case '6', '7', '9':
		calcDigit = calcIEMod11Digit
	}

	calc := func(doc string) int {
		sum := 0
		for i, r := range doc {
			sum += toInt(r) * (len(doc) + 1 - i)
		}

		return calcDigit(sum)
	}

	second := calc(base)
	first := calc(base + strconv.Itoa(second))

	return toInt(rune(doc[len(doc)-2])) == first &&
		toInt(rune(doc[len(doc)-1])) == second
}

// validIEMGDigits verifies the check digits of an IE of MG.
func validIEMGDigits(doc string) bool {
	// The first digit is calculated with a zero inserted after the code of the
	// municipality, and with the sum of the digits of the products.
	sum := 0
	for i, r := range doc[:3] + "0" + doc[3:11] {
		product := toInt(r) * (i%2 + 1)
		sum += product/10 + product%10
	}
	// Not official logic: the rules do not state the digit when the sum is
	// already a multiple of 10, so it is 0.
	first := (10 - sum%10) % 10

	sum = 0
	for i, weight := range []int{3, 2, 11, 10, 9, 8, 7, 6, 5, 4, 3} {
		sum += toInt(rune(doc[i])) * weight
	}
	sum += first * 2

	return toInt(rune(doc[11])) == first &&
		toInt(rune(doc[12])) == calcIEMod11Digit(sum)
}

// calcIEMod11Digit returns 11 minus the remainder of `sum` divided by 11, or 0
// if it is 10 or 11.
func calcIEMod11Digit(sum int) int {
	digit := 11 - sum%11
	if digit >= 10 {
		return 0
	}

	return digit
}

// calcIEMod11Minus10Digit returns 11 minus the remainder of `sum` divided by
// 11, minus 10 if it is 10 or 11.
func calcIEMod11Minus10Digit(sum int) int {
	digit := 11 - sum%11
	if digit >= 10 {
		return digit - 10
	}

	return digit
}

// calcIERemainderDigit returns the rightmost digit of the remainder of `sum`
// divided by 11.
func calcIERemainderDigit(sum int) int {
	return sum % 11 % 10
}

// calcIEMod9Digit returns the remainder of `sum` divided by 9.
func calcIEMod9Digit(sum int) int {
	return sum % 9
}
