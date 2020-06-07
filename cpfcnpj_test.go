package brdoc

import (
	"testing"
)

func TestIsCPF(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAAAAA"},
		{"InvalidPattern_ShouldReturnFalse", false, "000.000.000-00"},
		{"InvalidPattern_ShouldReturnFalse", false, "111.111.111-11"},
		{"InvalidPattern_ShouldReturnFalse", false, "222.222.222-22"},
		{"InvalidPattern_ShouldReturnFalse", false, "333.333.333-33"},
		{"InvalidPattern_ShouldReturnFalse", false, "444.444.444-44"},
		{"InvalidPattern_ShouldReturnFalse", false, "555.555.555-55"},
		{"InvalidPattern_ShouldReturnFalse", false, "666.666.666-66"},
		{"InvalidPattern_ShouldReturnFalse", false, "777.777.777-77"},
		{"InvalidPattern_ShouldReturnFalse", false, "888.888.888-88"},
		{"InvalidPattern_ShouldReturnFalse", false, "999.999.999-99"},
		{"InvalidDigits_ShouldReturnFalse", false, "248.438.034-08"},
		{"InvalidDigits_ShouldReturnFalse", false, "099.075.865-06"},
		{"InvalidFormat_ShouldReturnFalse", false, "248 438 034 80"},
		{"InvalidFormat_ShouldReturnFalse", false, "099-075-865.60"},
		{"Valid_ShouldReturnTrue", true, "248.438.034-80"},
		{"Valid_ShouldReturnTrue", true, "099.075.865-60"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCPF(tc.v))
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	for i, tc := range []struct {
		name     string
		expected bool
		v        string
	}{
		{"InvalidData_ShouldReturnFalse", false, "3467875434578764345789654"},
		{"InvalidData_ShouldReturnFalse", false, ""},
		{"InvalidData_ShouldReturnFalse", false, "AAAAAAAAAAAAAA"},
		{"InvalidPattern_ShouldReturnFalse", false, "00.000.000/0000-00"},
		{"InvalidPattern_ShouldReturnFalse", false, "11.111.111/1111-11"},
		{"InvalidPattern_ShouldReturnFalse", false, "22.222.222/2222-22"},
		{"InvalidPattern_ShouldReturnFalse", false, "33.333.333/3333-33"},
		{"InvalidPattern_ShouldReturnFalse", false, "44.444.444/4444-44"},
		{"InvalidPattern_ShouldReturnFalse", false, "55.555.555/5555-55"},
		{"InvalidPattern_ShouldReturnFalse", false, "66.666.666/6666-66"},
		{"InvalidPattern_ShouldReturnFalse", false, "77.777.777/7777-77"},
		{"InvalidPattern_ShouldReturnFalse", false, "88.888.888/8888-88"},
		{"InvalidPattern_ShouldReturnFalse", false, "99.999.999/9999-99"},
		{"InvalidDigits_ShouldReturnFalse", false, "26.637.142/0001-85"},
		{"InvalidDigits_ShouldReturnFalse", false, "74.221.325/0001-03"},
		{"InvalidFormat_ShouldReturnFalse", false, "26-637-142.0001/58"},
		{"InvalidFormat_ShouldReturnFalse", false, "74-221-325.0001/30"},
		{"Valid_ShouldReturnTrue", true, "26.637.142/0001-58"},
		{"Valid_ShouldReturnTrue", true, "74.221.325/0001-30"},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEqual(t, tc.expected, IsCNPJ(tc.v))
		})
	}
}
