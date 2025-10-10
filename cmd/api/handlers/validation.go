package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
	"budgetting-app/common"
	"fmt"
)

type ValidationError struct {
	Error string `json:"error"`
	Key   string `json:"key"`
	Condition string `json:"condition"`
}

func (h *Handler) ValidateBodyRequest(c echo.Context, payload interface{})[]*common.ValidationError{
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	var errors []*common.ValidationError
	
	err := validate.Struct(payload)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		reflected := reflect.ValueOf(payload)

		for _, validationErr := range validationErrors {
			field, _ := reflected.Type().FieldByName(validationErr.StructField())

			key := field.Tag.Get("json")

			if key == "" {
				key=strings.ToLower(validationErr.StructField())
			}

			condition := validationErr.Tag()
			keyToTitleCase := strings.Replace(key,"_","", -1)
			param := validationErr.Param()
			errMessage := keyToTitleCase + " field is "+condition

			switch condition {
				case "required":
					errMessage = keyToTitleCase + " is required"
				case "email":
					errMessage = keyToTitleCase + " must be a valid email address"
				case "min":
					if _, err := strconv.Atoi(param); err == nil {
						errMessage = fmt.Sprintf("%s must be at least %s characters", keyToTitleCase, param)
					}
				case "eqfield":
					errMessage = keyToTitleCase + " must be equal to "+strings.ToLower(param)
			}
			
		
			currentValidationError := &common.ValidationError{
				Error: errMessage,
				Key: key,
				Condition: condition,
			}

			errors = append(errors, currentValidationError)
		}
	}

	return errors
}