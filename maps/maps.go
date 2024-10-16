package main

import "fmt"

func main() {
	mapEstatico := map[string]int{"sp": 10000000, "cg": 900000}

	fmt.Println(mapEstatico)


	
	m := make(map[string]int)

	m["sp"] = 1000000
	m["cg"] = 900000

	valor, existe := m["rj"]
	
	if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Valor nao existe")
	}
}