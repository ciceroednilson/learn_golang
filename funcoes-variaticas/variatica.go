package main

import "fmt"

func soma(numeros ...int) int {
	var soma = 0
	fmt.Println(numeros)
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func escrever() string {
	return fmt.Sprint("escrever ", soma(1, 2, 3, 4, 5, 6, 7))
}

func escrever2(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	escrever2("Escreve 2 --> ", 123)
	fmt.Println(escrever())
}
