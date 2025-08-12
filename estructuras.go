//Las structs son colecciones de campos tipados, son el equivalente de los objetos en POO

package main

type persona struct { //struct persona con nombre y edad
	name string
	age  int
}

func newPerson(name string) *persona { //la funcion newPerson crea una nueva persona unicamente mediante el nombre

	p := persona{name: name}
	p.age = 42
	return &p
} //Go tiene un garbage collector: podes obtener un puntero a una variable local, y el garbage collector solo funcionara cuando no haya referencias activas a ella
