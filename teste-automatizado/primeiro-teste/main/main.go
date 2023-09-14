package main

import (
	"fmt"
	"unit-test-mod/primeiro-teste/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Avenida central")
	fmt.Println(tipoDeEndereco)
}
