package controllers

import "github.com/labstack/echo/v4"

type Response[T any] struct {
	Error         bool   `json:"error"`
	Code_Response int    `json:"code_response"`
	Message       string `json:"message"`
	Data          T      `json:"data"`
}

func NewResponse[T any](c echo.Context, statusCode int, Code_Response int, statusMessage bool, message string, data T) error {
	return c.JSON(statusCode, Response[T]{
		Error:         statusMessage,
		Code_Response: Code_Response,
		Message:       message,
		Data:          data,
	})
}
