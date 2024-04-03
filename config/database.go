package config

import (
	"fmt"
	"product-app-go/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Connected successfully to the database!")

	return db
}
