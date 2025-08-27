package controller

import (
	"go-project/model"
	"go-project/usecase"
	"go-project/utils"
	"net/http"
	"strconv"

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
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (pc *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Product id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Product id must be a int",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.productUsecase.FindProductById(productId)

	if product == nil && err == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
