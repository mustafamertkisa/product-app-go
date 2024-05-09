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

func (r *ProductRepositoryImpl) Save(product model.Product) error {
	result := r.Db.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryImpl) Update(product model.Product) error {
	var updateProduct = command.UpdateProductRequest{Id: int(product.Id), Name: product.Name, Price: product.Price}
	result := r.Db.Model(&product).Updates(updateProduct)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryImpl) Delete(productId int) error {
	var product model.Product
	result := r.Db.Where("id = ?", productId).Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	var product []model.Product
	result := r.Db.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *ProductRepositoryImpl) FindById(productId int) (model.Product, error) {
	var product model.Product
	result := r.Db.Find(&product, productId)
	if result == nil {
		return product, errors.New("product is not found")
	}

	return product, nil
}
