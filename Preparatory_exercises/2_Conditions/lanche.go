package main

import (
	"fmt"
)

func main() {
	var code, qtd float64

	fmt.Scanln(&code, &qtd)

	switch code {
	case 1:
		fmt.Printf("Total: R$ %.2f\n", 4.00*qtd)
	case 2:
		fmt.Printf("Total: R$ %.2f\n", 4.50*qtd)
	case 3:
		fmt.Printf("Total: R$ %.2f\n", 5.00*qtd)
	case 4:
		fmt.Printf("Total: R$ %.2f\n", 2.00*qtd)
	case 5:
		fmt.Printf("Total: R$ %.2f\n", 1.50*qtd)
	}
}
