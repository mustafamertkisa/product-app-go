package service

import (
	"product-app-go/internal/application/request"
	"product-app-go/internal/application/response"
)

type OrderService interface {
	Create(order request.CreateOrderRequest)
	Update(order request.UpdateOrderRequest)
	Delete(orderId int)
	FindById(orderId int) response.OrderResponse
	FindAll() []response.OrderResponse
}
