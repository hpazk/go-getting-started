package database

import (
	CategoryModel "github.com/hpazk/go-echo-rest-api/app/models/products"
	ProductModel "github.com/hpazk/go-echo-rest-api/app/models/products"
	UserModel "github.com/hpazk/go-echo-rest-api/app/models/users"
	"github.com/jinzhu/gorm"
)

func GetMigrations(db *gorm.DB) error {
	db.AutoMigrate(&CategoryModel.Category{})
	db.AutoMigrate(&ProductModel.Product{})
	db.AutoMigrate(&UserModel.User{})

	return nil
}

// func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
// 	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
// 		{
// 			ID: "2020080202",
// 			Migrate: func(tx *gorm.DB) error {
// 				if err := tx.AutoMigrate(&UserModel.User{}).Error; err != nil {
// 					return err
// 				}
// 				return nil
// 			},
// 			Rollback: func(tx *gorm.DB) error {
// 				if err := tx.DropTable("users").Error; err != nil {
// 					return nil
// 				}
// 				return nil
// 			},
// 		},
// 		{
// 			ID: "2020080216",
// 			Migrate: func(tx *gorm.DB) error {
// 				if err := tx.AutoMigrate(&ProductModel.Product{}).Error; err != nil {
// 					return err
// 				}
// 				return nil
// 			},
// 			Rollback: func(tx *gorm.DB) error {
// 				if err := tx.DropTable("products").Error; err != nil {
// 					return nil
// 				}
// 				return nil
// 			},
// 		},
// 		{
// 			ID: "2020080217",
// 			Migrate: func(tx *gorm.DB) error {
// 				if err := tx.AutoMigrate(&CategoryModel.Category{}).Error; err != nil {
// 					return err
// 				}
// 				return nil
// 			},
// 			Rollback: func(tx *gorm.DB) error {
// 				if err := tx.DropTable("categories").Error; err != nil {
// 					return nil
// 				}
// 				return nil
// 			},
// 		},
// 	})
// }
