package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematico(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	var f = func(texto string) string {
		fmt.Println("Funcao  F" + texto)
		return texto
	}
	resultado := f(" De novo")
	fmt.Println(resultado)

	resultoSoma, resultadoSubstracao := calculosMatematico(12, 6)
	fmt.Println(resultoSoma)
	fmt.Println(resultadoSubstracao)
	fmt.Println("============")
	resultoSoma2, _ := calculosMatematico(12, 6)
	fmt.Println(resultoSoma2)
}
