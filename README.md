# BR Doc

[![License][badge-1-img]][badge-1-link]
[![go.dev][badge-2-img]][badge-2-link]
[![CI][badge-3-img]][badge-3-link]
[![Version][badge-4-img]][badge-4-link]

Brazilian documents validator for Go.

Everything in this file, except this section and the [License](#license)
one, is in Portuguese.

## Descrição

BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de
documentos brasileiros.

Aceito PRs de todas as formas. Está permitido escrever em português, também. :)

## Uso

Instalação:

```sh
go get github.com/paemuri/brdoc/v4
```

Funções públicas:

- `func IsCPF(doc string) bool`
- `func IsCNPJ(doc string) bool`
- `func IsPhone(phone string) (valid bool, ufs []UF)`
- `func IsPhoneFrom(phone string, ufs ...UF) bool`
- `func IsCEP(doc string) (valid bool, uf UF)`
- `func IsCEPFrom(doc string, ufs ...UF) bool`
- `func IsPlate(doc string) bool`: placa veicular, nacional ou Mercosul
- `func IsNationalPlate(doc string) bool`: placa no padrão nacional antigo
  (`AAA-1111`)
- `func IsMercosulPlate(doc string) bool`: placa no padrão Mercosul
  (`AAA1A11`)
- `func IsVoterID(doc string) bool`: título de eleitor
- `func IsCNH(doc string) bool`
- `func IsPIS(doc string) bool`: PIS/PASEP, também funciona para NIS e NIT
- `func IsRENAVAM(doc string) bool`
- `func IsCNS(doc string) bool`
- `func IsIE(doc string, uf UF) bool`: Inscrição Estadual
- `func IsCNJ(doc string) bool`: número de processo judicial
- `func IsCivilCertificate(doc string) bool`: matrícula de certidão de registro
  civil (nascimento, casamento ou óbito)
- `func IsNFE(doc string) bool`: chave de acesso de NF-e, NFC-e, CT-e, MDF-e e
  outros documentos fiscais eletrônicos

Tipos públicos:

- `type UF string`: representa cada unidade federativa

Exemplo:

```go
import "github.com/paemuri/brdoc/v4"

func main() {
	brdoc.IsCPF("123.456.789-09")                  // true
	valid, ufs := brdoc.IsPhone("(11) 99999-9999") // true, [SP]
}
```

## License

This project code is in the public domain. See the [LICENSE file][1].

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you shall be in the public domain, without any
additional terms or conditions.

[1]: ./LICENSE

[badge-1-img]: https://img.shields.io/github/license/paemuri/brdoc?style=flat-square
[badge-1-link]: https://github.com/paemuri/brdoc/blob/main/LICENSE
[badge-2-img]: https://img.shields.io/badge/go.dev-reference-007d9c?style=flat-square&logo=go&logoColor=white
[badge-2-link]: https://pkg.go.dev/github.com/paemuri/brdoc/v4
[badge-3-img]: https://img.shields.io/github/actions/workflow/status/paemuri/brdoc/ci.yaml?branch=main&style=flat-square
[badge-3-link]: https://github.com/paemuri/brdoc/actions/workflows/ci.yaml
[badge-4-img]: https://img.shields.io/github/v/tag/paemuri/brdoc?sort=semver&style=flat-square
[badge-4-link]: https://github.com/paemuri/brdoc/tags
