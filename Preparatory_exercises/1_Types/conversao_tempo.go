package main

import "fmt"

func main() {
	var segundos int

	fmt.Scanln(&segundos)

	horas := segundos / 3600
	minutos := (segundos % 3600) / 60
	tempo_segundos := segundos % 60

	fmt.Printf("%d:%d:%d\n", horas, minutos, tempo_segundos)

}
