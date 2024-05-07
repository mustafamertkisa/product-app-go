package repository

import "product-app-go/internal/domain/model"

type OrderRepository interface {
	Save(order model.Order) error
	Update(order model.Order) error
	Delete(orderId int) error
	FindById(orderId int) (model.Order, error)
	FindAll() ([]model.Order, error)
	FindUserById(userId int) (model.User, error)
	FindProductById(productId int) (model.Product, error)
}
