## Herança no Go

A herança em Go é feito atraves da `struct`

```go
type Automovel struct {
	Ano int
	Placa string
	Modelo string
}

type Moto struct {
	Automovel
	Cilindradas int
}
```

No exemplo acima é definido a `struct` pai `Automovel` e depois a `struct` filha `Moto`
