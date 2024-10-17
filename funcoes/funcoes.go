package main

import "fmt"

func init() {
	fmt.Println("Chama a primeira função")
}

func main() {
	soma, subtracao, multiplicacao, divisao := Operacao(3,4)

	fmt.Println("Soma = ", soma, "Subtração = ", subtracao, "Mutiplcação = ", multiplicacao, "Divisão = ", divisao)
}

func Operacao(num1 int, num2 int) (soma int, subtracao int, multiplicacao int, divicao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	multiplicacao = num1 * num2
	divicao = num1 / num2

	return 
}