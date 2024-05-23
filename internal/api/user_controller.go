package api

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	createUserRequest := command.CreateUserRequest{}
	err := ctx.BodyParser(&createUserRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = controller.userService.Create(createUserRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	checkJWT := middleware.ValidateJWTClaims(ctx)
	if !checkJWT {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "unauthorized"})
	}

	updateUserRequest := command.UpdateUserRequest{}
	err := ctx.BodyParser(&updateUserRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := strconv.Atoi(ctx.Params("userId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updateUserRequest.Id = id

	err = controller.userService.Update(updateUserRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated user data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	checkJWT := middleware.ValidateJWTClaims(ctx)
	if !checkJWT {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "unauthorized"})
	}

	id, err := strconv.Atoi(ctx.Params("userId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = controller.userService.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted user data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) FindById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userResponse, err := controller.userService.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get user data",
		Data:    userResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	userResponse, err := controller.userService.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get users data",
		Data:    userResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
