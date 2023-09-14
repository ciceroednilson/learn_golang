package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pe1 := pessoa{"Cicero", 1}
	fmt.Println(pe1)
	estud := estudante{pe1, "TI", "USP"}
	fmt.Println(estud)
	fmt.Print(estud.nome)
}
