package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y  = 10
	fmt.Println("Main=======")
	fmt.Println(&x, *y)
	fmt.Println("ImprimirValores=======")
	ImprimirValores(&x, y)
}

func ImprimirValores(x *int, y *int) {
	fmt.Println(x, y)
	fmt.Println(&x, &y)
}