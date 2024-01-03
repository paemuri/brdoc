package brdoc

import "testing"

func TestIsPhone(t *testing.T) {
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
		{"Test_new_1", true, SP, "+55 (11) 3340-2800"},
		{"Test_new_2", true, SP, "+55 (12) 3340 2801"},
		{"Test_new_3", true, SP, "+55 (13) 33402802"},
		{"Test_new_4", true, SP, "+55 (14)3340-2803"},
		{"Test_new_5", true, SP, "+55 (15)33402804"},
		{"Test_new_6", true, SP, "+55 (16)3340 2805"},
		{"Test_new_7", true, SP, "+55 (17) 9 6340-2806"},
		{"Test_new_8", true, SP, "+55 (18) 9 6340 2807"},
		{"Test_new_9", true, SP, "+55 (19) 9 63402808"},
		{"Test_new_10", true, RJ, "+55 (21)9 7340-2809"},
		{"Test_new_11", true, RJ, "+55 (22)9 73402810"},
		{"Test_new_12", true, RJ, "+55 (24)9 7340 2811"},
		{"Test_new_13", true, ES, "+55 (27) 98340-2812"},
		{"Test_new_14", true, ES, "+55 (28) 98340 2813"},
		{"Test_new_15", true, MG, "+55 (31) 983402814"},
		{"Test_new_16", true, MG, "+55 (32)99340-2815"},
		{"Test_new_17", true, MG, "+55 (33)993402816"},
		{"Test_new_18", true, MG, "+55 (34)99340 2817"},
		{"Test_new_19", true, MG, "+55 35 3340-2818"},
		{"Test_new_20", true, MG, "+55 37 33402819"},
		{"Test_new_21", true, MG, "+55 38 3340 2820"},
		{"Test_new_22", true, PR, "+55 41 9 6340-2821"},
		{"Test_new_23", true, PR, "+55 42 9 63402822"},
		{"Test_new_24", true, PR, "+55 43 9 6340 2823"},
		{"Test_new_25", true, PR, "+55 44 97340-2824"},
		{"Test_new_26", true, PR, "+55 45 973402825"},
		{"Test_new_27", true, PR, "+55 46 97340 2826"},
	} {
		tC := tc
		t.Run(testName(i, tC.name), func(t *testing.T) {
			isValid, uf := IsPhone(tC.v)
			assertEqual(t, tC.expectedValid, isValid)
			assertEqual(t, tC.expectedUf, uf)
		})
	}
}
