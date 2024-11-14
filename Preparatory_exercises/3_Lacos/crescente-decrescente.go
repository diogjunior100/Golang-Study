package main

import (
	"fmt"
)

func main () {
	var x,y int

	fmt.Scan(&x,&y)

	for x != y {

		if x < y {
			fmt.Println("Crescente")
		} else if x > y{
			fmt.Println("Decrescente")
		}

		fmt.Scan(&x,&y)
	}

}

