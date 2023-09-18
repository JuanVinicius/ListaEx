package main

import (
	"fmt"
)

func main() {
	var comp float32
	var larg float32
	fmt.Println("Insira abaixo as dimensões de largura e comprimento para o calculo.\nIniciando com a medida de largura, digite abaixo:")
	fmt.Scanf("%g", &comp)

	fmt.Println("Insira a largura do terreno: ")
	fmt.Scanf("%g", &larg)

	areaTerreno := comp * larg

	fmt.Printf("Esta é a area do terreno: %g \n", areaTerreno)

}
