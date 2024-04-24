package service

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"
	"product-app-go/utils"

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

func (u *UserServiceImpl) Create(user command.CreateUserRequest) {
	err := u.validate.Struct(user)
	utils.ErrorPanic(err)
	userModel := model.User{
		Name:  user.Name,
		Email: user.Email,
	}
	u.UserRepository.Save(userModel)
}

func (u *UserServiceImpl) Update(user command.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	utils.ErrorPanic(err)

	userData.Name = user.Name
	userData.Email = user.Email
	u.UserRepository.Update(userData)
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindById(userId int) command.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	utils.ErrorPanic(err)
	userResponse := command.UserResponse{
		Id:    int(userData.Id),
		Name:  userData.Name,
		Email: userData.Email,
	}
	return userResponse
}

func (u *UserServiceImpl) FindAll() []command.UserResponse {
	result := u.UserRepository.FindAll()
	var users []command.UserResponse

	for _, value := range result {
		user := command.UserResponse{
			Id:    int(value.Id),
			Name:  value.Name,
			Email: value.Email,
		}
		users = append(users, user)
	}

	return users
}
