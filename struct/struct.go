package main

import (
	"fmt"
	model "golangestudos/struct/model"
	"time"
)


func main() {
	endereco := model.Endereco {
		Rua: "Rua X",
		Numero: 15,
		Cidade: "Campinas",
	}

	pessoa := model.Pessoa {
		Nome: "Carlos",
		Endereco: endereco,
		DataDeNascimento: time.Date(1995, 8, 26, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(pessoa)
	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)
}