package repository

import "product-app-go/internal/domain/model"

type ProductRepository interface {
	Save(product model.Product) error
	Update(product model.Product) error
	Delete(productId int) error
	FindById(productId int) (model.Product, error)
	FindAll() ([]model.Product, error)
}
