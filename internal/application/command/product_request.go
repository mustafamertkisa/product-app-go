package command

type CreateProductRequest struct {
	Name  string  `validate:"required,min=2,max=100" json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductRequest struct {
	Id    int     `validate:"required"`
	Name  string  `validate:"required,min=2,max=100" json:"name"`
	Price float64 `json:"price"`
}
