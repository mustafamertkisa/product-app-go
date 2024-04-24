package api

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/utils"
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
	createProductRequest := command.CreateProductRequest{}
	err := ctx.BodyParser(&createProductRequest)
	utils.ErrorPanic(err)

	controller.productService.Create(createProductRequest)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created product data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) Update(ctx *fiber.Ctx) error {
	updateProductRequest := command.UpdateProductRequest{}
	err := ctx.BodyParser(&updateProductRequest)
	utils.ErrorPanic(err)

	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	utils.ErrorPanic(err)

	updateProductRequest.Id = id

	controller.productService.Update(updateProductRequest)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated product data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	utils.ErrorPanic(err)
	controller.productService.Delete(id)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted product data",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindById(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	utils.ErrorPanic(err)

	productResponse := controller.productService.FindById(id)

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get product data",
		Data:    productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindAll(ctx *fiber.Ctx) error {
	productResponse := controller.productService.FindAll()

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get products data",
		Data:    productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
