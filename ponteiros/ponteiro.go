package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int16 = 1
	var variavel2 int16 = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int = 0
	var ponteiro *int
	variavel3 = 100
	ponteiro = &variavel3

	//*ponteiro desfazendo a referencia
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
