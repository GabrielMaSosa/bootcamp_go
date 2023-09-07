package main

import (
	"fmt"
)

func main() {
	meses := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio",
		7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	if inp := 10; inp <= 12 && inp >= 1 {

		for i, v := range meses {
			if i == inp {
				fmt.Println(i, ",", v)
				break //para no recorre todo rompemos el bucle
			}
		}
	} else {
		fmt.Println("MES INVALIDO debe ser de 1 al 12")
	}

}
