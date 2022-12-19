package main

import "fmt"

//funçãoe tbm são tipos em go
// funcão faz uma série de instruções dentro que vão ser executadas pelo seu programa

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

var f = func() { //função anonima
	fmt.Println("Função f")
}

var u = func(txt string) string {
	fmt.Println(txt)
	return txt
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { //funcao com mais de 1 retorno
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
	f()
	u("Olá")
	resultadSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadSoma, resultadoSubtracao)

	resultadSoma2, _ := calculosMatematicos(10, 15) // Ignora o segundo resultado
	fmt.Println(resultadSoma2)
}
