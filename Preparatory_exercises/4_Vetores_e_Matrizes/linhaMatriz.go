package main

import (
	"fmt"
)

func main () {
	var l int
	var t string
	var matrix [12][12] float32
	
	fmt.Scan(&l)
	fmt.Scan(&t)

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	var sum float32 = 0

	for j := 0; j < 12; j ++ {
		sum = sum + matrix[l][j]
	}
	
	if t == "s" {
		fmt.Printf("%.1f\n", sum)
	} else {
		fmt.Printf("%.1f\n", sum/12)
	}


}