package database

import (
	"github.com/yousafshah1214/go-comments-microservice/internal/comment"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&comment.Comment{}); err != nil {
		return err
	}

	return nil
}
