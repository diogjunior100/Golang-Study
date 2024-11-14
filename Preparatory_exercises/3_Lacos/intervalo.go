package main

import (
	"fmt"
)

func main () {
	var n, value int
	var in int = 0
	var out int = 0

	fmt.Scan(&n)

	for n != 0 {
		fmt.Scan(&value)

		if value >= 10 && value <= 20 {
			in = in + 1
		} else {
			out = out + 1
		}

		n--
	}

	fmt.Printf("%d in\n", in)
	fmt.Printf("%d out\n", out)
}
