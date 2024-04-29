package service

import (
	"product-app-go/internal/application/command"
)

type ProductService interface {
	Create(product command.CreateProductRequest)
	Update(product command.UpdateProductRequest)
	Delete(productId int)
	FindById(productId int) command.ProductResponse
	FindAll() []command.ProductResponse
}
