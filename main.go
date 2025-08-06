package main

import "fmt"

func main() {

	//MOSTRAR MENSAJES
	fmt.Println("Mostrar mensaje")
	mensajes()

	//Valores
	fmt.Println("Valores")
	valores()

	//Variables
	fmt.Println("Variables")
	variables()

	//Constantes
	fmt.Println("Constantes")
	constantes()

	//Bucles for
	fmt.Println("Bucles For")
	buclefor()

	//If & Else
	fmt.Println("If y Else")
	ifelse()

	//Switch case
	fmt.Println("Switch case y default")
	switchcase()

	//Arreglos
	fmt.Println("Arreglos")
	arreglos()

	//Slices
	fmt.Println("Slices")
	porciones()

	//Mapas
	fmt.Println("Mapas")
	mapas()

	//Funciones
	fmt.Println("Funciones de un solo retorno")
	res2 := plus2(1, 2)
	fmt.Println("La suma de dos valores es: ", res2)

	res3 := plus3(1, 2, 3)
	fmt.Println("La suma de tres valores es: ", res3)

	//Funciones con retorno de valores multiples
	fmt.Println("Funciones de multiple retorno")
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b) //usamos los dos valores para solicitar el retorno
	_, c := vals()
	fmt.Println(c) //solamente solicitamos el segundo valor, ignorando el primero

	//Funciones variadicas (numero variable de argumentos)
	fmt.Println("Funciones variadicas")
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//Funciones anonimas: cierre
	fmt.Println("Funciones anonimas")
	proximoEnt := secuenciaEnt() //asignamos el valor de secuenciaEnt a proximoEnt. Esta funcion captura su propio valor de i, siendo actualizado cada vez que llamamos a proximoEnt, por lo que si lo llamamos varias veces, este valor se va a ir actualizando

	fmt.Println(proximoEnt())
	fmt.Println(proximoEnt())
	fmt.Println(proximoEnt())

	nuevoEnt := secuenciaEnt()
	fmt.Println(nuevoEnt())

	//Funciones recursivas
	fmt.Println("Funcion factorial: ", fact(7))
	fmt.Println("Serie de Fibonacci: ", fib(7))

	//Rango en estructuras
	fmt.Println("Barrido con range")
	rango()

	//Punteros
	fmt.Println("Punteros")
	i := 1
	fmt.Println("Valor inicial: ", i)
	valor(i)
	fmt.Println("Reemplazo de valor: ", i)
	puntero(&i)
	fmt.Println("Reemplazo por puntero: ", i) //la sintaxis &i nos da la direccion de memoria de i
	fmt.Println("Puntero: ", &i)              //los punteros tambien pueden imprimirse

}
