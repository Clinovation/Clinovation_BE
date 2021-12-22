package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildSuccessResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err error, data interface{}) Response {
	errorMessage := err.Error()
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", err.Field())
			case "email":
				errorMessage = fmt.Sprintf("%s is required", err.Field())
			case "gte":
				errorMessage = fmt.Sprintf("%s is required", err.Field())
			case "lte":
				errorMessage = fmt.Sprintf("%s is required", err.Field())
			case "password":
				errorMessage = fmt.Sprintf("%s is not strong enough", err.Field())
			}
			break
		}
	}

	splitError := strings.Split(errorMessage, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
	return res
}