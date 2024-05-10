package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Println("Digite um número:")
	fmt.Scanf("%d", &numero)

	if numero < 0 {
		fmt.Println("Positivo!")

	} else if numero < 0 {
		fmt.Println("Negativo!")

	} else {
		fmt.Println("Zero!")
	}
}
