package main

import "fmt"

type Incrementavel interface {
	Incrementa(incremento int)
}

type Ponto struct {
	X, Y int
}

func (p *Ponto) Incrementa(incremento int) {
	p.X += incremento //atalho para (*p).X += incremento
	p.Y += incremento
}

func main() {

	var ponto Incrementavel
	//compila mas SIGSEG
	//função é chamada com valor nulo, não é um NullPointer
	//ponto.Incrementa(10)

	ponto = &Ponto{2, 4}
	fmt.Println(ponto)

	ponto.Incrementa(10)
	fmt.Println(ponto)

}
