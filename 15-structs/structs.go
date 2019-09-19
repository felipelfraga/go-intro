package main

import "fmt"

type Ponto struct {
	X, Y int
}

func main() {

	ponto := Ponto{1, 2}
	fmt.Println(ponto)

}
