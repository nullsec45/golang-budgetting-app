package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckStruct struct {
	Health bool `json:"health"`
}

func (h *Handler) HealthCheck(c echo.Context) error {
	healthCheckStruct := HealthCheckStruct{
		Health: true,
	}

	return c.JSON(http.StatusOK, healthCheckStruct)
}

func (h *Handler) HelloWorld (c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}