package config

import (
	"fmt"
	"product-app-go/internal/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	dsn := config.DBUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Connected successfully to the database")

	return db
}
