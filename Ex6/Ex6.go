package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Insira o numero para testar se é multiplo de 5: ")
	fmt.Scanf("%d", &num)

	if num%5 == 0 {
		fmt.Printf("É multiplo!")
	} else {
		fmt.Println("Não é multiplo!")
	}

}
