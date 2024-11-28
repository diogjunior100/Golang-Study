package main

import (
	"fmt"
)

func is_prime (x int) bool {
	if x < 2 {
		return false
	}

	if x == 2 {
		return true
	}

	if x % 2 == 0 {
		return false
	}

	for i := 3; i*i <= x; i += 2 {
		if x % i == 0 {
			return false
		}
	}

	return true
}

func main() {

	var n int
	var x int

	fmt.Scan(&n)

	for n > 0 {
		fmt.Scan(&x)

		if is_prime(x) {
			fmt.Printf("%d eh primo\n", x)
		} else {
			fmt.Printf("%d nao eh primo\n", x)
		}

		n--
	}

}
