# BR Doc

[govalidator][1] integration for [BR Doc][2]!

The description of this package is in brazilian portuguese.

### Descrição

Este pacote tem como objetivo integrar o [BR Doc][2] com o [govalidator][1]. Adiciona os valores `"cpf"`, `"cnpj"`, (e futuramente) `"cep"` e `"rg"` para a tag `valid`.

### Uso

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

[1]: https://github.com/asaskevich/govalidator
[1]: https://github.com/Nhanderu/brdoc

### License

This project code is in the public domain. See the `LICENSE` file.
