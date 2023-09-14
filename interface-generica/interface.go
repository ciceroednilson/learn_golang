package main

import "fmt"

func generico(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generico("Azul")
	generico(true)
	generico(1)
}
