package service

import (
	"product-app-go/data/request"
	"product-app-go/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}
