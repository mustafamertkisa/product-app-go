package api

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	createUserRequest := command.CreateUserRequest{}
	err := ctx.BodyParser(&createUserRequest)
	if err != nil {
		return err
	}

	controller.authService.Register(createUserRequest)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully register user",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	userLoginRequest := command.UserLoginRequest{}
	err := ctx.BodyParser(&userLoginRequest)
	if err != nil {
		return err
	}

	userToken, err := controller.authService.Login(userLoginRequest, ctx)
	if err != nil {
		return err
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "User logged in successfully",
		Data:    userToken,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
