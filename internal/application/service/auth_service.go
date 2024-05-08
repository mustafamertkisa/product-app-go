package service

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(user command.CreateUserRequest) error
	Login(user command.UserLoginRequest, ctx *fiber.Ctx) (string, error)
	GetUserFromToken(cookie string) (model.User, error)
	Logout(ctx *fiber.Ctx) error
}
