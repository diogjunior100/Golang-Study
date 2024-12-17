package main

import "fmt"

func main () {
	var n int
	fmt.Scan(&n)

	j := 0
	i := 1

	fmt.Printf("%d %d", j, i);
	n = n - 2

	for n > 0 {
		number := j + i
		j = i
		i = number

		fmt.Printf(" %d", number)
		n--
	}

	fmt.Println()
}
