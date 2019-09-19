package compartilhado

import "fmt"

var AtributoPublico int = 0
var atributoPrivado int = 1

//Exportado
func Imprime(s string) {
	fmt.Println(s)
}

//Privado
func imprime(s string) {
	fmt.Println(s)
}
