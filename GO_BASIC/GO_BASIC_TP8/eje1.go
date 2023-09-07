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
	_, err := os.Open("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}
