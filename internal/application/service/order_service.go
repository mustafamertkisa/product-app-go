package service

import (
	"product-app-go/internal/application/command"
)

type OrderService interface {
	Create(order command.CreateOrderRequest) error
	Update(order command.UpdateOrderRequest) error
	Delete(orderId int) error
	FindById(orderId int) (command.OrderResponse, error)
	FindAll() ([]command.OrderResponse, error)
}
