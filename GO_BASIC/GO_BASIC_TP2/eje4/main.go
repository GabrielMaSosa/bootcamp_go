package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La Edad de ", "Benjamin es ", employees["Benjamin"])
	counter1 := 0
	for _, v := range employees {

		if v > 21 {
			//	fmt.Println("Mayor de ", v, ": ", k)
			counter1++
		}

	}
	fmt.Println("Cantidad de Mayores", counter1)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
