/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes
y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes,
la categoría y que devuelva su salario.
*/
package main

import (
	"errors"
	"fmt"
)

func calcSalario(tipo string, minutos int) (float64, error) {
	horas := float64(minutos) / 60.0
	if minutos <= 0.0 {
		return 0.0, errors.New("NO se puede ingresar minutos negativos")
	}

	switch tipo {
	case "Categoria C":
		return horas * 1000.0, nil
	case "Categoria B":
		return horas * 1500.0 * 1.2, nil
	case "Categoria A":
		return horas * 3000.0 * 1.5, nil
	default:
		return 0.0, errors.New("Categoria Invalida")
	}

}

func main() {
	const (
		catC = "Categoria C"
		catB = "Categoria B"
		catA = "Categoria A"
	)

	salario, err := calcSalario(catA, -100)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("El salario del empleado de ", catA, "es :", int64(salario))
	}

}
