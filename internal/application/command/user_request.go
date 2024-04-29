package command

type CreateUserRequest struct {
	Name  string `validate:"required,min=2,max=100" json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Id    int    `validate:"required"`
	Name  string `validate:"required,min=2,max=100" json:"name"`
	Email string `json:"email"`
}
