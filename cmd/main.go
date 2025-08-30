package main

import (
	"go-project/controller"
	"go-project/db"
	"go-project/repository"
	"go-project/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
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
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)
	server.PUT("/product/:productId", productController.UpdateProductById)
	server.DELETE("/product/:productId", productController.DeleteProductById)

	server.Run(":81")
}
