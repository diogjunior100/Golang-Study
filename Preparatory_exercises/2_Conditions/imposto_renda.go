package main

import (
    "fmt"
)

func main() {

    var salary, op, tax float64

    fmt.Scanln(&salary)

    if salary <= 2000.00 {
        fmt.Println("Isento")
    } else if salary > 2000.00 && salary <= 3000.00 {
        op = salary - 2000.00
        tax = 0.08 * op
        fmt.Printf("R$ %.2f\n", tax)
    } else if salary > 3000.00 && salary <= 4500.00 {
        op = (salary - 3000.00) * 0.18
        tax = (1000*0.08) + op
        fmt.Printf("R$ %.2f\n", tax)
    } else {
        op = (salary - 4500.00)*0.28
        tax = (1000*0.08) + (1500.00*0.18) + op
        fmt.Printf("R$ %.2f\n", tax)
    }

}