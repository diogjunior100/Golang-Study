package main

import (
    "fmt"
)

func main() {

    var x,y,z float64

    fmt.Scanln(&x)
    fmt.Scanln(&y)
    fmt.Scanln(&z)

    media := ((2*x) + (3*y) + (5*z)) / 10

    fmt.Printf("MEDIA = %.1f\n", media)

}