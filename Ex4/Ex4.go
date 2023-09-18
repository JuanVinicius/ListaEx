package main

import (
	"fmt"
)

func main() {
	var nt1, nt2, nt3, nt4 float64

	fmt.Println("Insira abaixo suas notas para o calculo da média \nPrimeira nota: ")
	fmt.Scanf("%g", &nt1)

	fmt.Println("Segunda nota: ")
	fmt.Scanf("%g", &nt2)

	fmt.Println("Terceira nota: ")
	fmt.Scanf("%g", &nt3)

	fmt.Println("Quarta nota: ")
	fmt.Scanf("%g", &nt4)

	Media := (nt1 + nt2 + nt3 + nt4) / 4

	fmt.Println("\nSua média é de: ", Media)
}
