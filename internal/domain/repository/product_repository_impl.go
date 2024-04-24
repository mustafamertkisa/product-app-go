package repository

import (
	"errors"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/utils"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p *ProductRepositoryImpl) Save(product model.Product) {
	result := p.Db.Create(&product)
	utils.ErrorPanic(result.Error)
}

func (p *ProductRepositoryImpl) Update(product model.Product) {
	var updateProduct = command.UpdateProductRequest{Id: int(product.Id), Name: product.Name, Price: product.Price}
	result := p.Db.Model(&product).Updates(updateProduct)
	utils.ErrorPanic(result.Error)
}

func (p *ProductRepositoryImpl) Delete(productId int) {
	var product model.Product
	result := p.Db.Where("id = ?", productId).Delete(&product)
	utils.ErrorPanic(result.Error)
}

func (p *ProductRepositoryImpl) FindAll() []model.Product {
	var product []model.Product
	result := p.Db.Find(&product)
	utils.ErrorPanic(result.Error)
	return product
}

func (p *ProductRepositoryImpl) FindById(productId int) (model.Product, error) {
	var product model.Product
	result := p.Db.Find(&product, productId)
	if result == nil {
		return product, errors.New("product is not found")
	}

	return product, nil
}
