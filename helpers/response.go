package helpers

import (
	"fmt"
	"github.com/labstack/echo/v4"
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

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Page interface{} `json:"page,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c echo.Context, status int, data interface{}, args ...interface{}) error {
	res := BaseResponse{}
	res.Meta.Status = status
	res.Meta.Message = "Success"
	if data != "" && data != nil{
		res.Data = data
	}
	if len(args) > 0 {
		res.Page = args[0]
	}
	return c.JSON(status, res)
}