package api

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(service service.OrderService) *OrderController {
	return &OrderController{orderService: service}
}

func (controller *OrderController) Create(ctx *fiber.Ctx) error {
	createOrderRequest := command.CreateOrderRequest{}
	err := ctx.BodyParser(&createOrderRequest)
	if err != nil {
		return err
	}

	controller.orderService.Create(createOrderRequest)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created order data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *OrderController) Update(ctx *fiber.Ctx) error {
	updateOrderRequest := command.UpdateOrderRequest{}
	err := ctx.BodyParser(&updateOrderRequest)
	if err != nil {
		return err
	}

	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		return err
	}

	updateOrderRequest.Id = id

	controller.orderService.Update(updateOrderRequest)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated order data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *OrderController) Delete(ctx *fiber.Ctx) error {
	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		return err
	}

	controller.orderService.Delete(id)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted order data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *OrderController) FindById(ctx *fiber.Ctx) error {
	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		return err
	}

	orderResponse, err := controller.orderService.FindById(id)
	if err != nil {
		return err
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get order data",
		Data:    orderResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *OrderController) FindAll(ctx *fiber.Ctx) error {
	orderResponse, err := controller.orderService.FindAll()
	if err != nil {
		return err
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get orders data",
		Data:    orderResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
