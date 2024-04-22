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

	user, err := o.OrderRepository.FindUserById(order.UserId)
	helper.ErrorPanic(err)

	var products []model.Product
	for _, productId := range order.ProductIds {
		product, err := o.OrderRepository.FindProductById(productId)
		helper.ErrorPanic(err)

		products = append(products, product)
	}

	orderModel := model.Order{
		UserId:   order.UserId,
		User:     user,
		Products: products,
		Quantity: order.Quantity,
	}

	o.OrderRepository.Save(orderModel)
}

func (o *OrderServiceImpl) Update(order request.UpdateOrderRequest) {
	orderData, err := o.OrderRepository.FindById(order.Id)
	helper.ErrorPanic(err)

	user, err := o.OrderRepository.FindUserById(order.UserId)
	helper.ErrorPanic(err)

	var products []model.Product
	for _, productId := range order.ProductIds {
		product, err := o.OrderRepository.FindProductById(productId)
		helper.ErrorPanic(err)

		products = append(products, product)
	}

	orderData.UserId = order.UserId
	orderData.Quantity = order.Quantity
	orderData.User = user
	orderData.Products = products

	o.OrderRepository.Update(orderData)
}

func (o *OrderServiceImpl) Delete(orderId int) {
	o.OrderRepository.Delete(orderId)
}

func (o *OrderServiceImpl) FindById(orderId int) response.OrderResponse {
	orderData, err := o.OrderRepository.FindById(orderId)
	helper.ErrorPanic(err)

	orderResponse := response.OrderResponse{
		Id:       int(orderData.Id),
		UserId:   orderData.UserId,
		Quantity: orderData.Quantity,
		User:     orderData.User,
		Products: orderData.Products,
	}

	return orderResponse
}

func (o *OrderServiceImpl) FindAll() []response.OrderResponse {
	result := o.OrderRepository.FindAll()
	var orders []response.OrderResponse

	for _, value := range result {
		order := response.OrderResponse{
			Id:       int(value.Id),
			UserId:   value.UserId,
			Quantity: value.Quantity,
			User:     value.User,
			Products: value.Products,
		}
		orders = append(orders, order)
	}

	return orders
}
