package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scanln(&x, &y)

	if x > y {
		division := x % y

		if division == 0 {
			fmt.Printf("Sao Multiplos\n")
		} else {
			fmt.Printf("Nao sao Multiplos\n")
		}

	} else {
		division := y % x

		if division == 0 {
			fmt.Printf("Sao Multiplos\n")
		} else {
			fmt.Printf("Nao sao Multiplos\n")
		}
	}
}
