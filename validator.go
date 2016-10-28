package brdoc

import (
	"github.com/Nhanderu/tuyo/convert"
	valid "github.com/asaskevich/govalidator"
)

func init() {
	valid.CustomTypeTagMap.Set("cpf", valid.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return validate(i, IsCPF)
	}))
	valid.CustomTypeTagMap.Set("cnpj", valid.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return validate(i, IsCNPJ)
	}))
}

func validate(doc interface{}, v func(string) bool) bool {
	switch doc.(type) {
	case string:
		return v(convert.ToString(doc))
	default:
		return false
	}
}
