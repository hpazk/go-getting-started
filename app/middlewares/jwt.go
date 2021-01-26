package middlewares

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	uuid "github.com/satori/go.uuid"
)

type JWTCustomClaims struct {
	ID   uuid.UUID        `json:"id"`
	Name string           `json:"name"`
	Role helpers.UserRole `json:"role"`
	jwt.StandardClaims
}

func JWTMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(key),
	})
}
