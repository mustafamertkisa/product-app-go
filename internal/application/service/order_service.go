package service

import (
	"product-app-go/internal/application/command"
)

type OrderService interface {
	Create(order command.CreateOrderRequest)
	Update(order command.UpdateOrderRequest)
	Delete(orderId int)
	FindById(orderId int) command.OrderResponse
	FindAll() []command.OrderResponse
}
