package brdoc

import (
	"testing"
)

func TestIsCPF(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Random data.
		{"3467875434578764345789654", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},

		// Common invalid patterns.
		{"000.000.000-00", false},
		{"111.111.111-11", false},
		{"222.222.222-22", false},
		{"333.333.333-33", false},
		{"444.444.444-44", false},
		{"555.555.555-55", false},
		{"666.666.666-66", false},
		{"777.777.777-77", false},
		{"888.888.888-88", false},
		{"999.999.999-99", false},

		// Invalid digits.
		{"248.438.034-08", false},
		{"099.075.865-06", false},

		// Invalid format.
		{"248 438 034 80", false},
		{"099-075-865.60", false},

		// Valid.
		{"248.438.034-80", true},
		{"099.075.865-60", true},
	} {
		t.Logf("#%d CPF validation of %s should return %v: ", i, v.v, v.r)
		if IsCPF(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestIsCNPJ(t *testing.T) {
	for i, v := range []struct {
		v string
		r bool
	}{
		// Random data.
		{"3467875434578764345789654", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},

		// Common invalid patterns.
		{"00.000.000/0000-00", false},
		{"11.111.111/1111-11", false},
		{"22.222.222/2222-22", false},
		{"33.333.333/3333-33", false},
		{"44.444.444/4444-44", false},
		{"55.555.555/5555-55", false},
		{"66.666.666/6666-66", false},
		{"77.777.777/7777-77", false},
		{"88.888.888/8888-88", false},
		{"99.999.999/9999-99", false},

		// Invalid digits.
		{"26.637.142/0001-85", false},
		{"74.221.325/0001-03", false},

		// Invalid format.
		{"26-637-142.0001/58", false},
		{"74-221-325.0001/30", false},

		// Valid.
		{"26.637.142/0001-58", true},
		{"74.221.325/0001-30", true},
	} {
		t.Logf("#%d CNPJ validation of %s should return %v: ", i, v.v, v.r)
		if IsCNPJ(v.v) != v.r {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
