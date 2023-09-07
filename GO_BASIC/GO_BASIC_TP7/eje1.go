package main

import "fmt"

type Myerror struct {
	msg string
}

func (dta *Myerror) Error() string {

	return dta.msg

}

func main() {
	salary := 1800
	e := Myerror{msg: "Error: el salario ingresado no alcanza el mínimo imponible"}

	if salary < 150000 {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
