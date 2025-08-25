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

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.Id = id
	return product, nil
}
