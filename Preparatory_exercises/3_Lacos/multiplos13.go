package main

import (
	"fmt"
)

func main() {

	var x int
	var y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	sum := 0

	if x < y {
		for i:=x; i <=y; i++ {
			if i % 13 != 0 {
				sum += i
			}
		}
	} else {
		for i:=y; i <=x; i++ {
			if i % 13 != 0 {
				sum += i
			}
		}
	}

	fmt.Println(sum)

}
