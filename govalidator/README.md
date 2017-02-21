# BR Doc Govalidator

[govalidator][1] integration for [BR Doc][2]!

Everything in this file, but the [License section](#license).

### Descrição

This package integrate o [BR Doc][2] with [govalidator][1]. Add values `"cpf"`, `"cnpj"`, (and in the future) `"cep"` for tag `valid`.

### Usage

```go
pessoa := struct {
  Nome string ``
  CPF string `valid:"cpf"`
  CNPJ string `valid:"cnpj"`
}{
  Nome: "Beto Carreiro",
  CPF: "248.438.034-80",
  CNPJ: "26.637.142/0001-58",
}
r, _ := govalidator.ValidateStruct(pessoa) //=> true
```

### License

This project code is in the public domain. See the [LICENSE file][3].

[1]: https://github.com/asaskevich/govalidator
[2]: https://github.com/Nhanderu/brdoc
[3]: https://github.com/Nhanderu/brdoc/blob/master/govalidator/LICENSE
