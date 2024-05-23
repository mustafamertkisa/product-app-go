package test

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/test/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func initProductTestServices() service.ProductService {
	mockRepo := mocks.NewMockProductRepository()
	validate := validator.New()
	productService := service.NewProductServiceImpl(mockRepo, validate)

	return productService
}

func TestCreateProduct(t *testing.T) {
	productService := initProductTestServices()

	req := command.CreateProductRequest{
		Name:  "Test Product",
		Price: 10.50,
	}

	err := productService.Create(req)
	assert.Nil(t, err)

	createdProduct, err := productService.FindById(1)
	assert.Nil(t, err)

	expectedProduct := command.ProductResponse{
		Id:    1,
		Name:  "Test Product",
		Price: 10.50,
	}

	assert.Equal(t, expectedProduct, createdProduct)

}

func TestUpdateProduct(t *testing.T) {
	productService := initProductTestServices()

	createReq := command.CreateProductRequest{
		Name:  "Test Product",
		Price: 10.50,
	}
	_ = productService.Create(createReq)

	updateReq := command.UpdateProductRequest{
		Id:    1,
		Name:  "Updated Product",
		Price: 15.75,
	}
	err := productService.Update(updateReq)
	assert.Nil(t, err)

	updatedProduct, err := productService.FindById(updateReq.Id)
	assert.Nil(t, err)
	assert.Equal(t, "Updated Product", updatedProduct.Name)
	assert.Equal(t, 15.75, updatedProduct.Price)
}

func TestDeleteProduct(t *testing.T) {
	productService := initProductTestServices()

	createReq := command.CreateProductRequest{
		Name:  "Test Product",
		Price: 10.50,
	}
	_ = productService.Create(createReq)

	err := productService.Delete(1)
	assert.Nil(t, err)

	_, err = productService.FindById(1)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to find product: product not found", err.Error())
}

func TestFindProductById(t *testing.T) {
	productService := initProductTestServices()

	createReq := command.CreateProductRequest{
		Name:  "Test Product",
		Price: 10.50,
	}
	_ = productService.Create(createReq)

	product, err := productService.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, 10.50, product.Price)
}

func TestFindAllProducts(t *testing.T) {
	productService := initProductTestServices()

	createReq1 := command.CreateProductRequest{
		Name:  "Test Product 1",
		Price: 10.50,
	}
	createReq2 := command.CreateProductRequest{
		Name:  "Test Product 2",
		Price: 20.75,
	}

	_ = productService.Create(createReq1)
	_ = productService.Create(createReq2)

	products, err := productService.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
}
