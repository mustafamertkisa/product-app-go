package service

import (
	"product-app-go/data/request"
	"product-app-go/data/response"
	"product-app-go/helper"
	"product-app-go/model"
	"product-app-go/repository"

	"github.com/go-playground/validator"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
	}
}

func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Name:  user.Name,
		Email: user.Email,
	}
	u.UserRepository.Save(userModel)
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Name = user.Name
	userData.Email = user.Email
	u.UserRepository.Update(userData)
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)
	userResponse := response.UserResponse{
		Id:    int(userData.Id),
		Name:  userData.Name,
		Email: userData.Email,
	}
	return userResponse
}

func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.UserRepository.FindAll()
	var users []response.UserResponse

	for _, value := range result {
		user := response.UserResponse{
			Id:    int(value.Id),
			Name:  value.Name,
			Email: value.Email,
		}
		users = append(users, user)
	}

	return users
}
