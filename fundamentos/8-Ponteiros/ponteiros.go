package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1 // copia

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int // valor padrao é 0
	var ponteiro *int // valor padrao é nil // no ponteiro faz referencia do endereço de memoria

	variavel3 = 100
	ponteiro = &variavel3 // estou jogando o endereco de memoria "&variavel3" dentro do "ponteiro"
	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // para eu ver o valor do ponteiro eu coloco o *

}
