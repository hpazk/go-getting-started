package users

import (
	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/hpazk/go-echo-rest-api/app/models"
)

type User struct {
	models.Base
	Email    string `gorm:"type:varchar(255);unique_index"`
	Password string
	Name     string
	Role     helpers.UserRole
}

func (user User) String() string {
	return user.Name
}

func (user *User) BeforeSave() (err error) {
	hashed, err := helpers.GetPasswordUtil().HashPassword(user.Password)
	user.Password = hashed
	return
}
