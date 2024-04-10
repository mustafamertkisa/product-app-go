package service

import (
	"product-app-go/internal/application/request"
	"product-app-go/internal/application/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}
