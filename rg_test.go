package brdoc

import (
	"fmt"
	"testing"
)

func TestIsRG(t *testing.T) {
	for _, uf := range []UF{SP, RJ} {
		for i, tc := range []struct {
			name  string
			doc   string
			valid bool
		}{
			{"InvalidData", "3467875434578764345789654", false},
			{"InvalidData", "", false},
			{"InvalidData", "AAAAAAA", false},

			{"InvalidDigit", "34.112.513-A", false},
			{"InvalidDigit", "34.112.513-X", false},
			{"InvalidDigit", "34.112.513-x", false},
			{"InvalidDigit", "34.112.513-9", false},
			{"InvalidDigit", "39.406.714-X", false},
			{"InvalidDigit", "39.406.714-0", false},

			{"Valid", "34.112.513-1", true},
			{"Valid", "25.540.324-0", true},
			{"Valid", "39.406.714-9", true},
			{"Valid", "24.454.119-X", true},
			{"Valid", "24.454.119-x", true},
			{"Valid", "341125131", true},
			{"Valid", "255403240", true},
			{"Valid", "394067149", true},
			{"Valid", "24454119X", true},
			{"Valid", "24454119x", true},
		} {
			name := fmt.Sprintf("%s_%s", uf, tc.name)
			t.Run(testName(i, name), func(t *testing.T) {
				valid, err := IsRG(tc.doc, uf)
				assertEq(t, tc.valid, valid)
				assertEq(t, nil, err)
			})
		}
	}
}

func TestIsRG_NotImplemented(t *testing.T) {
	for i, uf := range []UF{AC, AL, AP, AM, BA, CE, DF, ES, GO, MA, MT, MS, MG, PA, PB, PR, PE, PI, RN, RS, RO, RR, SC, SE, TO} {
		name := fmt.Sprintf("%s_NotImplemented", uf)
		t.Run(testName(i, name), func(t *testing.T) {
			_, err := IsRG("", uf)
			assertNotEq(t, nil, err)
			assertEq(t, "federative unit not implemented", err.Error())
		})
	}
}

func TestCalcRGDigit(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		digit int
	}{
		{"CalcRGDigit", "34112513 ", 1},
		{"CalcRGDigit", "39406714 ", 9},
		{"CalcRGDigit", "24454119 ", 10},
		{"CalcRGDigit", "25540324 ", 11},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.digit, calcRGDigit(tc.doc))
		})
	}
}

func TestGetRGDigit(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		digit int
	}{
		{"GetRGDigit", "XXX1", 1},
		{"GetRGDigit", "XXX2", 2},
		{"GetRGDigit", "XXX3", 3},
		{"GetRGDigit", "XXX4", 4},
		{"GetRGDigit", "XXX5", 5},
		{"GetRGDigit", "XXX6", 6},
		{"GetRGDigit", "XXX7", 7},
		{"GetRGDigit", "XXX8", 8},
		{"GetRGDigit", "XXX9", 9},
		{"GetRGDigit", "XXXX", 10},
		{"GetRGDigit", "XXX0", 11},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.digit, getRGDigit(tc.doc))
		})
	}
}

func TestCleanRG(t *testing.T) {
	for i, tc := range []struct {
		name   string
		doc    string
		result string
	}{
		{"CleanRG", "34.112.513-1", "341125131"},
		{"CleanRG", "341125131", "341125131"},
		{"CleanRG", "39.406.714-9", "394067149"},
		{"CleanRG", "394067149", "394067149"},
		{"CleanRG", "24.454.119-X", "24454119X"},
		{"CleanRG", "24454119X", "24454119X"},
		{"CleanRG", "24.454.119-x", "24454119x"},
		{"CleanRG", "24454119x", "24454119x"},
		{"CleanRG", "24//45///41--1.9x", "24454119x"},
		{"CleanRG", "AAaaaAAaaa", ""},
		{"CleanRG", "abcdefghijklmnopqrstuvwxyz0123456789", "x0123456789"},
		{"CleanRG", "XxXxXx", "XxXxXx"},
		{"CleanRG", "XxXyxXx", "XxXxXx"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			doc := tc.doc
			cleanRG(&doc)
			assertEq(t, tc.result, doc)
		})
	}
}
