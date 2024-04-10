package service

import (
	"product-app-go/internal/application/request"
	"product-app-go/internal/application/response"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"
	"product-app-go/internal/helper"

	"github.com/go-playground/validator"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	validate        *validator.Validate
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		validate:        validate,
	}
}

func (o *OrderServiceImpl) Create(order request.CreateOrderRequest) {
	err := o.validate.Struct(order)
	helper.ErrorPanic(err)
	orderModel := model.Order{
		UserId:    order.UserId,
		ProductId: order.ProductId,
		Quantity:  order.Quantity,
		User:      order.User,
		Product:   order.Product,
	}
	o.OrderRepository.Save(orderModel)
}

func (o *OrderServiceImpl) Update(order request.UpdateOrderRequest) {
	orderData, err := o.OrderRepository.FindById(order.Id)
	helper.ErrorPanic(err)
	orderData.UserId = order.UserId
	orderData.ProductId = order.ProductId
	orderData.Quantity = order.Quantity
	orderData.User = order.User
	orderData.Product = order.Product
	o.OrderRepository.Update(orderData)
}

func (o *OrderServiceImpl) Delete(orderId int) {
	o.OrderRepository.Delete(orderId)
}

func (o *OrderServiceImpl) FindById(orderId int) response.OrderResponse {
	orderData, err := o.OrderRepository.FindById(orderId)
	helper.ErrorPanic(err)
	orderResponse := response.OrderResponse{
		Id:        int(orderData.Id),
		UserId:    orderData.UserId,
		ProductId: orderData.ProductId,
		Quantity:  orderData.Quantity,
		User:      orderData.User,
		Product:   orderData.Product,
	}
	return orderResponse
}

func (o *OrderServiceImpl) FindAll() []response.OrderResponse {
	result := o.OrderRepository.FindAll()
	var orders []response.OrderResponse

	for _, value := range result {
		order := response.OrderResponse{
			Id:        int(value.Id),
			UserId:    value.UserId,
			ProductId: value.ProductId,
			Quantity:  value.Quantity,
			User:      value.User,
			Product:   value.Product,
		}
		orders = append(orders, order)
	}

	return orders
}
