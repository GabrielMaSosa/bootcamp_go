/*

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
 Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.


*/

package main

import (
	"errors"
	"fmt"
)

func calcPromedio(notas ...float64) (float64, error) {
	acum := 0.0

	for _, nota := range notas {
		if nota < 0.0 {
			return 0.0, errors.New("No existen notas negativas")
		} else {
			acum += nota
		}

	}

	// va todo bien tambien hacemos casting del len
	return (acum / float64(len(notas))), nil
}

func main() {
	prom, err := calcPromedio(3, 2, 6, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio de las notas es ", prom)

	}

}
