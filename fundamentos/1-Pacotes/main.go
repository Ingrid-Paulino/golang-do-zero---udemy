package main

import (
	"fmt"                         //pacote interno
	"github.com/badoux/checkmail" // pacote externo
	"modulo/auxiliar"             //pacote interno
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("email@gmail.com") //retorna nil/null se estiver tudo certp
	fmt.Println(err)
}
