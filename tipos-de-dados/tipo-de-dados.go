package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 100
	fmt.Println(numero2)

	var numero3 rune = 12000
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str3 := "Texto2"
	fmt.Println(str3)

	var texto string
	fmt.Println(texto)

	var boolean1 bool = false
	fmt.Println(boolean1)

	var boolean2 bool = true
	fmt.Println(boolean2)

	var erro error = errors.New("Error new")
	fmt.Println(erro)
}
