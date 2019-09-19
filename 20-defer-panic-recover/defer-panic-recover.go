package main

import "fmt"

func main() {
	dizAlo()
	defer dizAlo()
	defer tomaRivotrilSePreciso()
	surta()
}

func dizAlo() {
	fmt.Println("_o/")
}

func tomaRivotrilSePreciso() {
	if s := recover(); s != nil {
		fmt.Printf("Surto acalmado. Problema foi: %s\n", s)
	} else {
		fmt.Println("Zen.")
	}
}

func surta() {
	fmt.Println("Surtando!!!!!!")
	panic("ahhhhhhhhhhhhhhh...")
}
