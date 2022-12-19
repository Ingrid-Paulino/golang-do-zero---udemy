package main

import "fmt"

// funcao variadica pode receber N parametros
// OBS: so pode ter uma variavel variadica por funcao, e ela tem que ser sempre o ultimo argumentp
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 0, 5, 7, 9, 10)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}
