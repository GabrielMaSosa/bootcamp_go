package main

import (
	"apiget/cmd/handlers"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	path := "products.json"
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(file), &handlers.ProductsList)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")

	products.GET("", handlers.GetAllProducts())
	products.GET(":id", handlers.GetProduct())
	products.GET("/search", handlers.SearchProduct())

	r.Run(":8080")
}
