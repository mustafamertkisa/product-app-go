package request

import "product-app-go/internal/domain/model"

type CreateOrderRequest struct {
	UserId    uint          `json:"userId"`
	ProductId uint          `json:"productId"`
	Quantity  int           `json:"quantity"`
	User      model.User    `json:"user"`
	Product   model.Product `json:"product"`
}

type UpdateOrderRequest struct {
	Id        int           `validate:"required"`
	UserId    uint          `json:"userId"`
	ProductId uint          `json:"productId"`
	Quantity  int           `json:"quantity"`
	User      model.User    `json:"user"`
	Product   model.Product `json:"product"`
}
