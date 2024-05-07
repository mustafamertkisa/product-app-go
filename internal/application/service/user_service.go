package service

import (
	"product-app-go/internal/application/command"
)

type UserService interface {
	Create(user command.CreateUserRequest) error
	Update(user command.UpdateUserRequest) error
	Delete(userId int) error
	FindById(userId int) (command.UserResponse, error)
	FindAll() ([]command.UserResponse, error)
}
