package main

import "fmt"

func main() {

	func() {
		fmt.Println("Função Anonima")
	}()

	func(nome string) {
		fmt.Println(nome)
	}("Cícero")

	var retorno = func(nome string) string {
		return fmt.Sprint("O nome é ", nome)
	}("Cícero 2")
	fmt.Println(retorno)

}
