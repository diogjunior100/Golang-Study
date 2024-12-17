package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	vtr := make([]int, n)

	for i:= 0; i < n; i++ {
		fmt.Scan(&vtr[i])
	}

	menorValor := vtr[0]
	indice := 0

	for i := 1; i < n; i++ {
		if vtr[i] < menorValor {
			menorValor = vtr[i]
			indice = i
		}
	}

	fmt.Printf("Menor valor: %d\n", menorValor)
	fmt.Printf("Posicao: %d\n", indice)

}
