package main

import (
	"fmt"
	"os"
)

func show() {
	err := recover()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ejecución finalizada")
}
func main() {
	defer show()
	archivo, err := os.Open("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	data, err2 := os.ReadFile("customers.txt")
	if err2 != nil {
		panic("No se puede abrir el archivo")
	}
	fmt.Printf("%s", data)
	defer archivo.Close()
	fmt.Println(archivo)

}
