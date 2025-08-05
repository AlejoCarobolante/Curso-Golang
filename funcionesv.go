package main

import "fmt"

func sum(nums ...int) { //funcion que toma un numero arbitrario de ints como parametro
	fmt.Print(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
