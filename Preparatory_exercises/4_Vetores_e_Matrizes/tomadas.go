package main

import (
	"fmt"
)

func main() {

	vtr := make([]int, 5)
	sum := 0

	for i := 0; i < 4; i++ {
		fmt.Scan(&vtr[i])
		sum = sum + vtr[i]
	}

	fmt.Printf("%d\n", sum - 3)

}
