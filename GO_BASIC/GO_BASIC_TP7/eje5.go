package main

import (
	"errors"
	"fmt"
)

func CalcSueldo(horas int, pphora float64) (salario float64, err error) {
	salario = float64(horas) * (pphora)

	switch {

	case salario >= 150000.0 && horas >= 80:
		salario = salario * 0.9
		err = nil
		return
	case salario < 150000.0 && horas >= 80:
		err = nil
		return
	case horas < 80:
		salario = 0.0
		err = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	default:
		salario = 0.0
		err = nil
		return

	}
}

func main() {
	horas := 70
	precio := 4500.0
	salario, err := CalcSueldo(horas, precio)
	if err != nil {
		fmt.Println(err)
		return //siempre devuelvo return para que termine el programa
	} else {
		fmt.Println("El salario es de :", salario)
	}

}
