package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32}

func main() {

	for i := range pow {
		fmt.Printf("Indice: %d\n", i)
	}

	for i, v := range pow {
		fmt.Printf("Indice: %d Valor: %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("Valor: %d\n", v)
	}

	for i, _ := range pow {
		fmt.Printf("Indice %d\n", i)
	}
}
