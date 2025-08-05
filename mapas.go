package main

import (
	"fmt"
	"maps"
)

func mapas() { //los mapas son tipos asociativos (clave-valor), en otros lenguajes son llamados Hash o Dict

	m := make(map[string]int) //para crear un mapa vacio, se usa la funcion make([tipo de dato de la clave]tipo de dat del valor)

	//para setear claves se usa la sintaxis: nombre[clave] = valor

	m["k1"] = 3
	m["k2"] = 162

	fmt.Println("Mapa: ", m) //Esto imprime todos los pares clave-valor del mapa

	v1 := m["k1"]
	fmt.Println("V1: ", v1) //Obtiene el valor correspondiente a la clave "k1"

	v3 := m["k3"]
	fmt.Println("V3: ", v3) //si el elemento no existe, devuelve el 0 del tipo de dato retornado

	fmt.Println("La longitud del mapa es: ", len(m)) //cuando se aplica len() a un mapa, devuelve la cantidad de pares clave-valor

	delete(m, "k2")
	fmt.Println("Mapa modificado: ", m) //la funcion delete borra un par relacionado a la clave pasada como parametro

	clear(m)
	fmt.Println("Mapa limpio: ", m) //La funcion clear remueve todos los pares clave-valor de un mapa

	_, prs := m["k2"]
	fmt.Println("Pares: ", prs) //El valor ignorado es opcional, y nos permite verificar si la clave esta presente en el mapa, lo cual permite distinguir entre claves faltantes y claves con valores nulos. En este caso, como no es necesario, lo indicamos con _

	n := map[string]int{"a1": 1, "a2": 2}
	fmt.Println("Mapa: ", n) //Declaramos e inicializamos los valores del mapa en una linea

	n2 := map[string]int{"messi": 10, "cristiano": 7}
	if maps.Equal(n, n2) {
		fmt.Println("Ambos mapas son iguales")
	} else {
		fmt.Println("Los mapas son distinos")
	} //Igual que para slices, el package maps contiene funciones que aportan utilidades interesantes a los mapas
}
