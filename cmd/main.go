package main

import (
	"go-project/controller"
	"go-project/db"
	"go-project/repository"
	"go-project/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	productRepository := repository.GetProductRepository(dbConnection)

	productUsecase := usecase.GetProductUsecase(productRepository)

	productController := controller.GetProductController(productUsecase)

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":81")
}
