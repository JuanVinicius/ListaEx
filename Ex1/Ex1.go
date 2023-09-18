package main

import (
	"fmt"
)

func calcSalario(x float64, y float64) int {
	var z float64 = x / y
	return int(z)
}
func main() {
	salaRecebido := 3200
	SalaMinimo := 1320

	fmt.Println(calcSalario(float64(salaRecebido), float64(SalaMinimo)), "Salarios minimos.")
}
