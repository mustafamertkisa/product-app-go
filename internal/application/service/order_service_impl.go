package service

import (
	"errors"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"

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

func (o *OrderServiceImpl) Create(order command.CreateOrderRequest) error {
	err := o.validate.Struct(order)
	if err != nil {
		return err
	}

	user, err := o.OrderRepository.FindUserById(order.UserId)
	if err != nil {
		return errors.New("failed to find user: " + err.Error())
	}

	var products []model.Product
	for _, productId := range order.ProductIds {
		product, err := o.OrderRepository.FindProductById(productId)
		if err != nil {
			return errors.New("failed to find product: " + err.Error())
		}

		products = append(products, product)
	}

	orderModel := model.Order{
		UserId:   order.UserId,
		User:     user,
		Products: products,
		Quantity: order.Quantity,
	}

	err = o.OrderRepository.Save(orderModel)
	if err != nil {
		return errors.New("failed to save order: " + err.Error())
	}

	return nil
}

func (o *OrderServiceImpl) Update(order command.UpdateOrderRequest) error {
	orderData, err := o.OrderRepository.FindById(order.Id)
	if err != nil {
		return errors.New("failed to find order: " + err.Error())
	}

	user, err := o.OrderRepository.FindUserById(order.UserId)
	if err != nil {
		return errors.New("failed to find user: " + err.Error())
	}

	var products []model.Product
	for _, productId := range order.ProductIds {
		product, err := o.OrderRepository.FindProductById(productId)
		if err != nil {
			return errors.New("failed to find product: " + err.Error())
		}

		products = append(products, product)
	}

	orderData.UserId = order.UserId
	orderData.Quantity = order.Quantity
	orderData.User = user
	orderData.Products = products

	err = o.OrderRepository.Update(orderData)
	if err != nil {
		return errors.New("failed to update order: " + err.Error())
	}

	return nil
}

func (o *OrderServiceImpl) Delete(orderId int) error {
	err := o.OrderRepository.Delete(orderId)
	if err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}

	return nil
}

func (o *OrderServiceImpl) FindById(orderId int) (command.OrderResponse, error) {
	orderData, err := o.OrderRepository.FindById(orderId)
	if err != nil {
		return command.OrderResponse{}, errors.New("failed to find order: " + err.Error())
	}

	orderResponse := command.OrderResponse{
		Id:       int(orderData.Id),
		UserId:   orderData.UserId,
		Quantity: orderData.Quantity,
		User:     orderData.User,
		Products: orderData.Products,
	}

	return orderResponse, nil
}

func (o *OrderServiceImpl) FindAll() ([]command.OrderResponse, error) {
	result, err := o.OrderRepository.FindAll()
	if err != nil {
		return nil, errors.New("failed to find orders: " + err.Error())
	}

	var orders []command.OrderResponse

	for _, value := range result {
		order := command.OrderResponse{
			Id:       int(value.Id),
			UserId:   value.UserId,
			Quantity: value.Quantity,
			User:     value.User,
			Products: value.Products,
		}
		orders = append(orders, order)
	}

	return orders, nil
}
