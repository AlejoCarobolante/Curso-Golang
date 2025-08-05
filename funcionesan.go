package main

func secuenciaEnt() func() int { //El uso de funciones anonimas se da cuando se quiere definir una funcion inline sin tener que nombrarla
	i := 0
	return func() int {
		i++
		return i
	}

} //la funcion secuenciaEnt devuelve otra funcion, la cual es definida anonimamente en el cuerpo de la primera. La funcion retornada se cierra sobre la variable i para formar un cierre propiamente dicho.
