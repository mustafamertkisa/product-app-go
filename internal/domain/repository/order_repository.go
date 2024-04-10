package repository

import "product-app-go/internal/domain/model"

type OrderRepository interface {
	Save(order model.Order)
	Update(order model.Order)
	Delete(orderId int)
	FindById(orderId int) (model.Order, error)
	FindAll() []model.Order
}
