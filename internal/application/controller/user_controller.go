package controller

import (
	"product-app-go/internal/application/request"
	"product-app-go/internal/application/response"
	"product-app-go/internal/domain/service"
	"product-app-go/internal/helper"
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
	createUserRequest := request.CreateUserRequest{}
	err := ctx.BodyParser(&createUserRequest)
	helper.ErrorPanic(err)

	controller.userService.Create(createUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.BodyParser(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.userService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated user data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.userService.Delete(id)

	webResponse := response.Response{
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
	helper.ErrorPanic(err)

	userResponse := controller.userService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get user data",
		Data:    userResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	userResponse := controller.userService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get users data",
		Data:    userResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
