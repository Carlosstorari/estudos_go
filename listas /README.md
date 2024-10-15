# Listas em GO

### Criando listas

```
    lista := []int{4, 9, 6, 7} // cria lista estatica
    lista1 := make([]int, 1) //cria lista dinamica do tipo int com uma posição
```

### Adicionando novo item na lista

```
    lista1 = append(lista1, 2) // adiciona o valor 2 na lista1
```

### Ver tamanho da lista

retorna o numero inteiro que corresponde ao tamanho da lista

```
    len(lista)
```

### Pegar partes expecificas de uma lista

#### Pega até um determinado ponto do começo

```
    lista := []int{4, 9, 6, 7, 0, 5, 8, 5, 3}
    lista2 := lista[:3] // pega os tres primeiros itens da lista
```

lista2 = {4, 9, 6}

#### Pega apartir de um ponto até o final da lista

```
    lista := []int{4, 9, 6, 7, 0, 5, 8, 5, 3}
    lista2 := lista[4:] // pega os tres primeiros itens da lista
```

lista2 = {0, 5, 8, 5, 3}
