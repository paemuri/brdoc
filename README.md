# BR Doc

[![Build Status](https://travis-ci.org/Nhanderu/brdoc.svg?branch=master)][tag1]
[![GoDoc](https://godoc.org/github.com/Nhanderu/brdoc?status.png)][tag2]
[![Go Report Card](https://goreportcard.com/badge/github.com/Nhanderu/brdoc)][tag3]
[![codecov](https://codecov.io/gh/Nhanderu/brdoc/branch/master/graph/badge.svg)][tag4]

CPF, CNPJ, CEP validator for Go, with [govalidator][1] integration!

Everything in this file, but the [License section](#license).

### Description

BR Doc is a package for validation of both the format and the digits of Brazilian documents, such as CPF, CNPJ, (and in the future) CEP This package has only the individual validation functions. To use integration with [govalidator] [1], it is necessary to import the package [BR Doc Govalidator] [2].

### Usage

- `func IsCPF (doc string) bool`
- `func IsCNPJ (doc string) bool`

Example of valid and invalid values:

```go
// Invalid because of the inconsistency of the digit:
Brdoc.IsCPF ("248.438.034-99") // => false
Brdoc.IsCNPJ ("26.637.142 / 0001-00") // => false

// Invalid because of the format:
Brdoc.IsCPF ("248 438 034 80") // => false
Brdoc.IsCNPJ ("26637142-0001.58") // => false

// Valid:
Brdoc.IsCPF ("248.438.034-80") // => true
Brdoc.IsCPF ("24843803480") // => true
Brdoc.IsCNPJ ("26.637.142 / 0001-58") // => true
Brdoc.IsCNPJ ("26637142000158") // => true
```

### To-do list

- [x] CPF validation
- [x] CNPJ validation
- [ ] CEP validation
- [ ] RG validation

### License

This project code is in the public domain. See the [LICENSE file][3].

[1]: https://github.com/asaskevich/govalidator
[2]: https://github.com/Nhanderu/brdoc/govalidator
[3]: https://github.com/Nhanderu/brdoc/blob/master/LICENSE

[tag1]: https://travis-ci.org/Nhanderu/brdoc
[tag2]: (https://godoc.org/github.com/Nhanderu/brdoc)
[tag3]: (https://goreportcard.com/report/github.com/Nhanderu/brdoc)
[tag4]: (https://codecov.io/gh/Nhanderu/brdoc)