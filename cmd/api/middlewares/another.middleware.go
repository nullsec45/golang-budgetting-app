package middlewares

import (
	"github.com/labstack/echo/v4"
	"fmt"
)

func AnotherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Another Middleware Executed")
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}