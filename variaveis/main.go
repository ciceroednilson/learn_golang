package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	variavel5, variavel6 := "variavel 5", "VARIAVEL 6"

	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	const constant1 string = "Const"
	fmt.Println(constant1)
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
