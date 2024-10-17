# Listas em GO

### Uso do `&`

O `&` é usado para retornar o endereço de memória de de uma variavel

```go
    x := 5
    fmt.Println(&x)
```

resultado:

```
    0xc00009a040
```

### Endereços de memória, passando paramentros em funções

Quando uma variavel é passada como parametro em uma função, a função faz uma copia dessa variavel. Com isso 2 espaços são alocados na memória, em grandes aplicações isso pode prejudicar o desenpenho.

```go
   func main() {
	x := 5
	fmt.Println(&x) //0xc000012110
	ImprimirValores(x)
}

func ImprimirValores(x int){
	fmt.Println(&x) //0xc00009a000
}
```

O `x` na função main tem o endereço `0xc000012110`, já o `x` na função `ImprimirValores` tem o endereço de variavel `0xc00009a000`.

### Ponteiros são usados para contornar esse problema

Para definir um ponteiro é usado o `*`

```go
    x := 5
    y := &x
    *y = 10
```

No trecho de codigo acima acontece o seguinte:

- variavel `x` é criada com valor 5
- `y` recebe o endereço de memoria da variavel `x`
- `y` é definido como ponteiro e compartilha o mesmo endereço de memoria de `x`
- por fim `x` e `y` tem o valor 10 po compartilharem o mesmo endereço de memoria
