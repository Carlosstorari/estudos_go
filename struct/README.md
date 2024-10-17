## Struct no Go

É uma estrutura simple de agrupamento de dados definida pela palavra `type` e `struct`

```go
type endereco struct {
	rua string
	numero int
	cidade string
}
```

Para criar uma nova instancia dessa struct

```go
endereco := endereco {
		rua: "Rua X",
		numero: 15,
		cidade: "Campinas",
	}
```

### Encapsulamento usando struct

Para que um struct possa ser visto em um arquivo de outro package é necessario que **a primeira letra esteja maiuscula**, por exemplo:

é criado um model de `Endereco`

**Arquivo model dentro da pasta model**

```go
package model

type Endereco struct {
	Rua string
	Numero int
	Cidade string
}
```

**Arquivo main**

```go
package main

import (
	"fmt"
	model "golangestudos/struct/model"
)


func main() {
	endereco := model.Endereco {
		Rua: "Rua X",
		Numero: 15,
		Cidade: "Campinas",
	}

	fmt.Println(endereco)
}
```
