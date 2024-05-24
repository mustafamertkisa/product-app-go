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

func (s *OrderServiceImpl) Create(order command.CreateOrderRequest) error {
	err := s.validate.Struct(order)
	if err != nil {
		return err
	}

	user, err := s.OrderRepository.FindUserById(order.UserId)
	if err != nil {
		return errors.New("failed to find user: " + err.Error())
	}

	var products []model.Product
	for _, productId := range order.ProductIds {
		product, err := s.OrderRepository.FindProductById(productId)
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

	err = s.OrderRepository.Save(orderModel)
	if err != nil {
		return errors.New("failed to save order: " + err.Error())
	}

	return nil
}

func (s *OrderServiceImpl) Update(order command.UpdateOrderRequest) error {
	orderData, err := s.OrderRepository.FindById(order.Id)
	if err != nil {
		return errors.New("failed to find order: " + err.Error())
	}

	user, err := s.OrderRepository.FindUserById(order.UserId)
	if err != nil {
		return errors.New("failed to find user: " + err.Error())
	}

	var products []model.Product
	for _, productId := range order.ProductIds {
		product, err := s.OrderRepository.FindProductById(productId)
		if err != nil {
			return errors.New("failed to find product: " + err.Error())
		}

		products = append(products, product)
	}

	orderData.UserId = order.UserId
	orderData.Quantity = order.Quantity
	orderData.User = user
	orderData.Products = products

	err = s.OrderRepository.Update(orderData)
	if err != nil {
		return errors.New("failed to update order: " + err.Error())
	}

	return nil
}

func (s *OrderServiceImpl) Delete(orderId int) error {
	err := s.OrderRepository.Delete(orderId)
	if err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}

	return nil
}

func (s *OrderServiceImpl) FindById(orderId int) (command.OrderResponse, error) {
	orderData, err := s.OrderRepository.FindById(orderId)
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

func (s *OrderServiceImpl) FindAll() ([]command.OrderResponse, error) {
	result, err := s.OrderRepository.FindAll()
	if err != nil {
		return nil, errors.New("failed to find orders: " + err.Error())
	}

	if len(result) == 0 {
		return nil, errors.New("no data found")
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
