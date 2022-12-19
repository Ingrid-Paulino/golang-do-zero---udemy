package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" //Tipo implicito  //Forma 1
	variavel2 := "Variável 2"           //Tipo explicito  //Forma 2

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( //Declara varias cariaveis ao mesmo tempo //Forma 3
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6" //Forma 4
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1" //Forma 5
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 //Inverte os valores das variaveis //Forma 6
	fmt.Println(variavel5, variavel6)
}
