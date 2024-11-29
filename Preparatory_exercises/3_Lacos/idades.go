package main

import (
	"fmt"
)

func main() {

	var (
		sum float64 = 0
		qtd float64 = 0
		x float64 = 1
		media float64
	)

	for x > 0 {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		sum = sum + x
		qtd = qtd + 1
	}

	media = sum / qtd

	fmt.Printf("%.2f", media)

}


