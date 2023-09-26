package main

import (
	"fmt"
	"time"
)

func sumar(a, b int) int {

	return a + b
}

// este es el middleware que agrega el tiempo
// recibe una funcion y retorna el mismo tipo de funcion
func mostrartiempo(f func(a, b int) int) func(a, b int) int {

	return func(a, b int) int {
		fmt.Println(time.Now())
		return f(a, b)
	}

}

//tambien podemos desencadenar middlewares

func mostrarPosterior(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		r := f(a, b)
		fmt.Println("se agrega algo mas ")
		return r
	}

}

func main() {
	fmt.Println("Resultado Middleware", mostrartiempo(sumar)(2, 3))
	fmt.Println("----------------------")
	fmt.Println("Resultado Middleware encadenado", mostrartiempo(mostrarPosterior(sumar))(2, 3))

}
