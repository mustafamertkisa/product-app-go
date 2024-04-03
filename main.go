package main

import (
	"fmt"
	"log"
	"product-app-go/config"
	"product-app-go/controller"
	"product-app-go/model"
	"product-app-go/repository"
	"product-app-go/router"
	"product-app-go/service"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Printf("Run service...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load env", err)
	}

	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("products").AutoMigrate(&model.Product{})

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository, validate)
	productController := controller.NewProductController(productService)

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, validate)
	userController := controller.NewUserController(userService)

	orderRepository := repository.NewOrderRepositoryImpl(db)
	orderService := service.NewOrderServiceImpl(orderRepository, validate)
	orderController := controller.NewOrderController(orderService)

	routes := router.NewRouter(productController, userController, orderController)

	app := fiber.New()
	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
