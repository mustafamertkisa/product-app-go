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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = controller.authService.Register(createUserRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userToken, err := controller.authService.Login(userLoginRequest, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	loginResponse := command.LoginResponse{
		Code:    200,
		Status:  "Ok",
		Message: "User logged in successfully",
		Token:   userToken,
	}

	return ctx.Status(fiber.StatusCreated).JSON(loginResponse)
}

func (c *AuthController) User(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	user, err := c.authService.GetUserFromToken(cookie)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(user)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	err := c.authService.Logout(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Logout success",
	})
}
