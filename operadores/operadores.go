package main

import "fmt"

func main() {
	var soma int32 = 1 + 2
	var subtracao int32 = 10 - 5
	var divisao int32 = 10 / 2
	var multiplicacao int32 = 2 * 5
	var restoDivisao int32 = 11 % 2
	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(divisao)
	fmt.Println(multiplicacao)
	fmt.Println(restoDivisao)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	var verdadeiro bool = true
	var falso bool = false
	fmt.Println(verdadeiro && true)
	fmt.Println(falso && false)
	fmt.Println(falso || verdadeiro)

	numero := 10
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)

	var texto string
	if numero > 5 {
		texto = "Maior"
	} else {
		texto = "Menor"
	}
	fmt.Println(texto)
}
