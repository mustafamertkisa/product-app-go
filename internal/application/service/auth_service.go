package service

import (
	"product-app-go/internal/application/command"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(user command.CreateUserRequest) error
	Login(user command.UserLoginRequest, ctx *fiber.Ctx) (string, error)
}
