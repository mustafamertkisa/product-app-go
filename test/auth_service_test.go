package test

import (
	"product-app-go/internal/application/command"
	"product-app-go/internal/application/service"
	"product-app-go/test/mocks"

	"testing"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func initAuthTestServices() (service.AuthService, service.UserService) {
	mockRepo := mocks.NewMockUserRepository()
	mockLogRepo := mocks.NewMockLogRepository()
	validate := validator.New()
	authService := service.NewAuthServiceImpl(mockRepo, mockLogRepo, validate)
	userService := service.NewUserServiceImpl(mockRepo, validate)

	return authService, userService
}

func TestRegister(t *testing.T) {
	authService, userService := initAuthTestServices()

	req := command.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "john123",
	}

	err := authService.Register(req)
	assert.Nil(t, err)

	createdUser, _ := userService.FindById(1)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, createdUser.Name)
	assert.Equal(t, req.Email, createdUser.Email)

	err = authService.Register(req)
	assert.NotNil(t, err)
	assert.Equal(t, "user already exists", err.Error())
}

func TestLogin(t *testing.T) {
	authService, _ := initAuthTestServices()

	req := command.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "john123",
	}
	err := authService.Register(req)
	assert.Nil(t, err)

	app := fiber.New()

	reqCtx := &fasthttp.RequestCtx{}
	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	loginReq := command.UserLoginRequest{
		Email:    "john.doe@example.com",
		Password: "john123",
	}

	token, err := authService.Login(loginReq, ctx)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	loggedInUser, err := authService.GetUserFromToken(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, loggedInUser)

	invalidLoginReq := command.UserLoginRequest{
		Email:    "john.doe@example.com",
		Password: "wrongpassword",
	}

	_, err = authService.Login(invalidLoginReq, ctx)
	assert.NotNil(t, err)
	assert.Equal(t, "incorrect password", err.Error())

	nonExistentUserReq := command.UserLoginRequest{
		Email:    "non.existent@example.com",
		Password: "password",
	}

	_, err = authService.Login(nonExistentUserReq, ctx)
	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestLogout(t *testing.T) {
	authService, _ := initAuthTestServices()

	req := command.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "john123",
	}
	err := authService.Register(req)
	assert.Nil(t, err)

	app := fiber.New()

	reqCtx := &fasthttp.RequestCtx{}
	ctx := app.AcquireCtx(reqCtx)
	defer app.ReleaseCtx(ctx)

	loginReq := command.UserLoginRequest{
		Email:    "john.doe@example.com",
		Password: "john123",
	}
	_, err = authService.Login(loginReq, ctx)
	assert.Nil(t, err)

	cookie := ctx.Response().Header.Peek("Set-Cookie")
	assert.Contains(t, string(cookie), "jwt=")

	err = authService.Logout(ctx)
	assert.Nil(t, err)

	cookie = ctx.Response().Header.Peek("Set-Cookie")
	assert.Contains(t, string(cookie), "jwt=;")
	assert.Contains(t, string(cookie), "expires=")
}
