package main

import (
    "fmt"
)

func main() {

    var raio float64
    var area float64

    fmt.Scanln(&raio)

    area = (raio*raio) * 3.14159

    fmt.Printf("A=%.4f\n", area)
}