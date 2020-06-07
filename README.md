# BR Doc

[![Build Status][tag1img]][tag1link]
[![GoDoc][tag2img]][tag2link]
[![Go Report Card][tag3img]][tag3link]
[![codecov][tag4img]][tag4link]

CPF, CNPJ, CEP, CNH, RENAVAM and license plate validator for Go!

Everything in this file, but the [License](#license) section, is in portuguese.

## Descrição

BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de documentos
brasileiros, como CPF, CNPJ, CEP, CNH, RENAVAM, placa veicular e futuramente RG.

Aceito PRs de todas as formas. Está permitido escrever em português, também. :)

## Uso

Principais funções:

- `func IsCPF(doc string) bool`
- `func IsCNPJ(doc string) bool`
- `func IsCEP(doc string, ufs ...FederativeUnit) bool`
- `func IsCNH(doc string) bool`
- `func IsRENAVAM(doc string) bool`
- `func IsPlate(doc string) bool`
- `func IsNationalPlate(doc string) bool`
- `func IsMercosulPlate(doc string) bool`

## To-do list

- [x] validação de CPF
- [x] validação de CNPJ
- [x] validação de CEP
- [x] validação de CNH (obrigado @eminetto!)
- [x] validação de RENAVAM (obrigado @leogregianin!)
- [x] validação de placa veicular
- [ ] entender como funciona hífen na CNH (ex. `0067600300-1`)
- [ ] validação de RG

Validação de RG não foi implementado porque cada estado tem as suas regras e eu não
estou com vontade de terminar isso. :zzz:

## License

This project code is in the public domain. See the [LICENSE file][1].

[1]: https://github.com/Nhanderu/brdoc/blob/master/LICENSE

[tag1img]: https://travis-ci.org/Nhanderu/brdoc.svg?branch=master
[tag1link]: https://travis-ci.org/Nhanderu/brdoc
[tag2img]: https://godoc.org/github.com/Nhanderu/brdoc?status.png
[tag2link]: https://godoc.org/github.com/Nhanderu/brdoc
[tag3img]: https://goreportcard.com/badge/github.com/Nhanderu/brdoc
[tag3link]: https://goreportcard.com/report/github.com/Nhanderu/brdoc
[tag4img]: https://codecov.io/gh/Nhanderu/brdoc/branch/master/graph/badge.svg
[tag4link]: https://codecov.io/gh/Nhanderu/brdoc
