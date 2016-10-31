# BR Doc

CPF, CNPJ, CEP and RG for [govalidator][1]!

The description of this package is in brazilian portuguese.

### Descrição

BR Doc é um pacote para validação, tanto do formato quanto dos dígitos, de documentos brasileiros, como CPF, CNPJ, (e futuramente) CEP e RG. Foi pensado para se usar com o [govalidator][1], mas é possível usar os métodos sem ele.

### Uso

Principais métodos:

```go
brdoc.IsCPF("111.111.111-11")
brdoc.IsCNPJ("11.111.111/1111-80")
```

A assinatura de ambos os métodos:

```go
func IsCPF(doc string) bool
```

Exemplo de valores válidos e inválidos:

````go
// Inválidos por causa da inconsistência do dígito:
brdoc.IsCPF("111.111.111-00") //=> false
brdoc.IsCNPJ("11.111.111/1111-99") //=> false

// Inválidos por causa do formato:
brdoc.IsCPF("111 111 111 11") //=> false
brdoc.IsCNPJ("11111111-1111.80") //=> false

// Válidos:
brdoc.IsCPF("111.111.111-11") //=> true
brdoc.IsCPF("11111111111") //=> true
brdoc.IsCNPJ("11.111.111/1111-80") //=> true
brdoc.IsCNPJ("11111111111180") //=> true
```

Usando o [govalidator][1]:

```go
pessoa := struct {
  Nome string ``
  CPF string `valid:"cpf"`
  CNPJ string `valid:"cnpj"`
}{
  Nome: "Beto Carreiro",
  CPF: "111.111.111-11",
  CNPJ: "11.111.111/1111-80",
}
govalidator.ValidateStruct(pessoa) //=> true
```

[1]: https://github.com/asaskevich/govalidator

### License

This project code is in the public domain.
