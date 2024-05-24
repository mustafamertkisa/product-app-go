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

func (s *ProductServiceImpl) Create(product command.CreateProductRequest) error {
	err := s.validate.Struct(product)
	if err != nil {
		return err
	}

	productModel := model.Product{
		Name:  product.Name,
		Price: product.Price,
	}

	err = s.ProductRepository.Save(productModel)
	if err != nil {
		return errors.New("failed to save product: " + err.Error())
	}

	return nil
}

func (s *ProductServiceImpl) Update(product command.UpdateProductRequest) error {
	productData, err := s.ProductRepository.FindById(product.Id)
	if err != nil {
		return errors.New("failed to find product: " + err.Error())
	}

	productData.Name = product.Name
	productData.Price = product.Price
	err = s.ProductRepository.Update(productData)
	if err != nil {
		return errors.New("failed to update product: " + err.Error())
	}

	return nil
}

func (s *ProductServiceImpl) Delete(productId int) error {
	err := s.ProductRepository.Delete(productId)
	if err != nil {
		return errors.New("failed to delete product: " + err.Error())
	}

	return nil
}

func (s *ProductServiceImpl) FindById(productId int) (command.ProductResponse, error) {
	productData, err := s.ProductRepository.FindById(productId)
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

func (s *ProductServiceImpl) FindAll() ([]command.ProductResponse, error) {
	result, err := s.ProductRepository.FindAll()
	if err != nil {
		return nil, errors.New("failed to find products: " + err.Error())
	}

	if len(result) == 0 {
		return nil, errors.New("no data found")
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
