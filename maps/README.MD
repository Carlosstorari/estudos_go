# Map em GO

### Criando Map com valores

```go
    mapEstatico := map[string]int{"sp": 10000000, "cg": 900000}
```

### Criando Map vazio para inserir valores depois

```go
    m := make(map[string]int) // cria map

    m["sp"] = 1000000 // insere valores
	m["cg"] = 900000
```

### Verificando se valor existe dentro do map

É possivel verificar se um valor existe dentro do map já que o map retorna mais de um valor

```go
    valor, existe := m["rj"]
    if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Valor nao existe")
	}
```

Essa verificação pode ser feita de outra forma também

```go
    if valor, existe := m["rj"]; existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Valor nao existe")
	}
```

### Iterando por um map com for

Para iterar por um map é usado o range

```go
	m := make(map[string]int)

	m["sp"] = 1000000
	m["cg"] = 900000

    // itera o map
    for chave, valor := range m {
		fmt.Println("Cidade", chave, "H:", valor)
	}
```

### Deletando um elemento do map

Para deletar é usado o metodo `delete`

```go
	m := make(map[string]int)

	m["sp"] = 1000000
	m["cg"] = 900000
	m["rj"] = 600000

	delete(m, "rj")
```

Ele recebe como parametro o map `m` e a chave `"rj"`
