package repository

import (
	"database/sql"
	"errors"
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
	query := "SELECT * FROM products"
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

func (pr *ProductRepostiroy) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO products " +
		"(name, price)" +
		"VALUES ($1, $2) RETURNING id",
	)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepostiroy) FindProductById(productId int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1 LIMIT 1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(productId).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepostiroy) UpdateProductById(productId int, productData model.Product) (*model.Product, error) {

	if !pr.exist(productId) {
		return nil, errors.New("product doesn't exists")
	}

	query, err := pr.connection.Prepare("UPDATE products SET name=$1, price=$2 WHERE id=$3")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(productData.Name, productData.Price, productId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	productData.Id = productId

	return &productData, nil
}

func (pr *ProductRepostiroy) DeleteProductById(productId int) (bool, error) {

	if !pr.exist(productId) {
		return false, errors.New("product doesn't exists")
	}

	query, err := pr.connection.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	_, err = query.Exec(productId)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}

func (pr *ProductRepostiroy) exist(productId int) (exist bool) {
	exists := false
	query, err := pr.connection.Prepare("SELECT EXISTS(SELECT id FROM products WHERE id=$1)")

	if err != nil {
		fmt.Println(err)
		return false
	}

	err = query.QueryRow(productId).Scan(&exists)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return exists
}
