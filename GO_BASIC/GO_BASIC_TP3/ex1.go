/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y
si gana más de $150.000 se le descontará además un 10 % (27% en total)


*/

package main

import "fmt"

func calcImpuesto(val float64) float64 {

	switch {
	case val > 150000.0:
		return val * (0.17 + 0.10)
	case val > 50000.0 && val <= 150000.0:
		return val * 0.17

	default:
		return 0.0

	}

}

func main() {

	sueldo := 3000.0

	fmt.Println("EL impuesto del sueldo ", sueldo, "es:", calcImpuesto(sueldo))

}
