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
	Rua string // os atributos que forem compartilhados tem que ter primeira lera maiuscula
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

### Tipos dentro de tipos no struct

É possivel criar uma struct que contenha uma outra struct como um de seus elementos

```go
package model

type Pessoa struct {
	Nome string
	Endereco Endereco
}
```

A atribuição de valores dessa struct é feita da seguinte forma

```go
	endereco := model.Endereco {
		Rua: "Rua X",
		Numero: 15,
		Cidade: "Campinas",
	}

	pessoa := model.Pessoa {
		Nome: "Carlos",
		Endereco: endereco,
	}
```

### Metodos de struct

É possivel adicionar um metodo dentro da logica de uma struct, para isso é necessario colocar qual struct quer usar, antes do nome da função

```go
package model

import "time"

type Pessoa struct {
	Nome string
	Endereco Endereco
	DataDeNascimento time.Time
}

func (p Pessoa)IdadeAtual() int { // metodo dentro da struct
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
```

No exemplo acima o metodo `IdadeAtual` esta vinculado a struct `Pessoa` o metodo pode ser usado quando a struct for chamada da seguinte forma

```go
	pessoa := model.Pessoa {
		Nome: "Carlos",
		Endereco: endereco,
		DataDeNascimento: time.Date(1995, 8, 26, 0, 0, 0, 0, time.Local),
	}
	idade := pessoa.IdadeAtual()
```
