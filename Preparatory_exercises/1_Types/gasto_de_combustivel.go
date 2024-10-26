package main

import "fmt"

func main() {
	var tempo, velocidade, distancia, litros float32

	fmt.Scanln(&tempo)
	fmt.Scanln(&velocidade)

	distancia = tempo * velocidade

	litros = distancia / 12

	fmt.Printf("%.3f\n", litros)
}
