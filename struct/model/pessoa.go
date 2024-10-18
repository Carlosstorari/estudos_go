package model

import "time"

type Pessoa struct {
	Nome string
	Endereco Endereco
	DataDeNascimento time.Time
}

func (p Pessoa)IdadeAtual() int { // metodo dentro da struct
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}