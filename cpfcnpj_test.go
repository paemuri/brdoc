package brdoc

import . "testing"

func TestIsCPF(t *T) {
	t.Run("Invalid CPF format", func(t *T) {
		for _, v := range []string{
			"3467875434578764345789654",
			"",
			"#$%¨&*(ABCDEF",
		} {
			r := IsCPF(v)
			assert(t, v, r, false)
		}
	})
	t.Run("Common invalid patterns in CPF", func(t *T) {
		for _, v := range []string{
			"000.000.000-00",
			"111.111.111-11",
			"222.222.222-22",
			"333.333.333-33",
			"444.444.444-44",
			"555.555.555-55",
			"666.666.666-66",
			"777.777.777-77",
			"888.888.888-88",
			"999.999.999-99",
		} {
			r := IsCPF(v)
			assert(t, v, r, false)
		}
	})
	t.Run("Invalid digits in CPF", func(t *T) {
		for _, v := range []string{
			"248.438.034-08",
			"099.075.865-06",
		} {
			r := IsCPF(v)
			assert(t, v, r, false)
		}
	})
	t.Run("Valid CPF", func(t *T) {
		for _, v := range []string{
			"248.438.034-80",
			"099.075.865-60",
		} {
			r := IsCPF(v)
			assert(t, v, r, true)
		}
	})
}

func TestIsCNPJ(t *T) {
	t.Run("Invalid CNPJ format", func(t *T) {
		for _, v := range []string{
			"3467875434578764345789654",
			"",
			"#$%¨&*(ABCDEF",
		} {
			r := IsCNPJ(v)
			assert(t, v, r, false)
		}
	})
	t.Run("Common invalid patterns in CNPJ", func(t *T) {
		for _, v := range []string{
			"00.000.000/0000-00",
			"11.111.111/1111-11",
			"22.222.222/2222-22",
			"33.333.333/3333-33",
			"44.444.444/4444-44",
			"55.555.555/5555-55",
			"66.666.666/6666-66",
			"77.777.777/7777-77",
			"88.888.888/8888-88",
			"99.999.999/9999-99",
		} {
			r := IsCNPJ(v)
			assert(t, v, r, false)
		}
	})
	t.Run("Invalid values in CNPJ", func(t *T) {
		v := "26.637.142/0000-58"
		r := IsCNPJ(v)
		assert(t, v, r, false)
	})
	t.Run("Invalid digits in CNPJ", func(t *T) {
		for _, v := range []string{
			"26.637.142/0001-85",
			"74.221.325/0001-03",
		} {
			r := IsCNPJ(v)
			assert(t, v, r, false)
		}
	})
	t.Run("Valid CNPJ", func(t *T) {
		for _, v := range []string{
			"26.637.142/0001-58",
			"74.221.325/0001-30",
		} {
			r := IsCNPJ(v)
			assert(t, v, r, true)
		}
	})
}
