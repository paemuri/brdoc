# BR Doc

[![Build Status][tag1img]][tag1link]
[![GoDoc][tag2img]][tag2link]
[![Go Report Card][tag3img]][tag3link]
[![codecov][tag4img]][tag4link]

CPF, CNPJ and CEP validator for Go!

Everything in this file, but the [License](#license) section, is in brazilian portuguese.

This package is also available in gopkg.in as [Nhanderu/brdoc.v1][1].

## Descrição


BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de documentos brasileiros, como CPF, CNPJ, CEP e (futuramente) RG.

## Uso

Principais funções:

- `func IsCPF(doc string) bool`
- `func IsCNPJ(doc string) bool`
- `func IsCEP(doc string, ufs ...FederativeUnit) bool`

Exemplo de valores válidos e inválidos:

```go
// Inválidos por causa da inconsistência do dígito:
brdoc.IsCPF("248.438.034-99") //=> false
brdoc.IsCNPJ("26.637.142/0001-00") //=> false

// Inválidos por causa do formato:
brdoc.IsCPF("248 438 034 80") //=> false
brdoc.IsCNPJ("26637142-0001.58") //=> false
brdoc.IsCEP("01 001.001") //=> false

// Inválidos por causa da UF:
brdoc.IsCEP("01001-001", brdoc.RJ) //=> false

// Válidos:
brdoc.IsCPF("248.438.034-80") //=> true
brdoc.IsCPF("24843803480") //=> true
brdoc.IsCNPJ("26.637.142/0001-58") //=> true
brdoc.IsCNPJ("26637142000158") //=> true
brdoc.IsCEP("01001-001", brdoc.SP) //=> true
brdoc.IsCEP("01001001", brdoc.SP) //=> true
```

## To-do list

- [x] validação de CPF
- [x] validação de CNPJ
- [x] validação de CEP
- [ ] validação de RG

Validação de RG não foi implementado porque cada estado tem as suas regras
e eu não estou com vontade de terminar isso. :zzz:

## License

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