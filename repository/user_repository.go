package repository

import "product-app-go/model"

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
}
