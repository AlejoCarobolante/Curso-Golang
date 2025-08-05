package main

import "fmt"

func variables() {

	var a = "initial"
	fmt.Println(a) //define una variable llamada "a" de tipo string (aunque no se indica se define por el tipo de datos asignado)

	var b, c int = 1, 2
	fmt.Println(b, c) //define dos variables "b y c" de tipo int, se asignan y se imprimen sus valores

	var d = true
	fmt.Println(d) //define una variable d, un booleano con valor true

	var e int
	fmt.Println(e) //se define un int e, pero al no asignarle un valor se inicializa por default en 0

	f := "apple"
	fmt.Println(f) //define una variable f (la asignación := implica no indicar un tipo de dato, sino que la variable es del tipo del valor asignado). Esta sintaxis sólo es válida dentro de funciones
}
