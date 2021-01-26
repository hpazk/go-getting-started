package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/hpazk/go-echo-rest-api/app/config"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

var onceDb sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := config.DatabaseNew().(*config.DatabaseConfig)
		// PostgreSQL
		db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			databaseConfig.Psql.DbHost,
			databaseConfig.Psql.DbPort,
			databaseConfig.Psql.DbUsername,
			databaseConfig.Psql.DbDatabase,
			databaseConfig.Psql.DbPassword,
		))

		// MySQL
		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// 	databaseConfig.Psql.DbUsername,
		// 	databaseConfig.Psql.DbPassword,
		// 	databaseConfig.Psql.DbHost,
		// 	databaseConfig.Psql.DbPort,
		// 	databaseConfig.Psql.DbDatabase,
		// )
		// db, err := gorm.Open("mysql", dsn)

		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}
		instance = db
	})
	return instance
}
