package router

import (
	"product-app-go/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(productController *controller.ProductController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to product app",
		})
	})

	router.Route("/products", func(router fiber.Router) {
		router.Post("/", productController.Create)
		router.Get("", productController.FindAll)
	})

	router.Route("/products/:productId", func(router fiber.Router) {
		router.Delete("", productController.Delete)
		router.Get("", productController.FindById)
		router.Patch("", productController.Update)
	})

	return router
}
