package main

import (
	"fmt"
)

func main() {
	dtain := "HOLA MUNDOJJ"

	//sacamos todo asci que no sea letra
	contador := 0
	for _, v := range dtain {

		switch {
		case v >= 65 && v <= 90:
			fmt.Println(v, string(v))
			contador++
		case v >= 97 && v <= 122:
			fmt.Println(v, string(v))
			contador++

		}

	}
	fmt.Println("La palabra", dtain, "continen", contador, "Letras")

}
