package service

import (
	"product-app-go/data/request"
	"product-app-go/data/response"
	"product-app-go/helper"
	"product-app-go/model"
	"product-app-go/repository"

	"github.com/go-playground/validator"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		validate:          validate,
	}
}

func (p *ProductServiceImpl) Create(product request.CreateProductRequest) {
	err := p.validate.Struct(product)
	helper.ErrorPanic(err)
	productModel := model.Product{
		Name:  product.Name,
		Price: product.Price,
	}
	p.ProductRepository.Save(productModel)
}

func (p *ProductServiceImpl) Update(product request.UpdateProductRequest) {
	productData, err := p.ProductRepository.FindById(product.Id)
	helper.ErrorPanic(err)
	productData.Name = product.Name
	productData.Price = product.Price
	p.ProductRepository.Update(productData)
}

func (p *ProductServiceImpl) Delete(productId int) {
	p.ProductRepository.Delete(productId)
}

func (p *ProductServiceImpl) FindById(productId int) response.ProductResponse {
	productData, err := p.ProductRepository.FindById(productId)
	helper.ErrorPanic(err)
	productResponse := response.ProductResponse{
		Id:    int(productData.Id),
		Name:  productData.Name,
		Price: productData.Price,
	}
	return productResponse
}

func (p *ProductServiceImpl) FindAll() []response.ProductResponse {
	result := p.ProductRepository.FindAll()
	var products []response.ProductResponse

	for _, value := range result {
		product := response.ProductResponse{
			Id:    int(value.Id),
			Name:  value.Name,
			Price: value.Price,
		}
		products = append(products, product)
	}

	return products
}
