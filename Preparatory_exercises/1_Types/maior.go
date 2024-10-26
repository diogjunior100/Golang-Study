package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, maior float64

	fmt.Scanln(&a, &b, &c)

	maior = ((a + b) + math.Abs(a-b)) / 2

	if maior > c {
		fmt.Printf("%.0f eh o maior\n", maior)
	} else {
		fmt.Printf("%.0f eh o maior\n", c)
	}
}
