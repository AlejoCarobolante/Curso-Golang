package main

import "fmt"

func buclefor() { //El for es la unica manera de hacer bucles en Go

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} //defino una variable i y le asigno el valor 1, luego, mientras sea menor o igual a 3 la imprimo y luego la incremento

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	} //inicializo la variable j dentro del loop y la imprimo e incremento siempre y cuando sea menor a 3

	for i := range 3 {
		fmt.Println("range", i)
	} //defino i dentro de un rango

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
