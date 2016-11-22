package brdoc

import (
	. "testing"

	valid "github.com/asaskevich/govalidator"
)

type cpfTest struct {
	CPF string `valid:"cpf"`
}
type cnpjTest struct {
	CNPJ string `valid:"cnpj"`
}

func TestIntegrationWithGoValidator(t *T) {
	t.Run("Struct with invalid CPF", func(t *T) {
		for _, cpf := range []string{
			"3467875434578764345789654",
			"",
			"000.000.000-00",
			"111.222.333-69",
		} {
			v := cpfTest{CPF: cpf}
			r, _ := valid.ValidateStruct(v)
			assert(t, cpf, r, false)
		}
	})
	t.Run("Struct with valid CPF", func(t *T) {
		cpf := "111.222.333-96"
		v := cpfTest{CPF: cpf}
		r, _ := valid.ValidateStruct(v)
		assert(t, cpf, r, true)
	})
	t.Run("Struct with invalid CNPJ", func(t *T) {
		for _, cnpj := range []string{
			"3467875434578764345789654",
			"",
			"00.000.000/0000-00",
			"11.222.333/4444-79",
		} {
			v := cnpjTest{CNPJ: cnpj}
			r, _ := valid.ValidateStruct(v)
			assert(t, cnpj, r, false)
		}
	})
	t.Run("Struct with valid CNPJ", func(t *T) {
		cnpj := "11.222.333/4444-97"
		v := cnpjTest{CNPJ: cnpj}
		r, _ := valid.ValidateStruct(v)
		assert(t, cnpj, r, true)
	})
}
