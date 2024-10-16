# Map em GO

### Criando Map com valores

```
    mapEstatico := map[string]int{"sp": 10000000, "cg": 900000}
```

### Criando Map vazio para inserir valores depois

```
    m := make(map[string]int) // cria map

    m["sp"] = 1000000 // insere valores
	m["cg"] = 900000
```

### Verificando se valor existe dentro do map

É possivel verificar se um valor existe dentro do map já que o map retorna mais de um valor

```
    valor, existe := m["rj"]
    if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Valor nao existe")
	}
```

Essa verificação pode ser feita de outra forma também

```
    if valor, existe := m["rj"]; existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Valor nao existe")
	}
```
