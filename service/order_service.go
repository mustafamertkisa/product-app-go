package service

import (
	"product-app-go/data/request"
	"product-app-go/data/response"
)

type OrderService interface {
	Create(order request.CreateOrderRequest)
	Update(order request.UpdateOrderRequest)
	Delete(orderId int)
	FindById(orderId int) response.OrderResponse
	FindAll() []response.OrderResponse
}
