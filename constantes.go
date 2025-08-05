package main

import (
	"fmt"
	"math"
)

func constantes() {
	const g string = "constant"
	fmt.Println(g) //defino una string con valor "constant" y la muestro

	const h = 15000
	const i = 60000 / h //defino dos constantes, de las cuales una depende de la otra

	fmt.Println(i)

	fmt.Println(math.Sin(i)) //la librería math sirve para hacer operaciones avanzadas
}
