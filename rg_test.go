package brdoc

import (
	"testing"
)

func TestIsRG(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAA"},
		{"InvalidData_ShouldReturnFalse", false, "34.112.513-A"},
		{"InvalidData_ShouldReturnFalse", false, "34.112.513-X"},
		{"InvalidData_ShouldReturnFalse", false, "34.112.513-x"},
		{"InvalidCheckDigit_ShouldReturnFalse", false, "34.112.513-9"},
		{"ValidFormat_ShouldReturnTrue", true, "34.112.513-1"},
		{"ValidFormat_ShouldReturnTrue", true, "341125131"},
		{"ValidFormat_ShouldReturnTrue", true, "25.540.324-0"},
		{"ValidFormat_ShouldReturnTrue", true, "255403240"},
		{"ValidFormat_ShouldReturnTrue", true, "39.406.714-9"},
		{"ValidFormat_ShouldReturnTrue", true, "24.454.119-X"},
		{"ValidFormat_ShouldReturnTrue", true, "24454119X"},
		{"ValidFormat_ShouldReturnTrue", true, "24.454.119-x"},
		{"ValidFormat_ShouldReturnTrue", true, "24454119x"},
		{"InvalidData_ShouldReturnFalse", false, "39.406.714-X"},
		{"InvalidData_ShouldReturnFalse", false, "39.406.714-0"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsRG(tc.v))
		})
	}
}

func TestCalculateRGCheckDigit(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected int
		v        string
	}{
		{"CalculateRGCheckDigit_ShouldReturn_1", 1, "34.112.513-1"},
		{"CalculateRGCheckDigit_ShouldReturn_1", 1, "341125131"},
		{"CalculateRGCheckDigit_ShouldReturn_9", 9, "39.406.714-9"},
		{"CalculateRGCheckDigit_ShouldReturn_9", 9, "394067149"},
		{"CalculateRGCheckDigit_ShouldReturn_11", 11, "25.540.324-0"},
		{"CalculateRGCheckDigit_ShouldReturn_11", 11, "255403240"},
		{"CalculateRGCheckDigit_ShouldReturn_10", 10, "24.454.119-X"},
		{"CalculateRGCheckDigit_ShouldReturn_10", 10, "24454119X"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertIntEqual(t, tc.expected, calculateRGCheckDigit(tc.v))
		})
	}
}

func TestGetRGCheckDigit(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected int
		v        string
	}{
		{"GetRGCheckDigit_ShouldReturn_1", 1, "34.112.513-1"},
		{"GetRGCheckDigit_ShouldReturn_1", 1, "341125131"},
		{"GetRGCheckDigit_ShouldReturn_9", 9, "39.406.714-9"},
		{"GetRGCheckDigit_ShouldReturn_9", 9, "394067149"},
		{"GetRGCheckDigit_ShouldReturn_11", 11, "25.540.324-0"},
		{"GetRGCheckDigit_ShouldReturn_11", 11, "255403240"},
		{"GetRGCheckDigit_ShouldReturn_10", 10, "24.454.119-X"},
		{"GetRGCheckDigit_ShouldReturn_10", 10, "24454119X"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertIntEqual(t, tc.expected, getRGCheckDigit(tc.v))
		})
	}
}

func TestCleanNonRGDigits(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected string
		v        string
	}{
		{"CleanNonRGDigits", "341125131", "34.112.513-1"},
		{"CleanNonRGDigits", "341125131", "341125131"},
		{"CleanNonRGDigits", "394067149", "39.406.714-9"},
		{"CleanNonRGDigits", "394067149", "394067149"},
		{"CleanNonRGDigits", "24454119X", "24.454.119-X"},
		{"CleanNonRGDigits", "24454119X", "24454119X"},
		{"CleanNonRGDigits", "24454119x", "24.454.119-x"},
		{"CleanNonRGDigits", "24454119x", "24454119x"},
		{"CleanNonRGDigits", "24454119x", "24//45///41--1.9x"},
		{"CleanNonRGDigits", "", "AAaaaAAaaa"},
		{"CleanNonRGDigits", "XxXxXx", "XxXxXx"},
		{"CleanNonRGDigits", "XxXxXx", "XxXyxXx"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertStringEqual(t, tc.expected, cleanNonRGDigits(tc.v))
		})
	}
}

func TestIsRGValidDigit(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        rune
	}{
		{"ValidDigit_ShouldReturnTrue", true, 'x'},
		{"ValidDigit_ShouldReturnTrue", true, 'X'},
		{"ValidDigit_ShouldReturnTrue", true, '1'},
		{"ValidDigit_ShouldReturnTrue", true, '2'},
		{"ValidDigit_ShouldReturnTrue", true, '3'},
		{"ValidDigit_ShouldReturnTrue", true, '9'},
		{"ValidDigit_ShouldReturnTrue", true, '0'},
		{"InvalidDigit_ShouldReturnFalse", false, 'Y'},
		{"InvalidDigit_ShouldReturnFalse", false, 'A'},
		{"InvalidDigit_ShouldReturnFalse", false, '-'},
		{"InvalidDigit_ShouldReturnFalse", false, '.'},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, isRGValidDigit(tc.v))
		})
	}
}
