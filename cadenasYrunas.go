//Una string en Go es un slice de bytes con permiso read-only. Tanto el lenguaje como las librerias estandares tratan las strings como contenedores de texto codificadas en UTF-8.
//Mientras en otros le guajes una string esta hecha de "caracteres", en Go el concepto de caracter es llamado una "runa", un entero que representa un puntero a un codigo Unicode

package main

import (
	"fmt"
	"unicode/utf8"
)

func cadenasrunas() {

	const s = "Lionel Messi"

	fmt.Println("Longitud de la cadena: ", len(s)) //como las strings son arreglos de bytes, la longityd de esta es la longitud de los bytes almacenados en este arreglo

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println() //Indexando la string obtenemos el valor del byte en cada indice. Luego, mediante este bucle obtenemos el valor en hexadecimal de cada byte al cual este apuntando el puntero

	fmt.Println("Contador de runas con RuneCountInString: ", utf8.RuneCountInString(s)) //Para contar cuantas runas hay en una string, usamos el paquete utf8. El tiempo de ejecucion de la funcion RuneCountInString depende del tamaño de la string, ya que tiene que decodificar una por una todas las runas UTF8
	for idx, runeValue := range s {
		fmt.Printf("%#U Empieza en %d\n", runeValue, idx)
	}

	fmt.Println("\nUsando DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U Empieza en %d\n", runeValue, i)
		w = width

		examineRune(runeValue) //pasaje de una runa en una funcion
	}

}

func examineRune(r rune) {

	switch r {
	case 't':
		fmt.Println("Se encuentra el nombre: ")
	case 'ส':
		fmt.Println("El apellido es: ")
	} //Cuando una runa es indicada entre comillas simples, se pasa su "literal", el cual puede ser comparado con un "valor"
}
