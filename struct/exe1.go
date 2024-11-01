package main

import "time"

/**
Criar um modelo que irá receber os itens para
a compra do mês, nesse modelo teremos a data que a
compra irá acontecer, mercado e os itens para comprar
**/

type Compra struct {
	Data time.Time
	Mercado string
	Itens [] ItemCompra
}

type ItemCompra struct {
	nome string
}

func (c* Compra)InicializaCompra() {
	c.Data = time.Now()
	c.Mercado = "Qualquer merda"
	c.Itens = 
}