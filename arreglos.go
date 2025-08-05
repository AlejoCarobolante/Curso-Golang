package main

import "fmt"

func arreglos() { //Un arreglo es una secuencia numerada de elementos con una longitud definida, pero son utiles en escenarios muy especificos. En Go generalmente se usan "slices"

	var a [5]int //Creamos un arreglo que va a tener 5 ints. Tanto tipo de elemento como tamaño del arreglo son parte del tipo del arreglo. Por defecto, estos son "zero-valued", lo cual para ints significa que esta lleno de 0.
	fmt.Println("vacio: ", a)

	a[4] = 100
	fmt.Println("set: ", a)    //podemos setear un valor en cierta posicion con la sintaxis arreglo[indice] = valor
	fmt.Println("get: ", a[4]) //podemos obbtener el valor de un arreglo en cierta posicion con la sintaxis arreglo[indice]

	fmt.Println("longitud: ", len(a)) //la funcion len nos devuelve la longitud de un array que le pasemos como parametro

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arreglo b (tamaño definido): ", b) //podemos declarar e inicializar un arreglo en la misma linea

	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("arreglo c (tamaño compilado): ", c) //podemos dejar que el compilador cuente el numero de elementos del arreglo con "..."

	d := [...]int{100, 3: 400, 500}
	fmt.Println("arreglo d (indices): ", d) // si se especifica un indice, todos los elementos desde el ultimo definido hasta el indice definido van a ser rellenados con 0.

	var twoD [2][3]int //composicion de arreglos para definir estructuras de datos multidimensionales

	for i := range 2 { //se define el indice para recorrer el arreglo exterior
		for j := range 3 {
			twoD[i][j] = i + j
		} //se define el indice para recorrer los arreglos interiores
	}
	fmt.Println("Array 2d: ", twoD)

	definedTwoD := [2][3]int{ //Se puede definir e inicializar arreglos multidimensionales al mismo tiempo
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println("Array 2d definido e inicializado: ", definedTwoD)

}
