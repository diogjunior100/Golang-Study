package main

import (
    "fmt"
    "math"
)

func main() {

    var a, b, c, delta float64

    fmt.Scan(&a, &b, &c)

    delta = (b * b) - (4 * a * c)

    if a == 0 || delta < 0{
        fmt.Println("Impossivel calcular")
    } else {
        r1 := (-b + math.Sqrt(delta)) / (2 * a)
        r2 := (-b - math.Sqrt(delta)) / (2 * a)
        fmt.Printf("R1 = %.5f\n", r1)
        fmt.Printf("R2 = %.5f\n", r2)
    }

}