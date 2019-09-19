package main

import "fmt"

func main() {

	//*******************************
	//declaração de variáveis

	var numeroInteiro int = 42
	var numeroPontoFlutuante float64 = 42.1
	//numeroPontoFlutuante = numeroInteiro //erro de compilação
	numeroPontoFlutuante = float64(numeroInteiro) //ok

	//*******************************	
	//declaração de variável com tipo implícito
	
	//numeroInteiro := 42 //erro de compilação, numeroInteiro já declarado
	numeroInteiro2 := 42
	umaString := "quarenta_e_dois"

	//*******************************	
	//múltiplas declarações

	var i, j, k int = 1, 1, 1

	//*******************************
	//Constantes

	const imutavel = "imutavel"
	//imutavel = "alterado" //erro de compilação

	fmt.Println(numeroInteiro, numeroInteiro2, numeroPontoFlutuante, umaString, i, j, k)
}
