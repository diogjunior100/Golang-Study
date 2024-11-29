package main

import (
"fmt"
)

func main() {

	var t int;
	var (
		pa int
		pb int
		g1 float64
		g2 float64
		years int
	)

	fmt.Scan(&t)

	for t > 0 {

		fmt.Scan(&pa, &pb, &g1, &g2)
		years = 0

		for pa <= pb {
			pa += int(float64(pa) * (g1 / 100.0))
			pb += int(float64(pb) * (g2 / 100.0))
			years++

			if years > 100 {
				fmt.Println("Mais de 1 seculo.")
				break
			}
		}

		if years <= 100 {
			fmt.Printf("%d anos.\n", years)
		}

		t--
	}
}

