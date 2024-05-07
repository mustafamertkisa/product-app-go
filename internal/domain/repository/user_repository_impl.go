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

func (u *UserRepositoryImpl) Save(user model.User) error {
	result := u.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserRepositoryImpl) Update(user model.User) error {
	var updateUser = command.UpdateUserRequest{Id: int(user.Id), Name: user.Name, Email: user.Email}
	result := u.Db.Model(&model.User{}).Where("id = ?", user.Id).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserRepositoryImpl) Delete(userId int) error {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var user []model.User
	result := u.Db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result == nil {
		return user, errors.New("user is not found")
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(userEmail string) (model.User, error) {
	var user model.User
	result := u.Db.Where("email = ?", userEmail).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
