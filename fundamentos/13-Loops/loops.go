package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		//time.Sleep(time.Second) //dorme por um segundo
	}

	fmt.Println(i)
	fmt.Println("---------------------")

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
		//time.Sleep(time.Second)
	}

	fmt.Println("---------------------")

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	fmt.Println("-----")

	for _, nome := range nomes {
		fmt.Println(nome)
	}
	fmt.Println("-----")

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra) //me traz o numero da tabela ascii
	}
	fmt.Println("-----")

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) //string converte o valor da tabela ascii para string
	}

	fmt.Println("---------------------")

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println("---------------------")

	for { //loop infinito
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second) //ctrl + c mata o processo
	}
}
