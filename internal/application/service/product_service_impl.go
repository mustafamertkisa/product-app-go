package service

import (
	"errors"
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

func (p *ProductServiceImpl) Create(product command.CreateProductRequest) error {
	err := p.validate.Struct(product)
	if err != nil {
		return err
	}

	productModel := model.Product{
		Name:  product.Name,
		Price: product.Price,
	}

	err = p.ProductRepository.Save(productModel)
	if err != nil {
		return errors.New("failed to save product: " + err.Error())
	}

	return nil
}

func (p *ProductServiceImpl) Update(product command.UpdateProductRequest) error {
	productData, err := p.ProductRepository.FindById(product.Id)
	if err != nil {
		return errors.New("failed to find product: " + err.Error())
	}

	productData.Name = product.Name
	productData.Price = product.Price
	err = p.ProductRepository.Update(productData)
	if err != nil {
		return errors.New("failed to update product: " + err.Error())
	}

	return nil
}

func (p *ProductServiceImpl) Delete(productId int) error {
	err := p.ProductRepository.Delete(productId)
	if err != nil {
		return errors.New("failed to delete product: " + err.Error())
	}

	return nil
}

func (p *ProductServiceImpl) FindById(productId int) (command.ProductResponse, error) {
	productData, err := p.ProductRepository.FindById(productId)
	if err != nil {
		return command.ProductResponse{}, errors.New("failed to find product: " + err.Error())
	}

	productResponse := command.ProductResponse{
		Id:    int(productData.Id),
		Name:  productData.Name,
		Price: productData.Price,
	}

	return productResponse, nil
}

func (p *ProductServiceImpl) FindAll() ([]command.ProductResponse, error) {
	result, err := p.ProductRepository.FindAll()
	if err != nil {
		return nil, errors.New("failed to find products: " + err.Error())
	}

	var products []command.ProductResponse

	for _, value := range result {
		product := command.ProductResponse{
			Id:    int(value.Id),
			Name:  value.Name,
			Price: value.Price,
		}
		products = append(products, product)
	}

	return products, nil
}
