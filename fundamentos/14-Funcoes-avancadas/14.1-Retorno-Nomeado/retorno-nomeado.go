package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) { //cria a variavel direto na funcao
	soma = n1 + n2
	subtracao = n1 - n2
	return //ja sabe qual variavel tem que retornar de forma implicita
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
