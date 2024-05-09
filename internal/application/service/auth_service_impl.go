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
	validate       *validator.Validate
}

func NewAuthServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
	}
}

func (s *AuthServiceImpl) Register(user command.CreateUserRequest) error {
	err := s.validate.Struct(user)
	if err != nil {
		return errors.New("validation failed: " + err.Error())
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

	var status int // 1 for success, 0 for failure
	var message string

	if userData.Id == 0 {
		status = 0
		message = "user not found"
	} else {
		if err := bcrypt.CompareHashAndPassword(userData.Password, []byte(user.Password)); err != nil {
			status = 0
			message = "incorrect password"
		} else {
			status = 1
			message = "successful login attempt for user"
		}
	}

	logEntry := model.LoginLog{
		Id:        primitive.NewObjectID(),
		Status:    status,
		Message:   message,
		UserId:    userData.Id,
		CreatedAt: time.Now(),
	}

	err = s.UserRepository.AddLogToMongo(logEntry)
	if err != nil {
		fmt.Println("Error logging login attempt:", err)
		return "", err
	}

	if status == 0 {
		return "", errors.New(message)
	}

	secretKey := os.Getenv("SECRET_KEY")
	secret := []byte(secretKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userData.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString(secret)
	if err != nil {
		return "", errors.New("could not create JWT token")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return token, nil
}

func (s *AuthServiceImpl) GetUserFromToken(cookie string) (model.User, error) {
	secretKey := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return model.User{}, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuerId, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to convert issuer ID to integer: %v", err)
	}

	var user model.User
	user, err = s.UserRepository.FindById(issuerId)
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
