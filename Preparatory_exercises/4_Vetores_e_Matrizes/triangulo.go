package main

import (
	"fmt"
	"sort"
)
 
func main() {

	var vtr[4] int

	for i := 0; i < 4; i ++ {
		fmt.Scan(&vtr[i])
	}

	sort.Ints(vtr[:])

	


}