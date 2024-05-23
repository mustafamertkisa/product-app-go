package mocks

import (
	"errors"
	"product-app-go/internal/domain/model"
)

type MockProductRepository struct {
	products map[int]model.Product
	nextID   int
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: make(map[int]model.Product),
		nextID:   1,
	}
}

func (m *MockProductRepository) Save(product model.Product) error {
	if _, exists := m.products[product.Id]; exists {
		return errors.New("product already exists")
	}
	product.Id = m.nextID // Assign the next available ID
	m.products[product.Id] = product
	m.nextID++ // Increment nextID for the next product
	return nil
}

func (m *MockProductRepository) Update(product model.Product) error {
	if _, exists := m.products[product.Id]; !exists {
		return errors.New("product not found")
	}
	m.products[product.Id] = product
	return nil
}

func (m *MockProductRepository) Delete(productId int) error {
	if _, exists := m.products[productId]; !exists {
		return errors.New("product not found")
	}
	delete(m.products, productId)
	return nil
}

func (m *MockProductRepository) FindById(productId int) (model.Product, error) {
	product, exists := m.products[productId]
	if !exists {
		return model.Product{}, errors.New("product not found")
	}
	return product, nil
}

func (m *MockProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil
}
