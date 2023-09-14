package main

import "fmt"

func main() {
	bufferSize := 2
	canal := make(chan string, bufferSize)
	canal <- "Olá mundo"
	canal <- "Olá mundo 2"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
