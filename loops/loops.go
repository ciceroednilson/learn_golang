package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("for 1")
		//time.Sleep(time.Second)
	}

	count := 10
	for j := 0; j < count; j++ {
		fmt.Println("for J", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"Cícero", "Machado", "Santos"}
	for i, nome := range nomes {
		fmt.Println(nome, i)
		//time.Sleep(time.Second)
	}

	for _, nome := range nomes {
		fmt.Println(nome, i)
		//time.Sleep(time.Second)
	}

	for i, letra := range "Palavra" {
		fmt.Println(string(letra), i)
	}

	pessoa := map[string]string{
		"nome":      "Cicero",
		"sobrenome": "Machado",
	}
	for key, value := range pessoa {
		fmt.Println(key, value)
	}

	for {
		fmt.Println("for true")
	}
}
