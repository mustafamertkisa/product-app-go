package controller

import (
	"product-app-go/data/request"
	"product-app-go/data/response"
	"product-app-go/helper"
	"product-app-go/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{productService: service}
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	createProductRequest := request.CreateProductRequest{}
	err := ctx.BodyParser(&createProductRequest)
	helper.ErrorPanic(err)

	controller.productService.Create(createProductRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created product data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) Update(ctx *fiber.Ctx) error {
	updateProductRequest := request.UpdateProductRequest{}
	err := ctx.BodyParser(&updateProductRequest)
	helper.ErrorPanic(err)

	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)

	updateProductRequest.Id = id

	controller.productService.Update(updateProductRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated product data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	controller.productService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted product data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindById(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)

	productResponse := controller.productService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get product data!",
		Data:    productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindAll(ctx *fiber.Ctx) error {
	productResponse := controller.productService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get products data!",
		Data:    productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
