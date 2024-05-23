package api

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
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
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = controller.productService.Create(createProductRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

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
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productId := ctx.Params("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updateProductRequest.Id = id

	err = controller.productService.Update(updateProductRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

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
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = controller.productService.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

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
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productResponse, err := controller.productService.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get product data",
		Data:    productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindAll(ctx *fiber.Ctx) error {
	productResponse, err := controller.productService.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	webResponse := command.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get products data",
		Data:    productResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
