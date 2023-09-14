package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando o usuário %s !!!\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade > 17
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Cícero 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())

	usuario2 := usuario{"Yuri", 9}
	fmt.Println(usuario2)
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade())
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
