package main

import (
	"errors"
	"fmt"
)

type Myerror struct{}

func (err *Myerror) Myerror() error {
	return errors.New("Error: el salario es menor a 10.000")

}

func main() {
	salary := 4000
	err := &Myerror{}

	if salary <= 10000 {
		fmt.Println(err.Myerror())

	} else {

		fmt.Println("OK")
	}

}
