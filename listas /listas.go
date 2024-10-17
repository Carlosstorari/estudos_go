package main

import "fmt"

func main() {
	lista := []int{4, 9, 6, 7, 0, 5, 8, 5, 3}
	fmt.Println("Lista", lista)
	fmt.Println("List pos1", lista[0])
	fmt.Println("List pos2", lista[1])
	fmt.Println("List pos3", lista[2])
	fmt.Println("Tamanho da lista", len(lista))

	lista1 := make([]int, 1)
	lista1 = append(lista1, 2)
	
	lista2 := lista[:3] // pega os tres primeiros itens da lista 
	lista3 :=lista[4:] // pega todos os elementos apartir do quarto item 

	fmt.Println(lista2, lista3)
}