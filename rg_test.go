package brdoc

import (
	"fmt"
	"testing"
)

func TestIsRG(t *testing.T) {
	for _, uf := range []UF{SP, RJ} {
		for i, tc := range []struct {
			name     string
			expected bool
			v        string
		}{
			{"Valid_ShouldReturnTrue", true, "34.112.513-1"},
			{"Valid_ShouldReturnTrue", true, "341125131"},
			{"Valid_ShouldReturnTrue", true, "25.540.324-0"},
			{"Valid_ShouldReturnTrue", true, "255403240"},
			{"Valid_ShouldReturnTrue", true, "39.406.714-9"},
			{"Valid_ShouldReturnTrue", true, "394067149"},
			{"Valid_ShouldReturnTrue", true, "24.454.119-X"},
			{"Valid_ShouldReturnTrue", true, "24454119X"},
			{"Valid_ShouldReturnTrue", true, "24.454.119-x"},
			{"Valid_ShouldReturnTrue", true, "24454119x"},

			{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
			{"InvalidData_ShouldReturnFalse", false, ""},
			{"InvalidData_ShouldReturnFalse", false, "AAAAAAA"},
			{"InvalidDigit_ShouldReturnFalse", false, "34.112.513-A"},
			{"InvalidDigit_ShouldReturnFalse", false, "34.112.513-X"},
			{"InvalidDigit_ShouldReturnFalse", false, "34.112.513-x"},
			{"InvalidDigit_ShouldReturnFalse", false, "34.112.513-9"},
			{"InvalidDigit_ShouldReturnFalse", false, "39.406.714-X"},
			{"InvalidDigit_ShouldReturnFalse", false, "39.406.714-0"},
		} {
			name := fmt.Sprintf("%s_%s", uf, tc.name)
			t.Run(testName(i, name), func(t *testing.T) {
				isValid, err := IsRG(tc.v, uf)
				assertEqual(t, err, nil)
				assertEqual(t, tc.expected, isValid)
			})
		}
	}
}

func TestIsRG_NotImplemented(t *testing.T) {
	for i, uf := range []UF{AC, AL, AP, AM, BA, CE, DF, ES, GO, MA, MT, MS, MG, PA, PB, PR, PE, PI, RN, RS, RO, RR, SC, SE, TO} {
		name := fmt.Sprintf("%s_NotImplemented", uf)
		t.Run(testName(i, name), func(t *testing.T) {
			_, err := IsRG("341125131", uf)
			assertEqual(t, err != nil, true) // Maybe a better assert.
		})
	}
}

func TestCalcRGDigit(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected int
		v        string
	}{
		{"CalcRGDigit", 1, "341125131"},
		{"CalcRGDigit", 9, "394067149"},
		{"CalcRGDigit", 11, "255403240"},
		{"CalcRGDigit", 10, "24454119X"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, calcRGDigit(tc.v))
		})
	}
}

func TestGetRGDigit(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected int
		v        string
	}{
		{"GetRGDigit", 1, "341125131"},
		{"GetRGDigit", 9, "394067149"},
		{"GetRGDigit", 11, "255403240"},
		{"GetRGDigit", 10, "24454119X"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, getRGDigit(tc.v))
		})
	}
}

func TestCleanRG(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected string
		v        string
	}{
		{"CleanRG", "341125131", "34.112.513-1"},
		{"CleanRG", "341125131", "341125131"},
		{"CleanRG", "394067149", "39.406.714-9"},
		{"CleanRG", "394067149", "394067149"},
		{"CleanRG", "24454119X", "24.454.119-X"},
		{"CleanRG", "24454119X", "24454119X"},
		{"CleanRG", "24454119x", "24.454.119-x"},
		{"CleanRG", "24454119x", "24454119x"},
		{"CleanRG", "24454119x", "24//45///41--1.9x"},
		{"CleanRG", "", "AAaaaAAaaa"},
		{"CleanRG", "XxXxXx", "XxXxXx"},
		{"CleanRG", "XxXxXx", "XxXyxXx"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			doc := tc.v
			cleanRG(&doc)
			assertEqual(t, tc.expected, doc)
		})
	}
}
