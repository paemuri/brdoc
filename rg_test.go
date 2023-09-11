package brdoc

import (
	"testing"
)

func TestIsRG(t *testing.T) {
	for i, tc := range []struct {
		name            string
		expectedIsValid bool
		expectedErr     error
		v               string
		uf              FederativeUnit
	}{
		{"InvalidData_ShouldReturnFalse", false, nil, "3467875434578764345789654", SP},
		{"InvalidData_ShouldReturnFalse", false, nil, "", SP},
		{"InvalidData_ShouldReturnFalse", false, nil, "AAAAAAA", SP},
		{"InvalidData_ShouldReturnFalse", false, nil, "34.112.513-A", SP},
		{"InvalidData_ShouldReturnFalse", false, nil, "34.112.513-X", SP},
		{"InvalidData_ShouldReturnFalse", false, nil, "34.112.513-x", SP},
		{"InvalidCheckDigit_ShouldReturnFalse", false, nil, "34.112.513-9", SP},
		{"ValidFormat_ShouldReturnTrue", true, nil, "34.112.513-1", SP},
		{"ValidFormat_ShouldReturnTrue", true, nil, "341125131", SP},
		{"ValidFormat_ShouldReturnTrue", true, nil, "25.540.324-0", RJ},
		{"ValidFormat_ShouldReturnTrue", true, nil, "255403240", RJ},
		{"ValidFormat_ShouldReturnTrue", true, nil, "39.406.714-9", RJ},
		{"ValidFormat_ShouldReturnTrue", true, nil, "24.454.119-X", RJ},
		{"ValidFormat_ShouldReturnTrue", true, nil, "24454119X", RJ},
		{"ValidFormat_ShouldReturnTrue", true, nil, "24.454.119-x", RJ},
		{"ValidFormat_ShouldReturnTrue", true, nil, "24454119x", RJ},
		{"InvalidData_ShouldReturnFalse", false, nil, "39.406.714-X", RJ},
		{"InvalidData_ShouldReturnFalse", false, nil, "39.406.714-0", RJ},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "39.406.714-0", PR},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "25.540.324-0", ES},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "255403240", BA},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "39.406.714-9", MT},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "24.454.119-X", RS},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "24.454.119-X", RS},
		{"NotImplementedUF_ShouldReturnError", false, errNotImplementedUF, "24454119X", DF},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			isValid, err := IsRG(tc.v, tc.uf)
			assertEqual(t, tc.expectedErr, err)
			if err == nil {
				assertEqual(t, tc.expectedIsValid, isValid)
			}
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
			assertEqual(t, tc.expected, calculateRGCheckDigit(tc.v))
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
			assertEqual(t, tc.expected, getRGCheckDigit(tc.v))
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
			assertEqual(t, tc.expected, cleanNonRGDigits(tc.v))
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
