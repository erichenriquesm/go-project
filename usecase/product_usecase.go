package usecase

import (
	"go-project/model"
	"go-project/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepostiroy
}

func GetProductUsecase(repository repository.ProductRepostiroy) ProductUsecase {
	return ProductUsecase{
		repository: repository,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
