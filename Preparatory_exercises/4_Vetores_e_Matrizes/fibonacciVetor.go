package main

import (
	"fmt"
)

func main() {
	var t,n int
	fmt.Scan(&t)

	vtr := make([]int, 62)

	for i := 0; i <= 60; i++ {
		if i == 0 {
			vtr[i] = 0
		} else if i == 1 || i == 2 {
			vtr[i] = 1
		} else {
			vtr[i] = vtr[(i - 2)] + vtr[(i - 1)]
		}
	}

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		fmt.Printf("Fib(%d) = %d\n", n, vtr[n])
	}
	

}