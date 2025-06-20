package brdoc

import (
	"testing"
)

func TestIsCPF(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAA", false},

		{"InvalidPattern", "000.000.000-00", false},
		{"InvalidPattern", "111.111.111-11", false},
		{"InvalidPattern", "222.222.222-22", false},
		{"InvalidPattern", "333.333.333-33", false},
		{"InvalidPattern", "444.444.444-44", false},
		{"InvalidPattern", "555.555.555-55", false},
		{"InvalidPattern", "666.666.666-66", false},
		{"InvalidPattern", "777.777.777-77", false},
		{"InvalidPattern", "888.888.888-88", false},
		{"InvalidPattern", "999.999.999-99", false},

		{"InvalidDigits", "248.438.034-08", false},
		{"InvalidDigits", "099.075.865-06", false},

		{"InvalidFormat", "248 438 034 80", false},
		{"InvalidFormat", "099-075-865.60", false},

		{"Valid", "248.438.034-80", true},
		{"Valid", "09907586560", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsCPF(tc.doc))
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		valid bool
	}{
		{"InvalidData", "3467875434578764345789654", false},
		{"InvalidData", "", false},
		{"InvalidData", "AAAAAAAAAAAAAA", false},

		{"InvalidPattern", "00.000.000/0000-00", false},
		{"InvalidPattern", "11.111.111/1111-11", false},
		{"InvalidPattern", "22.222.222/2222-22", false},
		{"InvalidPattern", "33.333.333/3333-33", false},
		{"InvalidPattern", "44.444.444/4444-44", false},
		{"InvalidPattern", "55.555.555/5555-55", false},
		{"InvalidPattern", "66.666.666/6666-66", false},
		{"InvalidPattern", "77.777.777/7777-77", false},
		{"InvalidPattern", "88.888.888/8888-88", false},
		{"InvalidPattern", "99.999.999/9999-99", false},

		{"InvalidDigits", "26.637.142/0001-85", false},
		{"InvalidDigits", "74.221.325/0001-03", false},

		{"InvalidFormat", "26-637-142.0001/58", false},
		{"InvalidFormat", "74-221-325.0001/30", false},

		{"Valid", "26.637.142/0001-58", true},
		{"Valid", "74221325000130", true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsCNPJ(tc.doc))
		})
	}
}
