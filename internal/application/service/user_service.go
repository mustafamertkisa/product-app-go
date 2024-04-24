package service

import (
	"product-app-go/internal/application/command"
)

type UserService interface {
	Create(user command.CreateUserRequest)
	Update(user command.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) command.UserResponse
	FindAll() []command.UserResponse
}
