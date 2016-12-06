package brdoc

import (
	"github.com/Nhanderu/brdoc"
	"github.com/Nhanderu/tuyo/convert"
	valid "github.com/asaskevich/govalidator"
)

func init() {
	valid.CustomTypeTagMap.Set("cpf", valid.CustomTypeValidator(func(i, o interface{}) bool {
		return brdoc.IsCPF(convert.ToString(i))
	}))
	valid.CustomTypeTagMap.Set("cnpj", valid.CustomTypeValidator(func(i, o interface{}) bool {
		return brdoc.IsCNPJ(convert.ToString(i))
	}))
}
