package response

import (
	"github.com/labstack/echo/v4"
)

// RegisterResponse is used for register responses.
type RegisterResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// LoginResponse is used for login responses.
type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

// ErrorResponse is used for error responses.
type ErrorsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RegisterSuccess(c echo.Context, code int, message string) error {
	response := RegisterResponse{
		Status:  "success",
		Message: message,
	}

	return c.JSON(code, response)
}

func LoginSuccess(c echo.Context, code int, token string) error {
	response := LoginResponse{
		Status: "success",
		Token:  token,
	}
	return c.JSON(code, response)
}

func ErrorResponsesAuth(c echo.Context, code int, message string) error {
	response := ErrorsResponse{
		Status:  "error",
		Message: message,
	}
	return c.JSON(code, response)
}