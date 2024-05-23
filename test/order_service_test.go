package test

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"

	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/test/mocks"
)

func initOrderTestServices() service.OrderService {
	mockRepo := mocks.NewMockOrderRepository()
	validate := validator.New()
	orderService := service.NewOrderServiceImpl(mockRepo, validate)

	return orderService
}

func TestCreateOrder(t *testing.T) {
	orderService := initOrderTestServices()

	req := command.CreateOrderRequest{
		UserId:     1,
		ProductIds: []int{1, 2, 3},
		Quantity:   2,
	}

	err := orderService.Create(req)
	assert.Nil(t, err)

	createdOrder, err := orderService.FindById(1)
	assert.Nil(t, err)

	assert.Equal(t, createdOrder.Id, 1)
	assert.Equal(t, createdOrder.UserId, req.UserId)
	assert.Equal(t, createdOrder.Quantity, req.Quantity)
	assert.Equal(t, len(createdOrder.Products), len(req.ProductIds))
}

func TestUpdateOrder(t *testing.T) {
	orderService := initOrderTestServices()

	req := command.CreateOrderRequest{
		UserId:     1,
		ProductIds: []int{1, 2, 3},
		Quantity:   2,
	}

	err := orderService.Create(req)
	assert.Nil(t, err)

	updateReq := command.UpdateOrderRequest{
		Id:         1,
		UserId:     2,
		ProductIds: []int{4, 5},
		Quantity:   3,
	}
	err = orderService.Update(updateReq)
	assert.Nil(t, err)

	updatedOrder, err := orderService.FindById(1)
	assert.Nil(t, err)

	assert.Equal(t, updateReq.UserId, updatedOrder.UserId)
	assert.Equal(t, updateReq.Quantity, updatedOrder.Quantity)
	assert.Len(t, updatedOrder.Products, 2)

}

func TestFindOrderById(t *testing.T) {
	orderService := initOrderTestServices()

	req := command.CreateOrderRequest{
		UserId:     1,
		ProductIds: []int{1, 2, 3},
		Quantity:   2,
	}

	err := orderService.Create(req)
	assert.Nil(t, err)

	foundOrder, err := orderService.FindById(1)
	assert.Nil(t, err)

	assert.Equal(t, req.UserId, foundOrder.UserId)
	assert.Equal(t, req.Quantity, foundOrder.Quantity)
	assert.Len(t, foundOrder.Products, 3)
}
