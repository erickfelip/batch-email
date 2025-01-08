package internalErrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0] // <-retornando o primeiro erro

	field := strings.ToLower(validationError.StructField()) // <- change to lower case
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "min":
		return errors.New(field + " is required with min " + validationError.Param())
	case "max":
		return errors.New(field + " is required with max " + validationError.Param())
	case "email":
		return errors.New(field + " is invalid")
	}
	return nil
}
