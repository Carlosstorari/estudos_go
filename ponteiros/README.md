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

#### Sintaxe Ponteiros

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

isso significa que se o valor de `x` for auterado, o de `y` também vai ser auterado

#### Endereço de memoria de ponteiro VS variavel comum

Para ver qual endereço de memoria um ponteiro esta apontando no caso do `y` por exemplo, é só escrever o nome da variavel sem o `*`.

```go
   x := 5
	y := &x
	*y  = 10
   fmt.Println(&x, y)
```

Resultado:

```
    0xc00010c040 0xc00010c040
```

Os dois resultados são iguais pois `x` e `y` compartilham o mesmo endereço de variavel, além disso no print do resultado o `&` é usado no `x` porque ele não é um ponteiro diferente do `y` que não precisa do `&`

Caso seja necessario exibir o valor de um ponteiro o `*` também é usado

```go
   x := 5
	y := &x
	*y  = 10
   fmt.Println(*y)
```

Resultado:

```
10
```

### Passando ponteiros como paramentro

Para definir ponteiros como parametro é só usar o `*` antes do tipo da variavel

```go
   func ImprimirValores(x *int, y *int){}
```

segue exemplo completo

```go
 func main() {
	x := 5
	y := &x
	*y  = 10

	ImprimirValores(&x, y)
	fmt.Println(&x, *y)
	fmt.Println(&x, y)
}

func ImprimirValores(x *int, y *int) {
	*x = 20
}
```

Resultado:

```
0xc000102040 0xc000102040
0xc000102040 20
0xc000102040 0xc000102040
```

repare que o endereço no print de dentro da função, é igual ao de fora
