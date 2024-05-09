package repository

import (
	"errors"
	"product-app-go/internal/domain/model"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrderRepositoryImpl(Db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{Db: Db}
}

func (r *OrderRepositoryImpl) Save(order model.Order) error {
	result := r.Db.Create(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryImpl) Update(order model.Order) error {
	updateData := map[string]interface{}{
		"user_id":  order.UserId,
		"quantity": order.Quantity,
	}

	r.Db.Model(&order).Association("Products").Replace(order.Products)
	result := r.Db.Model(&order).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryImpl) Delete(orderId int) error {
	var order model.Order
	result := r.Db.Where("id = ?", orderId).Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryImpl) FindAll() ([]model.Order, error) {
	var orders []model.Order
	if err := r.Db.Preload("User").Preload("Products").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepositoryImpl) FindById(orderId int) (model.Order, error) {
	var order model.Order
	result := r.Db.Preload("User").Preload("Products").First(&order, orderId)
	if result == nil {
		return order, errors.New("order is not found")
	}

	return order, nil
}

func (r *OrderRepositoryImpl) FindUserById(userId int) (model.User, error) {
	var user model.User
	result := r.Db.First(&user, userId)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *OrderRepositoryImpl) FindProductById(productId int) (model.Product, error) {
	var product model.Product
	result := r.Db.First(&product, productId)
	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
