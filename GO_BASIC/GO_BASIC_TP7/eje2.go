package main

import (
	"errors"
	"fmt"
)

var Msg = "Error: el salario es menor a 10.000"

type Myerror struct {
}

func (dta Myerror) Error() (out string) {

	return
}

func ValidSalario(val int) (err error) {
	er2 := Myerror{}

	if val < 150000 {
		return &er2
	} else {
		return nil
	}
}

func main() {

	salary := 1800
	var err2 error

	err := ValidSalario(salary)
	err2 = &Myerror{}
	if errors.Is(err, err2) {
		fmt.Println(err2.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
