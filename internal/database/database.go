package database

import (
	"fmt"

	"github.com/yousafshah1214/go-comments-microservice/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("connecting to database")

	config, err := config.LoadConfig(".")

	if err != nil {
		return nil, err
	}

	conString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBPassword)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
