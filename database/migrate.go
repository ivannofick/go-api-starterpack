package database

import (
	"go-crud/app/models"
	"go-crud/config"
)

func RunMigration() {
	config.DB.AutoMigrate(&models.User{})
}
