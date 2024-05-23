package service

import (
	"errors"
	"fmt"
	"os"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	LogRepository  repository.LogRepository
	validate       *validator.Validate
}

func NewAuthServiceImpl(userRepository repository.UserRepository, logRepository repository.LogRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		LogRepository:  logRepository,
		validate:       validate,
	}
}

func (s *AuthServiceImpl) Register(user command.CreateUserRequest) error {
	err := s.validate.Struct(user)
	if err != nil {
		return errors.New("validation failed: " + err.Error())
	}

	existingUser, _ := s.UserRepository.FindByEmail(user.Email)
	if existingUser.Email != "" {
		return errors.New("user already exists")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}

	s.UserRepository.Save(userModel)

	return nil
}

func (s *AuthServiceImpl) Login(user command.UserLoginRequest, ctx *fiber.Ctx) (string, error) {
	userData, err := s.UserRepository.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	var success bool
	var message string

	if userData.Id == 0 {
		success = false
		message = "user not found"
	} else {
		if err := bcrypt.CompareHashAndPassword(userData.Password, []byte(user.Password)); err != nil {
			success = false
			message = "incorrect password"
		} else {
			success = true
			message = "successful login attempt for user"
		}
	}

	logEntry := model.LoginLog{
		Id:        primitive.NewObjectID(),
		Success:   success,
		Message:   message,
		UserId:    userData.Id,
		CreatedAt: time.Now(),
	}

	err = s.LogRepository.AddLogToMongo(logEntry)
	if err != nil {
		fmt.Println("Error logging login attempt:", err)
	}

	if !success {
		return "", errors.New(message)
	}

	secretKey := os.Getenv("SECRET_KEY")
	secret := []byte(secretKey)

	expectedExpTime := time.Now().Add(time.Hour * 2) // two hour

	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(int(userData.Id)),
		"exp":     expectedExpTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  expectedExpTime,
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return tokenString, nil
}

func (s *AuthServiceImpl) GetUserFromToken(cookie string) (model.User, error) {
	secretKey := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return model.User{}, err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return model.User{}, errors.New("invalid token")
	}

	userIdStr, ok := (*claims)["user_id"].(string)
	if !ok || userIdStr == "" {
		return model.User{}, errors.New("user ID is empty")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to convert user ID to integer: %v", err)
	}

	var user model.User
	user, err = s.UserRepository.FindById(userId)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *AuthServiceImpl) Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return nil
}
