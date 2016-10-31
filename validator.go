package brdoc

import (
	"github.com/Nhanderu/tuyo/convert"
	valid "github.com/asaskevich/govalidator"
)

func init() {
	valid.CustomTypeTagMap.Set("cpf", valid.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return IsCPF(convert.ToString(i))
	}))
	valid.CustomTypeTagMap.Set("cnpj", valid.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return IsCNPJ(convert.ToString(i))
	}))
}
