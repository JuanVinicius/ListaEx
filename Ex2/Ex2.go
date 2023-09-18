package main

import (
	"fmt"
	"time"
)

func main() {
	hora := time.Now()
	minutos := hora.Hour()*60 + hora.Minute()

	fmt.Print("Ja se passaram ", minutos, " minutos. \n\n")
}
