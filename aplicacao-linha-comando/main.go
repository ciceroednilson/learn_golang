package main

import (
	"fmt"
	"linha_de_comando/application"
	"log"
	"os"
)

func main() {
	// exemplo para executar
	// go run main.go ip --host amazon.com.br
	// go run main.go servidores --host bing.com.br
	// go build
	// ./linha_de_comando ip --host yahoo.com.br
	app := application.Gerar()
	fmt.Println(app.Name)
	error := app.Run(os.Args)
	if error != nil {
		log.Fatal("Erro fatal")
	}

}
