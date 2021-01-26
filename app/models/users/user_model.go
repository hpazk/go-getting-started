package users

import (
	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/hpazk/go-echo-rest-api/app/models"
)

type User struct {
	models.Base
	Email       string `gorm:"type:varchar(255);unique_index"`
	Password    string
	Role        helpers.UserRole
	Name        string
	FirstName   string
	LastName    string
	PhoneNumber string
	City        string
	Address     string
}

func (user User) String() string {
	return user.Name
}

func (user *User) BeforeSave() (err error) {
	hashed, err := helpers.GetPasswordUtil().HashPassword(user.Password)
	user.Password = hashed
	return
}
