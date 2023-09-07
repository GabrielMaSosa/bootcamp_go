package main

import (
	"errors"
	"fmt"
)

type Clientes struct {
	Legajo    string
	Nombre    string
	DNI       int64
	Telefono  int64
	Domicilio string
}

var MisClientes = []Clientes{
	{
		Legajo:    "123",
		Nombre:    "Carlos",
		DNI:       123,
		Telefono:  123,
		Domicilio: "Av Libertad",
	},
	{
		Legajo:    "456",
		Nombre:    "Pedro",
		DNI:       456,
		Telefono:  456,
		Domicilio: "Av Sarmiento",
	},
	{
		Legajo:    "789",
		Nombre:    "Raul",
		DNI:       789,
		Telefono:  789,
		Domicilio: "Av Corrientes",
	},
}

var (
	ErrIsEqual = errors.New("error: el cliente ya existe")
	ErrIsNull  = errors.New("error: Hay campos Nulo")
)

func (val *Clientes) IsDifferent() (flag bool, err error) {
	fmt.Println(len(MisClientes))

	for i, v := range MisClientes {
		fmt.Println(i, v)
		if *val == v {
			fmt.Println("iguales", v, *val)
			flag = false
			err = fmt.Errorf("Error val igual %w", ErrIsEqual)

		} else {
			flag = true
			err = nil
		}

	}
	return
}

func (val *Clientes) IsNotNull() (flag bool, err error) {
	if (*val).DNI != 0 && (*val).Telefono != 0 && (*val).Domicilio != "" && (*val).Legajo != "" && (*val).Nombre != "" {
		flag = true
		err = nil
	} else {
		flag = false
		err = fmt.Errorf("Error val Nulo %w", ErrIsNull)
	}
	return
}
func main() {
	/* 	val := Clientes{
		Legajo:    "123",
		Nombre:    "Gabriel",
		DNI:       456,
		Telefono:  789,
		Domicilio: "Av Libertad",
	} */

	/* val := Clientes{
		Legajo:    "",
		Nombre:    "Gabriel",
		DNI:       456,
		Telefono:  789,
		Domicilio: "Av Libertad",
	} */

	val := &Clientes{Legajo: "789",
		Nombre:    "Raul",
		DNI:       789,
		Telefono:  789,
		Domicilio: "Av Corrientes",
	}

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Fin de la ejecución")
	}()

	stat, err := val.IsDifferent()
	fmt.Println(errors.Is(err, ErrIsEqual))
	if err != nil {
		defer func() {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")

		}()
		panic(err)
	} else {
		if stat {
			fmt.Println("OKOKOK")
			fmt.Println(MisClientes)
		}
	}
	stat2, err2 := val.IsNotNull()
	if err2 != nil {
		defer func() {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")

		}()
		panic(err2)
	} else {
		if stat2 {
			fmt.Println("OKOKOK")
			fmt.Println(MisClientes)

		}

	}

}
