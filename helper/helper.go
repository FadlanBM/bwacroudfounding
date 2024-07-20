package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func NewResponse(message string, code int, status string, data interface{}) *Response {
	return &Response{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}

func FormatValidationInputError(err error) []string {
	var errors []string

	for _, e:= range err.(validator.ValidationErrors) {
		errors=append(errors, e.Error())
	}

	return errors
}
