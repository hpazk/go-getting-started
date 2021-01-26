package product

import (
	"github.com/hpazk/go-echo-rest-api/app/models"
)

type Product struct {
	models.Base
	Name          string
	Description   string
	Price         int
	ProductRating int
	PicturePath   string
	CategoryID    int
	Category      Category `gorm:"ForeignKey:CategoryID"`
}
