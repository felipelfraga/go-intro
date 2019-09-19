package main

import "fmt"

type Incrementavel interface {
	Incrementa(incremento int) error
}

type Ponto struct {
	X, Y int
}

type ArgumentoInvalido int

func (ai ArgumentoInvalido) Error() string {
	return fmt.Sprintf("O argumento %d é inválido", ai)
}

func (p *Ponto) Incrementa(incremento int) error {
	if incremento < 1 {
		return ArgumentoInvalido(incremento)
	}
	p.X += incremento
	p.Y += incremento
	return nil
}

func main() {

	var ponto Incrementavel = &Ponto{2, 4}

	ponto.Incrementa(10)
	fmt.Println(ponto)

	erro := ponto.Incrementa(-10)
	if erro == nil {
		fmt.Println(ponto)
	} else {
		fmt.Println(ponto, erro)
	}

}
