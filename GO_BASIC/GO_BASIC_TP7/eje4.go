package main

import "fmt"

type Myerror struct{}

func (err *Myerror) MError(val int) error {

	return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d ", val)

}

func main() {

	salary := 4000
	err := &Myerror{}

	if salary <= 10000 {
		fmt.Println(err.MError(salary))

	} else {

		fmt.Println("OK")
	}

}
