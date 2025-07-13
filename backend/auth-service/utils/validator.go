package utils

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/shared-libs/pkg/response"
)

func ValidateStruct(data interface{}, validate *validator.Validate) *response.APIError{
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return response.NewAPIError(http.StatusBadRequest, "Validation failed", err.Error())
	}

	errorMessages := make(map[string]string)
	for _, fieldErr := range validationErrs {
		errorMessages[fieldErr.Field()] = fmt.Sprintf(
			"Field validation for '%s' failed on the '%s' tag (value: '%v', param: '%s')",
			fieldErr.Field(),
			fieldErr.Tag(),
			fieldErr.Value(),
			fieldErr.Param(),
		)
	}

	return response.NewAPIError(http.StatusBadRequest, "Validation failed", errorMessages)
}