package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gthub.com/GabrielMaSosa/test-funcional/cmd/handlers"
	product "gthub.com/GabrielMaSosa/test-funcional/internal/protuct"
	"gthub.com/GabrielMaSosa/test-funcional/pkg"
)

func main() {

	// inyectamos las dependencias
	slice, err := pkg.InitilizeBD("../products.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(slice))
	fmt.Println(slice)
	repo := product.NewRepository(slice)
	servi := product.NewServiceProduct(&repo)
	hdler := handlers.NewHandlerProduct(servi)
	//fin de las inyecciones

	//inicio server
	server := gin.Default()
	productsrout := server.Group("/products")
	handlers.Rutas(productsrout, hdler)

	server.Run(":8080")

}
