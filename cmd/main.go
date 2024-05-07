package main

import (
	"fmt"
	"log"
	"os"
	"product-app-go/internal/api"
	"product-app-go/internal/application/router"
	"product-app-go/internal/application/service"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"
	"product-app-go/internal/infrastructure"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("Run service...")

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	dbUrl := os.Getenv("POSTGRES_URL")

	db := infrastructure.ConnectionDB(dbUrl)
	validate := validator.New()

	db.AutoMigrate(&model.Product{}, &model.User{}, &model.Order{})

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository, validate)
	productController := api.NewProductController(productService)

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, validate)
	userController := api.NewUserController(userService)

	orderRepository := repository.NewOrderRepositoryImpl(db)
	orderService := service.NewOrderServiceImpl(orderRepository, validate)
	orderController := api.NewOrderController(orderService)

	authService := service.NewAuthServiceImpl(userRepository, validate)
	authController := api.NewAuthController(authService)

	routes := router.NewRouter(productController, userController, orderController, authController)

	app := fiber.New()
	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
