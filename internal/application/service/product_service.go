package service

import (
	"product-app-go/internal/application/command"
)

type ProductService interface {
	Create(product command.CreateProductRequest) error
	Update(product command.UpdateProductRequest) error
	Delete(productId int) error
	FindById(productId int) (command.ProductResponse, error)
	FindAll() ([]command.ProductResponse, error)
}
