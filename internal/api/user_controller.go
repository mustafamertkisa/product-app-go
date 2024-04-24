package api

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/utils"
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
	utils.ErrorPanic(err)

	controller.userService.Create(createUserRequest)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	updateUserRequest := command.UpdateUserRequest{}
	err := ctx.BodyParser(&updateUserRequest)
	utils.ErrorPanic(err)

	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.userService.Update(updateUserRequest)

	webResponse := command.Response{
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
	utils.ErrorPanic(err)
	controller.userService.Delete(id)

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
	utils.ErrorPanic(err)

	userResponse := controller.userService.FindById(id)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get user data",
		Data:    userResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	userResponse := controller.userService.FindAll()

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get users data",
		Data:    userResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
