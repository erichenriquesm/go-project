package repository

import (
	"database/sql"
	"fmt"
	"go-project/model"
)

type ProductRepostiroy struct {
	connection *sql.DB
}

func GetProductRepository(connection *sql.DB) ProductRepostiroy {
	return ProductRepostiroy{
		connection: connection,
	}
}

func (pr *ProductRepostiroy) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}
