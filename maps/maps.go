package main

import "fmt"

func main() {

	fmt.Println("Maps")

	fmt.Println("Maps-------------------1")
	usuario := map[string]string{
		"nome":      "Cicero",
		"sobrenome": "Machado",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	fmt.Println("Maps-------------------2")
	usuario2 := map[string]map[string]string{
		"nome": {
			"subNome":  "Lucas 2",
			"subNome2": "Lucas 2",
		},
		"sobre": {
			"subNome":  "Luca 3",
			"subNome2": "Lucas 3",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "sobre")
	fmt.Println(usuario2)
	fmt.Println("Maps-------------------2")

	fmt.Println("Maps-------------------3")
	usuario2["employee"] = map[string]string{
		"nome_emp": "Cicero Emp",
	}
	fmt.Println(usuario2)
	fmt.Println("Maps-------------------3")
}
