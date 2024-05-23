package mocks

import (
	"errors"
	"product-app-go/internal/domain/model"
)

type MockOrderRepository struct {
	orders map[int]model.Order
	nextID int
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{
		orders: make(map[int]model.Order),
		nextID: 1,
	}
}

func (m *MockOrderRepository) Save(order model.Order) error {
	order.Id = m.nextID
	m.orders[m.nextID] = order
	m.nextID++
	return nil
}

func (m *MockOrderRepository) Update(order model.Order) error {
	if _, exists := m.orders[order.Id]; !exists {
		return errors.New("order not found")
	}
	m.orders[order.Id] = order
	return nil
}

func (m *MockOrderRepository) Delete(orderId int) error {
	if _, exists := m.orders[orderId]; !exists {
		return errors.New("order not found")
	}
	delete(m.orders, orderId)
	return nil
}

func (m *MockOrderRepository) FindById(orderId int) (model.Order, error) {
	order, exists := m.orders[orderId]
	if !exists {
		return model.Order{}, errors.New("order not found")
	}
	return order, nil
}

func (m *MockOrderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	for _, o := range m.orders {
		orders = append(orders, o)
	}
	return orders, nil
}

func (m *MockOrderRepository) FindUserById(userId int) (model.User, error) {
	// Dummy implementation, as it's not essential for testing OrderServiceImpl
	return model.User{}, nil
}

func (m *MockOrderRepository) FindProductById(productId int) (model.Product, error) {
	// Dummy implementation, as it's not essential for testing OrderServiceImpl
	return model.Product{}, nil
}
