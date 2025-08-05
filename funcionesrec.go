package main

func fact(n int) int { //creamos una funcion factorial que se llama a si misma hasta llegar al caso base (0)
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(m int) int { //creamos una funcion que nos describa la serie de fibonacci para un numero elegido
	if m < 2 {
		return m
	}
	return fib(m-1) + fib(m-2)
}
