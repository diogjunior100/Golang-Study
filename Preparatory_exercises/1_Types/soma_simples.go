package main
 
import (
    "fmt"
)
 
func main() {

    var x int
    var y int
    
    fmt.Scanln(&x);
    fmt.Scanln(&y);
    
    sum := x + y
    
    fmt.Printf("SOMA = %d\n", sum)

}
