package main

import (
	"fmt"
	"time" //Esta libreria sirve para obtener datos de un calendario
)

func switchcase() { //Esta declaracion expresa condicionales a traves de varias ramas

	i := 2
	fmt.Println("Escribir", i, "como: ")
	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	} //Swtitch basico

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //Para separar condiciones en un mismo case, se usa coma
		fmt.Println("Es fin de semana")
	default: //se usa el caso "default" que es lo que pasaria en caso de no cumplirse el case
		fmt.Println("Es dia de semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Todavia no pasamos del mediodia")
	default:
		fmt.Println("Ya pasamos del mediodia")
	} //Un switch sin expresion funciona como un if/else logico

	whatAmI := func(i interface{}) {
		switch t := i.(type) { //se usa el type switch para comparar tipos en lugar de valores
		case bool:
			fmt.Println("Soy booleano")
		case int:
			fmt.Println("Soy un int")
		default:
			fmt.Println("No conozco el tipo", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hola")

}
