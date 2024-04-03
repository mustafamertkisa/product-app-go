package repository

import (
	"errors"
	"product-app-go/data/request"
	"product-app-go/helper"
	"product-app-go/model"

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
	var updateOrder = request.UpdateOrderRequest{Id: int(order.Id), UserId: order.UserId, ProductId: order.ProductId, Quantity: order.Quantity, User: order.User, Product: order.Product}
	result := o.Db.Model(&order).Updates(updateOrder)
	helper.ErrorPanic(result.Error)
}

func (o *OrderRepositoryImpl) Delete(orderId int) {
	var order model.Order
	result := o.Db.Where("id = ?", orderId).Delete(&order)
	helper.ErrorPanic(result.Error)
}

func (o *OrderRepositoryImpl) FindAll() []model.Order {
	var order []model.Order
	result := o.Db.Find(&order)
	helper.ErrorPanic(result.Error)
	return order
}

func (o *OrderRepositoryImpl) FindById(orderId int) (model.Order, error) {
	var order model.Order
	result := o.Db.Find(&order, orderId)
	if result == nil {
		return order, errors.New("order is not found")
	}

	return order, nil
}
