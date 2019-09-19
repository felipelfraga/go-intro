package main

import "fmt"

type Ponto struct {
	X, Y int
}

func (p Ponto) TentaIncrementar(incremento int) {
	p.X += incremento
	p.Y += incremento
}

func (p *Ponto) Incrementa(incremento int) {
	p.X += incremento //atalho para (*p).X += incremento
	p.Y += incremento
}

func main() {

	ponto := Ponto{5, 8}
	fmt.Println(ponto)

	ponto.TentaIncrementar(2)
	fmt.Println(ponto)

	ponto.Incrementa(2)
	fmt.Println(ponto)

	ponteiroParaPonto := &ponto

	ponteiroParaPonto.Incrementa(2) //syntax sugar
	fmt.Println(ponto)

	(*ponteiroParaPonto).Incrementa(2)
	fmt.Println(ponto)

}
