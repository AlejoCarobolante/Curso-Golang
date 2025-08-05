package main

import (
	"fmt"
	"slices"
)

func porciones() {

	var s []string
	fmt.Println("tamaño del arreglo: ", s, s == nil, len(s) == 0) //A diferencia de los arreglos, los slices son definidos unicamente por los elementos que contiene (no su numero). Un slice sin inicializar es igual a "nil" y su longitud es 0.

	t := make([]string, 3)
	fmt.Println("vacio: ", t, "longitud: ", len(t), "capacidad: ", cap(t)) //Para crear una slice con longitud distinta de 0, se usa la funcion make([]tipo, tamaño). En este caso creamos una slice de longitud 3 de strings (inicialmente "zero-valued"). Por defecto, la capacidad de cada slice es igual a su longitud: si una slice va a crecer con el tiempo, se puede pasar el tamaño explicitamente como parametro a la funcion make() (el resto de valores no definidos son "zero-valued")

	t[0] = "a"
	t[1] = "b"
	t[2] = "c"
	fmt.Println("set: ", t)                       //se hace el set igual que en un arreglo
	fmt.Println("get: ", t[2])                    // se hace el get igual que en un arreglo
	fmt.Println("longitud de la slice: ", len(t)) //la funcion len() tambien funciona para slices.

	t = append(t, "d")
	t = append(t, "e", "f") //La funcion append nos devuelve un nuevo array con los elementos insertados al final. Su tamaño se ve modificado.
	fmt.Println("Slice con append: ", t)

	c := make([]string, len(t))
	copy(c, t)
	fmt.Println("Slice copiado: ", c) //las slices pueden ser copiadas. En este caso creamos una slice "c" del mismo tamaño que "t" y copiamos sus elementos con copy(destino, origen)

	m1 := t[2:5]
	fmt.Println("Slice recortado inferior y superiormente: ", m1) //Se pueden realizar "operaciones" a los slices con la sintaxis slice[min:max], donde el valor minimo si se incluye mientras que el maximo no.

	m2 := t[:5]
	fmt.Println("Slice recortado superiormente: ", m2) //En este caso el slice m2 nos va a mostrar todos los valores de t hasta (sin incluir) el 5to

	v := []string{"g", "h", "i"}
	fmt.Println("Slice declarado e inicializado: ", v) //Se puede declarar e inicializar un slice en la misma linea

	v2 := []string{"g", "h", "i"}
	if slices.Equal(v, v2) {
		fmt.Println("Las dos slices son iguales")
	} //Dentro del package slices podemos encontrar muchas funciones interesantes como esta que sirve para comparar

	slice2d := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1                  //definimos la longitud de los slices interiores (como el slice es de  tamaño 3, el indice va desde 0 a 2, por lo cual los slices interiores iran de 1 a 3)
		slice2d[i] = make([]int, innerLen) //creamos los slices interiores que contienen datos de tipo int y cuyo tamaño es igual a innerLen
		for j := range innerLen {
			slice2d[i][j] = i + j //para cada lugar se rellena con el valor de la suma de los subindices
		}
	}
	fmt.Println("El slice de dos dimensiones es: ", slice2d)
}
