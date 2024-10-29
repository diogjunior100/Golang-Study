package main

import (
	"fmt"
)

func main(){
	var n, a, l, p int

	fmt.Scanln(&n)
	fmt.Scan(&a,&l,&p)

	if n <= a && n <= l && n <= p{
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}

}