package repository

import (
	"errors"
	"product-app-go/internal/application/request"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/helper"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{Id: int(user.Id), Name: user.Name, Email: user.Email}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) FindAll() []model.User {
	var user []model.User
	result := u.Db.Find(&user)
	helper.ErrorPanic(result.Error)
	return user
}

func (u *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result == nil {
		return user, errors.New("user is not found")
	}

	return user, nil
}
