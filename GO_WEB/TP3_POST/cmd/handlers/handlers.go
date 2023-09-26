package handlers

import (
	"eje1tp3/cmd/productos"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(len(productos.Productos))
		var data productos.RequestProduct
		var dtpure productos.Producto
		var dtout productos.ResponseBody
		var dto productos.Data
		if err := ctx.ShouldBindJSON(&data); err != nil {
			dto.ID = data.ID
			dto.CodeValue = data.CodeValue
			dto.Expiration = data.Expiration
			dto.IsPublished = data.IsPublished
			dto.Name = data.Name
			dto.Price = data.Price
			dto.Quantity = data.Quantity
			dtout.Data = &dto

			dtout.Mensjage = "i cant create this"

			ctx.JSON(http.StatusBadRequest, dtout)

			return
		}
		//vamos a trabajar con el tipo producto
		dtpure.ID = data.ID
		dtpure.Name = data.Name
		dtpure.Quantity = data.Quantity
		dtpure.CodeValue = data.CodeValue
		dtpure.IsPublished = data.IsPublished
		dtpure.Expiration = data.Expiration
		dtpure.Price = data.Price

		fmt.Println(dtpure)
		if err := productos.ValidateProduct(&dtpure); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := productos.ValidateCodeValue(&dtpure); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		if err := productos.ValidateDate(&dtpure); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		//esto hago para agregar a mi db simulada

		productos.Productos = append(productos.Productos, dtpure)
		//ahora armo el response lindo
		dto.CodeValue = dtpure.CodeValue
		dto.Expiration = dtpure.Expiration
		dto.IsPublished = dtpure.IsPublished
		dto.Name = dtpure.Name
		dto.Price = dtpure.Price
		dto.Quantity = dtpure.Quantity
		dto.ID = dtpure.ID
		dtout.Data = &dto
		dtout.Mensjage = "Successful..."

		ctx.JSON(201, dtout)
		return
	}
}

func GetByIdPatchUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dtout productos.ResponseBody
		var dto productos.Data
		idstr := ctx.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request url not valid",
			})
			return
		}
		producto, err := productos.SearchById(id)
		if err != nil {
			ctx.String(404, fmt.Sprintf("Not fount with %d", id))
			return
		}
		dto.ID = producto.ID
		dto.Name = producto.Name
		dto.Quantity = producto.Quantity
		dto.CodeValue = producto.CodeValue
		dto.IsPublished = producto.IsPublished
		dto.Expiration = producto.Expiration
		dto.Price = producto.Price
		dtout.Data = &dto
		dtout.Mensjage = "Successful.."
		ctx.JSON(200, dtout)
	}
}

func GetByPriceMayor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idq := ctx.Query("priceGt")
		idf, err := strconv.ParseFloat(idq, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})

		}

		datas, errf := productos.SearchMayPrice(idf)
		if errf != nil {
			ctx.String(http.StatusNotFound, errf.Error())
		}

		ctx.JSON(200, datas)

	}
}
