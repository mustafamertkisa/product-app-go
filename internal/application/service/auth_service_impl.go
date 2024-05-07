package service

import (
	"errors"
	"os"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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

func (a *AuthServiceImpl) Register(user command.CreateUserRequest) error {
	err := a.validate.Struct(user)
	if err != nil {
		return errors.New("validation failed: " + err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}

	a.UserRepository.Save(userModel)

	return nil
}

func (a *AuthServiceImpl) Login(user command.UserLoginRequest, ctx *fiber.Ctx) (string, error) {
	userData, err := a.UserRepository.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if userData.Id == 0 {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword(userData.Password, []byte(user.Password)); err != nil {
		return "", errors.New("incorrect password")
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
