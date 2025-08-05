package main

func plus2(a int, b int) int { //Esta funcion toma dos ints y devuelve su suma. Go requiere returns explicitos, por lo que no hay que inferir que por operar con ints debe devolver un int.
	return a + b
}

func plus3(a, b, c int) int { //Cuando hay multiples parametros de mismo tipo de manera consecutiva, se omite el nombre de tipo para cada uno y se pone al final
	return a + b + c
}
