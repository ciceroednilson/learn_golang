package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint32
}

func main() {
	fmt.Println("Struct")
	var u usuario
	u.nome = "Cícero"
	u.idade = 37
	fmt.Println(u)

	uusuario2 := usuario{"Cicero 2", 37, endereco{"Rua X", 1}}
	fmt.Println(uusuario2)

	uusuario3 := usuario{nome: "Cicero 3"}
	fmt.Println(uusuario3)

	uusuario4 := usuario{nome: "Cicero 4", idade: 2}
	fmt.Println(uusuario4)
}
