package common

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/labstack/echo/v4"
	"os"
	"fmt"
	"log"
)

type Application struct {
	logger echo.Logger
	server *echo.Echo
}

func NewMySQL() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Default().Println("Connected to database successfully")
	return db, nil
}