package auth

import (
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-echo-rest-api/app/config"
	"github.com/hpazk/go-echo-rest-api/app/middlewares"
	UserModel "github.com/hpazk/go-echo-rest-api/app/models/users"
)

type authService struct{}

var singleton AuthService
var once sync.Once

func GetAuthService() AuthService {
	once.Do(func() {
		singleton = &authService{}
	})
	return singleton
}

type AuthService interface {
	GetAccessToken(user *UserModel.User) (string, error)
}

func (s *authService) GetAccessToken(user *UserModel.User) (string, error) {
	claims := &middlewares.JWTCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * config.TokenExpiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
