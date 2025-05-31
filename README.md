# BR Doc

[![License][badge-1-img]][badge-1-link]
[![go.dev][badge-2-img]][badge-2-link]
[![Go Report Card][badge-3-img]][badge-3-link]

CPF, CNPJ, CEP, CNH, PIS/PASEP, RENAVAM, CNS, license plate and voter ID
validator for Go!

Everything in this file, but the [License](#license) section, is in portuguese.

## Descrição

BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de
documentos brasileiros.

Aceito PRs de todas as formas. Está permitido escrever em português, também. :)

## Uso

Principais funções:

- `func IsCPF(doc string) bool`
- `func IsCNPJ(doc string) bool`
- `func IsCEP(doc string, ufs ...FederativeUnit) bool`
- `func IsCNH(doc string) bool`
- `func IsPIS(doc string) bool`
- `func IsRENAVAM(doc string) bool`
- `func IsPlate(doc string) bool`
- `func IsNationalPlate(doc string) bool`
- `func IsMercosulPlate(doc string) bool`
- `func IsCNS(doc string) bool`
- `func IsRG(doc string, uf FederativeUnit) bool`
- `func IsVoterID(doc string) bool`
- `func IsPhone(phone string) (isValid bool, uf FederativeUnit)`

## Coisas a fazer

- [x] validação de CPF
- [x] validação de CNPJ
- [x] validação de CEP
- [x] validação de CNH (obrigado @eminetto!)
- [x] validação de RENAVAM (obrigado @leogregianin!)
- [x] validação de placa veicular
- [x] validação de CNS (obrigado @renatosuero!)
- [x] validação de título de eleitor (obrigado @tiagompalte!)
- [x] validação de telefone (obrigado @tiagompalte e @ArtusC!)
- [ ] validação de RG
  - [x] SP (obrigado @robas!)
  - [x] RJ (obrigado @robas!)
  - [ ] demais estados

## License

This project code is in the public domain. See the [LICENSE file][1].

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you shall be in the public domain, without any
additional terms or conditions.

[1]: ./LICENSE

[badge-1-img]: https://img.shields.io/github/license/paemuri/brdoc?style=flat-square
[badge-1-link]: https://github.com/paemuri/brdoc/blob/master/LICENSE
[badge-2-img]: https://img.shields.io/badge/go.dev-reference-007d9c?style=flat-square&logo=go&logoColor=white
[badge-2-link]: https://pkg.go.dev/github.com/paemuri/brdoc
[badge-3-img]: https://goreportcard.com/badge/github.com/paemuri/brdoc?style=flat-square
[badge-3-link]: https://goreportcard.com/report/github.com/paemuri/brdoc
