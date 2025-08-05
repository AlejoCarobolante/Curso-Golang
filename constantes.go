package main

import (
	"fmt"
	"math"
)

func constantes() {
	const g string = "constant"
	fmt.Println(g)

	const h = 15000
	const i = 60000 / h

	fmt.Println(i)

	fmt.Println(math.Sin(i)) //la librería math sirve para hacer operaciones avanzadas
}
