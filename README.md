# BR Doc

[![Build Status][tag1img]][tag1link]
[![GoDoc][tag2img]][tag2link]
[![Go Report Card][tag3img]][tag3link]
[![codecov][tag4img]][tag4link]

CPF, CNPJ, CEP validator for Go!

Everything in this file, but the [To-do](#to-do-list) and [License](#license) sections, is in brazilian portuguese.

This package is also available in gopkg.in as [Nhanderu/brdoc.v1][1].

### Descrição

BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de documentos brasileiros, como CPF, CNPJ, (e futuramente) CEP e RG. Este pacote possui apenas as funções individuais de validação. Para utilizar da integração com o [govalidator][2], é necessário importar o pacote [BR Doc Govalidator][3].

### Uso

Principais funções:

- `func IsCPF(doc string) bool`
- `func IsCNPJ(doc string) bool`

Exemplo de valores válidos e inválidos:

```go
// Inválidos por causa da inconsistência do dígito:
brdoc.IsCPF("248.438.034-99") //=> false
brdoc.IsCNPJ("26.637.142/0001-00") //=> false

// Inválidos por causa do formato:
brdoc.IsCPF("248 438 034 80") //=> false
brdoc.IsCNPJ("26637142-0001.58") //=> false

// Válidos:
brdoc.IsCPF("248.438.034-80") //=> true
brdoc.IsCPF("24843803480") //=> true
brdoc.IsCNPJ("26.637.142/0001-58") //=> true
brdoc.IsCNPJ("26637142000158") //=> true
```

### To-do list

- [x] CPF validation
- [x] CNPJ validation
- [x] CEP validation
- [ ] RG validation

### License

This project code is in the public domain. See the [LICENSE file][4].

[1]: https://gopkg.in/Nhanderu/brdoc.v1
[2]: https://github.com/asaskevich/govalidator
[3]: https://github.com/Nhanderu/brdoc-govalidator
[4]: https://github.com/Nhanderu/brdoc/blob/master/LICENSE

[tag1img]: https://travis-ci.org/Nhanderu/brdoc.svg?branch=master
[tag1link]: https://travis-ci.org/Nhanderu/brdoc
[tag2img]: https://godoc.org/github.com/Nhanderu/brdoc?status.png
[tag2link]: https://godoc.org/github.com/Nhanderu/brdoc
[tag3img]: https://goreportcard.com/badge/github.com/Nhanderu/brdoc
[tag3link]: https://goreportcard.com/report/github.com/Nhanderu/brdoc
[tag4img]: https://codecov.io/gh/Nhanderu/brdoc/branch/master/graph/badge.svg
[tag4link]: https://codecov.io/gh/Nhanderu/brdoc