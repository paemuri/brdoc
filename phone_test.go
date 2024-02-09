package brdoc

import "testing"

func TestIsPhone(t *testing.T) {
	t.Parallel()

	for i, tc := range []struct {
		name          string
		expectedValid bool
		expectedUf    FederativeUnit
		v             string
	}{
		{"Valid_ShouldReturnTrueUfSP", true, SP, "+55(11)999999999"},
		{"Valid_ShouldReturnTrueUfSP", true, SP, "55(11)999999999"},
		{"Valid_ShouldReturnTrueUfPR", true, PR, "(41)999999999"},
		{"Valid_ShouldReturnTrueUfDF", true, DF, "(61)32222222"},
		{"Valid_ShouldReturnTrueUfAC", true, AC, "(68) 99988-1234"},
		{"Valid_ShouldReturnTrueUfPE", true, PE, "8198988888"},
		{"Valid_ShouldReturnTrueUfPE", true, PE, "558198988888"},
		{"Valid_ShouldReturnFalseUfAC", true, SP, "12 9999-9999"},
		{"InvalidDDI_ShouldReturnFalseUfAC", false, AC, "+01(11)999999999"},
		{"InvalidDDD_ShouldReturnFalseUfAC_1", false, AC, "(01)999999999"},
		{"InvalidDDD_ShouldReturnFalseUfAC_2", false, AC, "55(01)999999999"},
		{"InvalidPattern_ShouldReturnFalseUfAC_1", true, SP, "11 9999 9999"},
		{"InvalidPattern_ShouldReturnFalseUfAC_2", true, SP, "11 9 9999 9999"},
		{"InvalidPattern_ShouldReturnFalseUfAC_3", false, AC, "11.99999.9999"},
		{"InvalidPattern_ShouldReturnFalseUfAC_4", false, AC, "11 99999/9999"},
		{"InvalidPattern_ShouldReturnFalseUfAC_5", false, AC, "(11)9999999-99"},
		{"Teste_novo_1", true, SP, "+55 (11) 3340-2800"},
		{"Teste_novo_2", true, SP, "+55 (12) 3340 2801"},
		{"Teste_novo_3", true, SP, "+55 (13) 33402802"},
		{"Teste_novo_4", true, SP, "+55 (14)3340-2803"},
		{"Teste_novo_5", true, SP, "+55 (15)33402804"},
		{"Teste_novo_6", true, SP, "+55 (16)3340 2805"},
		{"Teste_novo_7", true, SP, "+55 (17) 9 6340-2806"},
		{"Teste_novo_8", true, SP, "+55 (18) 9 6340 2807"},
		{"Teste_novo_9", true, SP, "+55 (19) 9 63402808"},
		{"Teste_novo_10", true, RJ, "+55 (21)9 7340-2809"},
		{"Teste_novo_11", true, RJ, "+55 (22)9 73402810"},
		{"Teste_novo_12", true, RJ, "+55 (24)9 7340 2811"},
		{"Teste_novo_13", true, ES, "+55 (27) 98340-2812"},
		{"Teste_novo_14", true, ES, "+55 (28) 98340 2813"},
		{"Teste_novo_15", true, MG, "+55 (31) 983402814"},
		{"Teste_novo_16", true, MG, "+55 (32)99340-2815"},
		{"Teste_novo_17", true, MG, "+55 (33)993402816"},
		{"Teste_novo_18", true, MG, "+55 (34)99340 2817"},
		{"Teste_novo_19", true, MG, "+55 35 3340-2818"},
		{"Teste_novo_20", true, MG, "+55 37 33402819"},
		{"Teste_novo_21", true, MG, "+55 38 3340 2820"},
		{"Teste_novo_22", true, PR, "+55 41 9 6340-2821"},
		{"Teste_novo_23", true, PR, "+55 42 9 63402822"},
		{"Teste_novo_24", true, PR, "+55 43 9 6340 2823"},
		{"Teste_novo_25", true, PR, "+55 44 97340-2824"},
		{"Teste_novo_26", true, PR, "+55 45 973402825"},
		{"Teste_novo_27", true, PR, "+55 46 97340 2826"},
	} {
		tC := tc
		t.Run(testName(i, tC.name), func(t *testing.T) {
			t.Parallel()
			t.Helper()
			isValid, uf := IsPhone(tC.v)
			assertEqual(t, tC.expectedValid, isValid)
			assertEqual(t, tC.expectedUf, uf)
		})
	}
}
