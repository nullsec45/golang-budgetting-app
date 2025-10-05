package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"budgetting-app/cmd/api/requests"
	"fmt"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	payload := new(requests.RegisterUserRequest)	

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
		fmt.Println(payload)

	c.Logger().Info(payload)

	return c.String(http.StatusBadRequest, "good request")
}