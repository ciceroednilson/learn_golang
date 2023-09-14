package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá mundo")
	auxiliar.Escrever()
	fmt.Println(checkmail.ValidateFormat("ciceroednilsongmail.com"))
	fmt.Println(checkmail.ValidateFormat("ciceroednilson@gmail.com"))
}
