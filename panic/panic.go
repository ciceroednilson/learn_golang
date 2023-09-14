package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado")
	}

}

func testePanic(numero int) string {
	defer recuperar()
	if numero == 1 {
		return "Cicero"
	} else if numero == 2 {
		return "Yuri"
	}
	panic("Não sei o que fazer")

}

func main() {
	fmt.Println(testePanic(1))
	fmt.Println(testePanic(2))
	fmt.Println("Antes do panic")
	fmt.Println(testePanic(3))
	fmt.Println("Depois do panic")
}
