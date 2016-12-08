# BR Doc

CPF, CNPJ, CEP and RG validator for Go, with [govalidator][1] integration!

Everything in this file, but the [License section](#license), is in brazilian portuguese.

### Descrição

BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de documentos brasileiros, como CPF, CNPJ, (e futuramente) CEP e RG. Este pacote possui apenas as funções individuais de validação. Para utilizar da integração com o [govalidator][1], é necessário importar o pacote [BR Doc Govalidator][2].

### Uso

Principais funções:

- `func IsCPF(doc string) bool`
- `func IsCNPJ(doc string) bool`

Exemplo de valores válidos e inválidos:

````go
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

### License

This project code is in the public domain. See the [LICENSE file][3].

[1]: https://github.com/asaskevich/govalidator
[2]: https://github.com/Nhanderu/brdoc/govalidator
[3]: https://github.com/Nhanderu/brdoc/blob/master/LICENSE
