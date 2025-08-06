package main

import (
	"fmt"
)

func rango() { //"range" itera sobre una variedad de elementos y estructuras de datos
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	} //Esta funcion suma los elementos del slice, ignorando los indices
	fmt.Println("La suma de los elementos del slice es: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	} //"range" en arrays y slices nos brinda tanto el indice como el valor para cada entrada

	mapa1 := map[string]string{"a": "azul", "b": "blanco", "c": "celeste"}
	for k, v := range mapa1 {
		fmt.Printf("%s -> %s \n", k, v)
	}
	for k := range mapa1 {
		fmt.Println("Clave: ", k) //range tambien puede iterar sobre las claves de un mapa
	}
	for i, c := range "go" {
		fmt.Println(i, c) //range sobre un String recorre sobre cada caracter, por lo que el i tomara el valor de la posicion que ocupe el caracter sobre la cadena, mientras que c muestra su codigo en Unicode (runa)
	}

}
