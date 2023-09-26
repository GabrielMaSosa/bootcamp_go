package main

import (
	"eje1tp3/cmd/handlers"
	"eje1tp3/cmd/productos"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {

	jsonfile, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}
	defer jsonfile.Close()
	bytevalues, err2 := io.ReadAll(jsonfile)
	if err != nil {
		panic(err2)
	}

	json.Unmarshal(bytevalues, &productos.Productos)

	//fmt.Printf("Productos : %+v", productos)
	fmt.Println(len(productos.Productos))

	r := gin.Default()

	r.POST("/", handlers.Save())
	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productos.Productos)
		return

	})

	r.GET("/products/:id", handlers.GetByIdPatchUrl())

	r.GET("/products/search", handlers.GetByPriceMayor())
	r.Run(":8080")
}
