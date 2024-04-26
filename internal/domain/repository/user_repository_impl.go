package repository

import (
	"errors"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"

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
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = command.UpdateUserRequest{Id: int(user.Id), Name: user.Name, Email: user.Email}
	result := u.Db.Model(&model.User{}).Where("id = ?", user.Id).Updates(updateUser)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *UserRepositoryImpl) FindAll() []model.User {
	var user []model.User
	result := u.Db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

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
