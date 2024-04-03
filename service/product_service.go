package service

import (
	"product-app-go/data/request"
	"product-app-go/data/response"
)

type ProductService interface {
	Create(product request.CreateProductRequest)
	Update(product request.UpdateProductRequest)
	Delete(productId int)
	FindById(productId int) response.ProductResponse
	FindAll() []response.ProductResponse
}
