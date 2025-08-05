package main

import "fmt"

func ifelse() {

	if 7%2 == 0 {
		fmt.Println("El numero 7 es par")
	} else {
		fmt.Println("El numero 7 es impar")
	} //Ejemplo basico

	if 8%4 == 0 {
		fmt.Println("El numero 8 es divisible por 4")
	} //If sin Else

	if 8%4 == 0 || 7%4 == 0 {
		fmt.Println("Al menos uno de los siguientes numeros: (7 u 8) es divisible por 4")
	} //Utilizacion de compuertas logicas en la condicion del if

	if num := 9; num < 0 {
		fmt.Println(num, "es negativo")
	} else if num < 10 {
		fmt.Println(num, "tiene un digito")
	} else {
		fmt.Println(num, "tiene mas de un digito")
	} //If anidados, cualquier variable delcarada es accesible desde las ramas interiores
} //No existe el if ternario, incluso para las condiciones basicas hay que definir un if completo
