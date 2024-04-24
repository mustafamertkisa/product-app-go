package command

import "product-app-go/internal/domain/model"

type OrderResponse struct {
	Id       int             `json:"id"`
	UserId   int             `json:"userId"`
	Quantity int             `json:"quantity"`
	User     model.User      `json:"user"`
	Products []model.Product `json:"products"`
}
