package main

import (
	"budgetting-app/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	app.server.GET("/health", handler.HealthCheck)
	app.server.GET("/hello-world", handler.HelloWorld)
	app.server.POST("/register", handler.RegisterHandler)
}