package helpers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Errors  interface{} `json:"errors"`
}

func ApiResponse(message string, code int, data interface{}) Response {

	status := "failed"
	if code == http.StatusOK {
		status = "success"
	}

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
		Errors:  nil,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func ValidatorError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) { // change err to form validation error
		errors = append(errors, e.Error())
	}

	return errors
}
