package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	url := "/ping"
	router.GET(url, func(c *gin.Context) {
		c.String(200, "PONG")
	})
	router.Run(":8080")

}
