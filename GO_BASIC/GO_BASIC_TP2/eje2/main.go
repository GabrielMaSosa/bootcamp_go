package main

import (
	"fmt"
)

func main() {
	var (
		edad       = 22
		trabaja    = true
		antiguedad = 1
		sueldo     = 100000.0
	)
	switch {
	case edad >= 22 && trabaja == true && antiguedad >= 1 && sueldo >= 100000.0:
		fmt.Println("Tenes Prestamo sin interes")
	case edad >= 22 && trabaja == true && antiguedad >= 1 && sueldo < 100000.0:
		fmt.Println("Tenes Prestamo con una pequeña taza de interes")
	default:
		fmt.Println("Prestamo Rechazado..")
	}

}
