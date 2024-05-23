package test

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/test/mocks"

	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func initUserTestServices() service.UserService {
	mockRepo := mocks.NewMockUserRepository()
	validate := validator.New()
	userService := service.NewUserServiceImpl(mockRepo, validate)

	return userService
}

func TestCreateUser(t *testing.T) {
	userService := initUserTestServices()

	req := command.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	err := userService.Create(req)
	assert.Nil(t, err)

	err = userService.Create(req)
	assert.NotNil(t, err)
	assert.Equal(t, "user already exists", err.Error())
}

func TestUpdateUser(t *testing.T) {
	userService := initUserTestServices()

	createReq := command.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	_ = userService.Create(createReq)

	updateReq := command.UpdateUserRequest{
		Id:    1,
		Name:  "John Smith",
		Email: "john.smith@example.com",
	}

	err := userService.Update(updateReq)
	assert.Nil(t, err)

	_ = userService.Create(command.CreateUserRequest{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	})

	updateReq.Email = "jane.doe@example.com"
	err = userService.Update(updateReq)
	assert.NotNil(t, err)
	assert.Equal(t, "this email is already in use", err.Error())
}

func TestDeleteUser(t *testing.T) {
	userService := initUserTestServices()

	createReq := command.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	_ = userService.Create(createReq)

	err := userService.Delete(1)
	assert.Nil(t, err)

	err = userService.Delete(2)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to delete order: user not found", err.Error())
}

func TestFindById(t *testing.T) {
	userService := initUserTestServices()

	createReq := command.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	_ = userService.Create(createReq)

	user, err := userService.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)

	_, err = userService.FindById(2)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to find user: user not found", err.Error())
}

func TestFindAllUsers(t *testing.T) {
	userService := initUserTestServices()

	_ = userService.Create(command.CreateUserRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	})

	_ = userService.Create(command.CreateUserRequest{
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	})

	users, err := userService.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}
