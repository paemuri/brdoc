package brdoc

import (
	valid "github.com/asaskevich/govalidator"
	. "testing"
)

type cpfTest struct {
	CPF string `valid:"cpf"`
}
type cnpjTest struct {
	CNPJ string `valid:"cnpj"`
}

func TestIntegrationWithGoValidator(t *T) {
	t.Run("Struct with invalid CPF", func(t *T) {
		cpf := ""
		v := cpfTest{CPF: cpf}
		r, _ := valid.ValidateStruct(v)
		assert(t, cpf, r, false)

		cpf = "#$%¨&*(ABCDEF"
		v = cpfTest{CPF: cpf}
		r, _ = valid.ValidateStruct(v)
		assert(t, cpf, r, false)

		cpf = "111.111.111-22"
		v = cpfTest{CPF: cpf}
		r, _ = valid.ValidateStruct(v)
		assert(t, cpf, r, false)
	})
	t.Run("Struct with invalid CNPJ", func(t *T) {
		cnpj := ""
		v := cnpjTest{CNPJ: cnpj}
		r, _ := valid.ValidateStruct(v)
		assert(t, cnpj, r, false)

		cnpj = "#$%¨&*(ABCDEF"
		v = cnpjTest{CNPJ: cnpj}
		r, _ = valid.ValidateStruct(v)
		assert(t, cnpj, r, false)

		cnpj = "11.111.111/1111-00"
		v = cnpjTest{CNPJ: cnpj}
		r, _ = valid.ValidateStruct(v)
		assert(t, cnpj, r, false)
	})
	t.Run("Struct with valid CPF", func(t *T) {
		cpf := "111.111.111-11"
		v := cpfTest{CPF: cpf}
		r, _ := valid.ValidateStruct(v)
		assert(t, cpf, r, true)
	})
	t.Run("Struct with valid CNPJ", func(t *T) {
		cnpj := "11.111.111/1111-80"
		v := cnpjTest{CNPJ: cnpj}
		r, _ := valid.ValidateStruct(v)
		assert(t, cnpj, r, true)
	})
}
