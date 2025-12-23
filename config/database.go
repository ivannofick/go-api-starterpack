package config

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	driver := strings.ToLower(os.Getenv("DB_DRIVER"))

	var (
		db  *gorm.DB
		err error
	)

	switch driver {

	case "postgres", "postgresql":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	default:
		panic("Unsupported DB_DRIVER. Use 'postgres' or 'mysql'")
	}

	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	DB = db
}
