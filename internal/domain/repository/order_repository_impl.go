package repository

import (
	"errors"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/helper"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrderRepositoryImpl(Db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{Db: Db}
}

func (o *OrderRepositoryImpl) Save(order model.Order) {
	result := o.Db.Create(&order)
	helper.ErrorPanic(result.Error)
}

func (o *OrderRepositoryImpl) Update(order model.Order) {
	updateData := map[string]interface{}{
		"user_id":  order.UserId,
		"quantity": order.Quantity,
	}

	o.Db.Model(&order).Association("Products").Replace(order.Products)
	o.Db.Model(&order).Updates(updateData)
}

func (o *OrderRepositoryImpl) Delete(orderId int) {
	var order model.Order
	result := o.Db.Where("id = ?", orderId).Delete(&order)
	helper.ErrorPanic(result.Error)
}

func (o *OrderRepositoryImpl) FindAll() []model.Order {
	var orders []model.Order
	if err := o.Db.Preload("User").Preload("Products").Find(&orders).Error; err != nil {
		return nil
	}

	return orders
}

func (o *OrderRepositoryImpl) FindById(orderId int) (model.Order, error) {
	var order model.Order
	result := o.Db.Preload("User").Preload("Products").First(&order, orderId)
	if result == nil {
		return order, errors.New("order is not found")
	}

	return order, nil
}

func (o *OrderRepositoryImpl) FindUserById(userId int) (model.User, error) {
	var user model.User
	result := o.Db.First(&user, userId)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (o *OrderRepositoryImpl) FindProductById(productId int) (model.Product, error) {
	var product model.Product
	result := o.Db.First(&product, productId)
	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
