package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS INTERIOS
	var numeroPadrao int = 10 //Por traz dos panos ele coloca qual o bits que seu computador é 32 ou 64
	fmt.Println(numeroPadrao)

	var numero int8 = -100
	fmt.Println(numero)

	var numero1 int16 = -10000
	fmt.Println(numero1)

	var numero2 int32 = 10000
	fmt.Println(numero2)

	var numero3 int64 = -100000000000
	fmt.Println(numero3)

	var numero4 uint32 = 10000 //uint é um int sem sinal,  entao se eu colocar -10000 vai dar erro. Inclusive ele aceita todos os tipos de tamanho do int
	fmt.Println(numero4)

	// alias
	// INT32 = RUNE
	var numero5 rune = 12456 // rune é um alias de int32 - são a mesma coisa
	fmt.Println(numero5)

	// BYTE = UINT8
	var numero6 byte = 123 // byte é um alias de int8
	fmt.Println(numero6)

	// NÚMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67 // A variavel declara o valor como um float e ele vai ser 32 ou 64 dependendo do computador
	fmt.Println(numeroReal3)

	// STRINGS
	var str string = "Texto" //Sempre aspas duplas para string
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A' //char é tipo int e aceita so um caracter // char sempre aspas simples e ele me retorna sempre um numero, esse numero é o valor que esta na tabela ASCII da letra //https://web.fe.up.pt/~ee96100/projecto/Tabela%20ascii.htm
	fmt.Println(char)

	// VALORES INICIAIS DOS DADOS

	var texto2 string //valor padrao é "" (string vazias)
	fmt.Println(texto2)

	var num int //valor padrao é 0
	fmt.Println(num)

	var num2 float32 //valor padrao é 0
	fmt.Println(num2)

	var booleano1 bool // valor padrao é false
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno") // erro em go tem um tipo que é error e seu valor padrao é nil(null)
	fmt.Println(erro)                           // errors.New --> pacote nativo do GO
}
