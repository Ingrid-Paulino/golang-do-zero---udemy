package main

import "fmt"

func main() {
	// funcao anonima nao tem nome
	func() string {
		return fmt.Sprintf("Funcao anonima")
	}() // assim que ela rodar ela ja vai executar

	//ou

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
