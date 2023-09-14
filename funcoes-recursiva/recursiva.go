package main

import "fmt"

func fibonacci(numero uint) uint {
	if numero <= 1 {
		return numero
	}
	return fibonacci(numero-2) + fibonacci(numero-1)
}

func main() {
	fmt.Println("Recursiva")
	var posicao = fibonacci(uint(10))
	for i := uint(0); i < posicao; i++ {
		fmt.Println(i)
	}
}
