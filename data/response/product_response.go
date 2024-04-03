package response

type ProductResponse struct {
	Id    int
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
