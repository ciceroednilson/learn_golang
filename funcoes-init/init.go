package main

import "fmt"

var numero int

func init() {
	fmt.Println("Primeiro init")
	numero = 10
}
func main() {
	fmt.Println("Segundo main")
	fmt.Println(numero)
}
