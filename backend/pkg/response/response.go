package response

import "github.com/labstack/echo/v4"

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Success(c echo.Context, code int, data interface{}) error {
	response := SuccessResponse{
		Status: "success",
		Data:   data,
	}
	return c.JSON(code, response)
}

func Error(c echo.Context, code int, message string) error {
	response := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	return c.JSON(code, response)
}