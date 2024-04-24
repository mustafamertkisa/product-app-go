package command

import "product-app-go/internal/domain/model"

type CreateOrderRequest struct {
	UserId     int             `json:"userId"`
	ProductIds []int           `json:"productIds"`
	Quantity   int             `json:"quantity"`
	User       model.User      `json:"user"`
	Products   []model.Product `json:"products"`
}

type UpdateOrderRequest struct {
	Id         int             `validate:"required"`
	UserId     int             `json:"userId"`
	ProductIds []int           `json:"productIds"`
	Quantity   int             `json:"quantity"`
	User       model.User      `json:"user"`
	Products   []model.Product `json:"products"`
}
