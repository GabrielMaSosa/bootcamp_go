package main

import "fmt"

//Ejercicio 1 - Registro de estudiantes
//muestra los valores y tambien adicionalmente hice para agregar a un array y que muestre el array

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      uint64
	Fecha    string
}

func (data *Alumno) Detalle() {
	fmt.Println("Address Memory: ", &data)
	fmt.Println("Nombre: ", data.Nombre)
	fmt.Println("Apellido: ", data.Apellido)
	fmt.Println("DNI: ", data.DNI)
	fmt.Println("Fecha de Nac: ", data.Fecha)
}

func (data *Alumno) DetalleArray(vect *[]Alumno) {

	for _, v := range *vect {

		fmt.Println("Address Memory: ", &v)
		fmt.Println("Nombre: ", v.Nombre)
		fmt.Println("Apellido: ", v.Apellido)
		fmt.Println("DNI: ", v.DNI)
		fmt.Println("Fecha de Nac: ", v.Fecha)
	}
}

func (data *Alumno) Save(additem Alumno, vect *[]Alumno) {
	(*vect) = append((*vect), additem)

}

func main() {
	var p Alumno
	var arr []Alumno
	fmt.Println(&p)
	p.Apellido = "Sosa"
	p.Nombre = "Gabriel"
	p.DNI = 35005287
	p.Fecha = "30/07/1991"

	p.Detalle()
	p.Save(p, &arr)
	p.DetalleArray(&arr)
	p.Apellido = "Pedrosky"
	p.Nombre = "Raul"
	p.DNI = 35005287
	p.Fecha = "22/07/1991"
	p.Detalle()
	p.Save(p, &arr)
	p.DetalleArray(&arr)

}
