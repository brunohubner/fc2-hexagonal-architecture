package dtos

import "github.com/brunohubner/fc2-hexagonal-architecture/src/application"

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *ProductDto {
	return &ProductDto{}
}

func (p *ProductDto) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	if _, err := product.IsValid(); err != nil {
		return &application.Product{}, err
	}

	return product, nil
}
