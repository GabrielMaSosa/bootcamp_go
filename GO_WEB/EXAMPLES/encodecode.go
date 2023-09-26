package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	url := "/hello/world"
	router.GET(url, func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello World"})
	})
	router.Run(":8080")

}
