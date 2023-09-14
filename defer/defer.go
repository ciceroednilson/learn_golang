package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Fim da nota")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	//defer funcao1()
	funcao2()
	alunoEstaAprovado(7, 8)
}
