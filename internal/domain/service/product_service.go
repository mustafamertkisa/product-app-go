package service

import (
	"product-app-go/internal/application/request"
	"product-app-go/internal/application/response"
)

type ProductService interface {
	Create(product request.CreateProductRequest)
	Update(product request.UpdateProductRequest)
	Delete(productId int)
	FindById(productId int) response.ProductResponse
	FindAll() []response.ProductResponse
}
