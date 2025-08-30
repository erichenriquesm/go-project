package model

type Product struct {
	Id    int     `json:"product_id"`
	Name  string  `json:"name" binding:"omitempty,min=3,max=100"`
	Price float64 `json:"price" binding:"omitempty,gt=0"`
}
