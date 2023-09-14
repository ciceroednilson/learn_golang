package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	default:
		return "Sábado"

	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough //joga para a próxima condição
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	default:
		diaDaSemana = "Sábado"
	}
	return diaDaSemana
}

func main() {
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(6))
	fmt.Println(diaDaSemana(7))
	fmt.Println("=============diaDaSemana2=============")
	fmt.Println(diaDaSemana2(1))
	fmt.Println(diaDaSemana2(7))
}
