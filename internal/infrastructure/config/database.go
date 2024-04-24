package config

import (
	"fmt"
	"product-app-go/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	dsn := config.DBUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.ErrorPanic(err)

	fmt.Println("Connected successfully to the database")

	return db
}
