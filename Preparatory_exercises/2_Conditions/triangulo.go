package main

import (
    "fmt"
)

func main() {

    var a, b, c float64

    fmt.Scan(&a, &b, &c)

    if a < (b + c) && b < (a+c) && c < (a+b){
        perimeter := a + b + c
        fmt.Printf("Perimetro = %.1f\n", perimeter)
    } else {
        area := ((a+b)*c)/2
        fmt.Printf("Area = %.1f\n", area)
    }

}
