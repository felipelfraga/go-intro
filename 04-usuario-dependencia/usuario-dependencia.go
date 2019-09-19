package main

import (
	"github.com/felipelfraga/go-intro/03-compartilhado/compartilhado"
	"fmt"
)

func main() {
	compartilhado.Imprime("um_texto_qualquer")
	fmt.Println(compartilhado.AtributoPublico)
	
	//erro de compilação. atributoPrivado é privado
	//fmt.Println(compartilhado.atributoPrivado)
}