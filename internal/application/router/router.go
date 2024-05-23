package router

import (
	"product-app-go/internal/api"
	"product-app-go/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(productController *api.ProductController, userController *api.UserController, orderController *api.OrderController, authController *api.AuthController, logController *api.LogController) *fiber.App {
	router := fiber.New()

	authRequired := middleware.NewAuthMiddleware()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to product app",
		})
	})

	// Product routes
	productRoutes := router.Group("/products")
	{
		productRoutes.Post("/", productController.Create)
		productRoutes.Get("/", productController.FindAll)

		productRoutes.Route("/:productId", func(productRouter fiber.Router) {
			productRouter.Delete("", productController.Delete)
			productRouter.Get("", productController.FindById)
			productRouter.Put("", productController.Update)
		})
	}

	// User routes
	userRoutes := router.Group("/users")
	{
		userRoutes.Post("/", userController.Create)
		userRoutes.Get("/", userController.FindAll)

		userRoutes.Route("/:userId", func(userRouter fiber.Router) {
			userRouter.Delete("", authRequired, userController.Delete)
			userRouter.Get("", userController.FindById)
			userRouter.Put("", authRequired, userController.Update)
		})
	}

	// Order routes
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.Post("/", orderController.Create)
		orderRoutes.Get("/", orderController.FindAll)

		orderRoutes.Route("/:orderId", func(orderRouter fiber.Router) {
			orderRouter.Delete("", orderController.Delete)
			orderRouter.Get("", orderController.FindById)
			orderRouter.Put("", orderController.Update)
		})
	}

	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.Post("/register", authController.Register)
		authRoutes.Post("/login", authController.Login)
		authRoutes.Post("/user", authController.User)
		authRoutes.Post("/logout", authController.Logout)
	}

	// Log routes
	logRoutes := router.Group("/logs")
	{
		logRoutes.Delete("/", logController.DeleteAllLogs)
		logRoutes.Get("/", logController.GetAllLogs)

		logRoutes.Route("/:id", func(logRouter fiber.Router) {
			logRouter.Get("", logController.GetLogById)
			logRouter.Delete("", logController.DeleteLogById)
		})

		logRoutes.Get("/user/:userId", logController.GetLogsByUserId)
	}

	return router
}
