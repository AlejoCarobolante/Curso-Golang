package main

func valor(val int) {
	val = 0
} //esta funcion tiene como parametro un int, por lo que los argumentos seran pasados por valor. Valor va a tomar una copia de val distinta de la que llama la funcion

func puntero(pnt *int) {
	*pnt = 0
} //esta funcion recibe como parametro un *int (puntero en formato int), por lo que *pnt en el body de la funcion busca el valor en esa direccion. Esto modifica el valor del puntero indicado y lo reemplaza por 0
