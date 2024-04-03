package response

type UserResponse struct {
	Id    int
	Name  string `json:"name"`
	Email string `json:"email"`
}
