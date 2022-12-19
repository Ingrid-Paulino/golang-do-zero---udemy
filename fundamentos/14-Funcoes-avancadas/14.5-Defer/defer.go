package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

// defer adia a chamada de uma funcao, deixa ela por ultimo
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!") // msg aparece por ultimo
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer funcao1() // essa funcao vai rodar por ultimo
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}
