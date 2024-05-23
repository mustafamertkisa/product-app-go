package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
)

func NewAuthMiddleware() fiber.Handler {
	secretKey := os.Getenv("SECRET_KEY")
	secret := []byte(secretKey)

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func ValidateJWTClaims(ctx *fiber.Ctx) bool {
	jwtCookie := ctx.Cookies("jwt")
	if jwtCookie == "" {
		fmt.Println("token cookie is missing")
		return false
	}

	token, err := jwt.Parse(jwtCookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println("error parsing JWT:", err)
		return false
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("failed to parse claims")
		return false
	}

	claimUserId, ok := claims["user_id"].(string)
	if !ok {
		fmt.Println("user_id claim is missing or not a string")
		return false
	}

	userId := ctx.Params("userId")
	if claimUserId != userId {
		fmt.Println("user_id claim does not match userId parameter")
		return false
	}

	return true
}
