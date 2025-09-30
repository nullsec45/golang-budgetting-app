package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"github.com/joho/godotenv"
	"budgetting-app/common"
	"budgetting-app/cmd/api/handlers"
	"budgetting-app/cmd/api/middlewares"
	"fmt"
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
		e.Logger.Fatal("Error loading .env file, Error : ", err.Error())
	}


	h := handlers.Handler{
		DB: db,
	}

	app := Application{
		logger: e.Logger,
		server: e,
		handler: h,
	}

	e.Use(middlewares.CustomMiddleware, middleware.Logger(), middlewares.AnotherMiddleware)

	app.routes(h)

	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}