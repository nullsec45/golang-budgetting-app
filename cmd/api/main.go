package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"budgetting-app/common"
	"budgetting-app/cmd/api/handlers"
)

type Application struct {
	logger echo.Logger
	server *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()
	err := godotenv.Load()

	db, err := common.NewMySQL()

	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := handlers.Handler{
		DB: db,
	}

	app := Application{
		logger: e.Logger,
		server: e,
		handler: h,
	}

	fmt.Println(app)
	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}