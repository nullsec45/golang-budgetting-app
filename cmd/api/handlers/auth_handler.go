package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"budgetting-app/cmd/api/requests"
	// "github.com/go-playground/validator/v10"
	"budgetting-app/common"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	payload := new(requests.RegisterUserRequest)	

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(c, *payload)

	if validationErrors != nil {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	c.Logger().Info(payload)
	// var validate *validator.Validate
	// validate = validator.New(validator.WithRequiredStructEnabled())
	// err := validate.Struct(payload)
	// validationErrors := err.(validator.ValidationErrors)
	// c.Logger().Print(validationErrors)

	return c.String(http.StatusOK, "good request")
}