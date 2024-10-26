package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2, distance float64

	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)

	distance = math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2)

	fmt.Printf("%.4f\n", math.Sqrt(distance))

}
