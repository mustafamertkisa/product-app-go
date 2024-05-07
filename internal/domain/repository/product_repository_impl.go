package repository

import (
	"errors"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p *ProductRepositoryImpl) Save(product model.Product) error {
	result := p.Db.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepositoryImpl) Update(product model.Product) error {
	var updateProduct = command.UpdateProductRequest{Id: int(product.Id), Name: product.Name, Price: product.Price}
	result := p.Db.Model(&product).Updates(updateProduct)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepositoryImpl) Delete(productId int) error {
	var product model.Product
	result := p.Db.Where("id = ?", productId).Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	var product []model.Product
	result := p.Db.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (p *ProductRepositoryImpl) FindById(productId int) (model.Product, error) {
	var product model.Product
	result := p.Db.Find(&product, productId)
	if result == nil {
		return product, errors.New("product is not found")
	}

	return product, nil
}
