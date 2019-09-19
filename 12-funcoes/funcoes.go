package main

import "fmt"

func main() {
	fmt.Println(vai("foi"))
}

func vai(texto string) (saida1, saida2 string) {
	return texto + ": 1", texto + ": 2"
}
