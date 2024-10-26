package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scanln(&s)

	if s == "ABC" || s == "ACB" || s == "BAC" || s == "BCA" || s == "CAB" || s == "CBA" {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
