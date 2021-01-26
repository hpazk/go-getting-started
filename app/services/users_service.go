package services

import (
	"sync"

	"github.com/hpazk/go-echo-rest-api/app/database"
	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/hpazk/go-echo-rest-api/app/models/users"
	UserModel "github.com/hpazk/go-echo-rest-api/app/models/users"
)

type usersService struct{}

var singleton UsersService
var once sync.Once

func GetUsersService() UsersService {
	if singleton != nil {
		return singleton
	}
	once.Do(func() {
		singleton = &usersService{}
	})
	return singleton
}

func SetUsersService(service UsersService) UsersService {
	original := singleton
	singleton = service
	return original
}

type UsersService interface {
	FindUserByEmail(email string) *UserModel.User
	AddUser(email string, password string, name string) *UserModel.User
}

func (u *usersService) FindUserByEmail(email string) *UserModel.User {
	db := database.GetInstance()
	var user UserModel.User
	err := db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}
	return nil
}

func (u *usersService) AddUser(email string, password string, name string) *UserModel.User {
	user := users.User{
		Role:     helpers.Customer,
		Email:    email,
		Password: password,
		Name:     name,
	}
	db := database.GetInstance()
	db.Create(&user)
	return &user
}
