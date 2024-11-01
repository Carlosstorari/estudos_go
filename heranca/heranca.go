package main

import (
	"fmt"
	model "golangestudos/heranca/model"
)

func main() {
	automovelMoto := model.Automovel {
		Ano: 2022, 
		Placa: "XPTO",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel: automovelMoto,
		Cilindradas: 125,
	} 
	fmt.Println(moto)
	fmt.Println(moto.Modelo)
}