package middlewares

import (
	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Budgeting-App-Server")
		return next(c)
	}
}