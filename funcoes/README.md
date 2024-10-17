# Funcoes em GO

### Fuções que retornam mais de um valor

Em Go as funções podem retornar mais de um valor, os tipos retornados são colocados entre parenteses

```go
    func Operacao(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	subtracao := num1 - num2
	multiplicacao := num1 * num2
	divicao := num1 / num2

	return soma, subtracao, multiplicacao, divicao
}
```

para pegar o valor retornado na hora de chamar a função é feito da seguinte forma

```go
    soma, subtracao, multiplicacao, divisao := Operacao(3,4)
```

#### Nomeando retornos

Os retornos da função podem ser nomeados como no exemplo abaixo

```go
    func Operacao(num1 int, num2 int) (soma int, subtracao int, multiplicacao int, divicao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	multiplicacao = num1 * num2
	divicao = num1 / num2

	return
}
```

Observe que o retorno esta vazio e na atribuição da variavel não é mais usado `:=` pois a variavel e seu tipo já foram declarados da definição da função\*\*

### Função init

A função `init` é a primeira função a ser inicializada e não precisa ser chamada em nenhum lugar, pode ser usada pra fazer uma configuração necessária antes do codigo rodar

```go
func init() {
	fmt.Println("Chama a primeira função")
}
```
