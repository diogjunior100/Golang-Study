package main

import (
    "fmt"
)

func main() {

    var x,y,z, comparador int

    fmt.Scan(&x, &y, &z)

    a := x
    b:= y
    c := z

    if a < b {
        comparador = a
        a = b
        b = comparador
    }

    if b < c {
        comparador = b
        b = c
        c = comparador
    }

    if a < b {
        comparador = a
        a = b
        b = comparador
    }

    fmt.Printf("%d\n%d\n%d\n", c,b,a)
    fmt.Println()
    fmt.Printf("%d\n%d\n%d\n", x,y,z)
}
