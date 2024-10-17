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