package service

import (
	"errors"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"

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

func (u *UserServiceImpl) Create(user command.CreateUserRequest) error {
	err := u.validate.Struct(user)
	if err != nil {
		return err
	}

	userModel := model.User{
		Name:  user.Name,
		Email: user.Email,
	}

	err = u.UserRepository.Save(userModel)
	if err != nil {
		return errors.New("failed to save user: " + err.Error())
	}

	return nil
}

func (u *UserServiceImpl) Update(user command.UpdateUserRequest) error {
	userData, err := u.UserRepository.FindById(user.Id)
	if err != nil {
		return errors.New("failed to find user: " + err.Error())
	}

	userData.Name = user.Name
	userData.Email = user.Email

	err = u.UserRepository.Update(userData)
	if err != nil {
		return errors.New("failed to update user: " + err.Error())
	}

	return nil
}

func (u *UserServiceImpl) Delete(userId int) error {
	err := u.UserRepository.Delete(userId)
	if err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}

	return nil
}

func (u *UserServiceImpl) FindById(userId int) (command.UserResponse, error) {
	userData, err := u.UserRepository.FindById(userId)
	if err != nil {
		return command.UserResponse{}, errors.New("failed to find user: " + err.Error())
	}

	userResponse := command.UserResponse{
		Id:    int(userData.Id),
		Name:  userData.Name,
		Email: userData.Email,
	}

	return userResponse, nil
}

func (u *UserServiceImpl) FindAll() ([]command.UserResponse, error) {
	result, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, errors.New("failed to find users: " + err.Error())
	}

	var users []command.UserResponse

	for _, value := range result {
		user := command.UserResponse{
			Id:    int(value.Id),
			Name:  value.Name,
			Email: value.Email,
		}
		users = append(users, user)
	}

	return users, nil
}
