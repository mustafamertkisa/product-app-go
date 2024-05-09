package repository

import "product-app-go/internal/domain/model"

type UserRepository interface {
	Save(user model.User) error
	Update(user model.User) error
	Delete(userId int) error
	FindById(userId int) (model.User, error)
	FindByEmail(userEmail string) (model.User, error)
	FindAll() ([]model.User, error)
	AddLogToMongo(log model.LoginLog) error
}
