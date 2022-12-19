package main

import "fmt"

func main() {
	// map cria objetos
	//map[tipo das chaves]tipo dos valores
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"]) // acessa as chaves

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome") // deleta uma chave
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{ // adiciona uma chave no map
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
