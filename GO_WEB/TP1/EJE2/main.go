package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
}

func Saludar(ctx *gin.Context) {
	var request Persona

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.String(http.StatusOK, fmt.Sprintf("Hola %s %s", request.Nombre, request.Apellido))
		return
	}

}

func main() {
	r := gin.Default()

	r.POST("/saludo", Saludar)
	r.Run(":8080")
}
