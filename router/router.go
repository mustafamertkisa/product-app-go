package router

import (
	"product-app-go/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(productController *controller.ProductController, userController *controller.UserController, orderController *controller.OrderController) *fiber.App {
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

	router.Route("/users", func(router fiber.Router) {
		router.Post("/", userController.Create)
		router.Get("", userController.FindAll)
	})

	router.Route("/users/:userId", func(router fiber.Router) {
		router.Delete("", userController.Delete)
		router.Get("", userController.FindById)
		router.Patch("", userController.Update)
	})

	router.Route("/orders", func(router fiber.Router) {
		router.Post("/", orderController.Create)
		router.Get("", orderController.FindAll)
	})

	router.Route("/orders/:orderId", func(router fiber.Router) {
		router.Delete("", orderController.Delete)
		router.Get("", orderController.FindById)
		router.Patch("", orderController.Update)
	})

	return router
}
