package controller

import (
	"go-project/model"
	"go-project/usecase"
	"go-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func GetProductController(productUsecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: productUsecase,
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

func (pc *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	if errs := utils.ValidateJSON(ctx, &product); errs != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	insertedProduct, err := pc.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}
