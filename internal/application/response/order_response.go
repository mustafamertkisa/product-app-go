package response

import "product-app-go/internal/domain/model"

type OrderResponse struct {
	Id        int           `json:"id"`
	UserId    uint          `json:"userId"`
	ProductId uint          `json:"productId"`
	Quantity  int           `json:"quantity"`
	User      model.User    `json:"user"`
	Product   model.Product `json:"product"`
}
