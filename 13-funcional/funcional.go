package main

import "fmt"

func main() {
	s1, s2 := biProdutor("um_texto_qualquer")
	fmt.Println(s1, s2)

	s1, s2 = biProdutorIndireto(biProdutor)
	fmt.Println(s1, s2)

	s1, s2 = biProdutorIndireto(func(texto string)(s1, s2 string){
		return texto + "_lambda1", texto + "_lambda2"
	})
	fmt.Println(s1, s2)

}

func biProdutorIndireto(funcaoComoValor func (string)(string, string)) (string, string) {
	return funcaoComoValor("funcao_como_valor")
}

func biProdutor(texto string) (s1, s2 string) {
	return texto + "1", texto + "2"
}
