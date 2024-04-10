package repository

import "product-app-go/internal/domain/model"

type ProductRepository interface {
	Save(product model.Product)
	Update(product model.Product)
	Delete(productId int)
	FindById(productId int) (model.Product, error)
	FindAll() []model.Product
}
