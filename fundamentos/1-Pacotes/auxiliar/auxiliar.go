package auxiliar

import "fmt"

func Escrever() { //Nome das funçoes precisa ser declarado a primeira letra em maiusculo - isso indica que a funçao é publica e não privada
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() //Funçoes de pacotes diferentes privadas, eu consigo chamae em arquivos diferentes mas que esteja no mesmo pacote
}
