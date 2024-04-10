package controller

import (
	"product-app-go/internal/application/request"
	"product-app-go/internal/application/response"
	"product-app-go/internal/domain/service"
	"product-app-go/internal/helper"
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
	createOrderRequest := request.CreateOrderRequest{}
	err := ctx.BodyParser(&createOrderRequest)
	helper.ErrorPanic(err)

	controller.orderService.Create(createOrderRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created order data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *OrderController) Update(ctx *fiber.Ctx) error {
	updateOrderRequest := request.UpdateOrderRequest{}
	err := ctx.BodyParser(&updateOrderRequest)
	helper.ErrorPanic(err)

	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	helper.ErrorPanic(err)

	updateOrderRequest.Id = id

	controller.orderService.Update(updateOrderRequest)

	webResponse := response.Response{
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
	helper.ErrorPanic(err)
	controller.orderService.Delete(id)

	webResponse := response.Response{
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
	helper.ErrorPanic(err)

	orderResponse := controller.orderService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get order data",
		Data:    orderResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *OrderController) FindAll(ctx *fiber.Ctx) error {
	orderResponse := controller.orderService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get orders data",
		Data:    orderResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
