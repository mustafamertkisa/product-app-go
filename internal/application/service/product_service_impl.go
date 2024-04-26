package service

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"

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

func (p *ProductServiceImpl) Create(product command.CreateProductRequest) {
	err := p.validate.Struct(product)
	if err != nil {
		panic(err)
	}
	productModel := model.Product{
		Name:  product.Name,
		Price: product.Price,
	}
	p.ProductRepository.Save(productModel)
}

func (p *ProductServiceImpl) Update(product command.UpdateProductRequest) {
	productData, err := p.ProductRepository.FindById(product.Id)
	if err != nil {
		panic(err)
	}
	productData.Name = product.Name
	productData.Price = product.Price
	p.ProductRepository.Update(productData)
}

func (p *ProductServiceImpl) Delete(productId int) {
	p.ProductRepository.Delete(productId)
}

func (p *ProductServiceImpl) FindById(productId int) command.ProductResponse {
	productData, err := p.ProductRepository.FindById(productId)
	if err != nil {
		panic(err)
	}
	productResponse := command.ProductResponse{
		Id:    int(productData.Id),
		Name:  productData.Name,
		Price: productData.Price,
	}
	return productResponse
}

func (p *ProductServiceImpl) FindAll() []command.ProductResponse {
	result := p.ProductRepository.FindAll()
	var products []command.ProductResponse

	for _, value := range result {
		product := command.ProductResponse{
			Id:    int(value.Id),
			Name:  value.Name,
			Price: value.Price,
		}
		products = append(products, product)
	}

	return products
}
