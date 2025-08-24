package controller

import (
	"go-project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func GetProductController(productUseCase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: productUseCase,
	}
}

func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
