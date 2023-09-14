package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array slices")

	var array [5]int
	array[0] = 1

	fmt.Println(array)

	arrayStr := [5]string{"Cicero", "Cicero 2", "Cicero 3", "Cicero 4", "Cicero 5"}
	fmt.Println(arrayStr)

	arrayInt := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arrayInt)

	fmt.Println("=======================")
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 10)

	fmt.Println(slice)
	fmt.Println("=======================")
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arrayInt))

	slice2 := arrayStr[1:3]
	fmt.Println(slice2)
	arrayStr[2] = "Cicero Alterado"
	fmt.Println(slice2)

	fmt.Println("===========SLICE 3============")
	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println("==========SLICE 3=============")

	fmt.Println("===========SLICE 4============")
	slice4 := make([]float32, 5)
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println("===========SLICE 4============")

}
